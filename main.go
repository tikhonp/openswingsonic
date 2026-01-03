package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tikhonp/openswingsonic/internal/config"
	"github.com/tikhonp/openswingsonic/internal/db"
	opensubsonicauth "github.com/tikhonp/openswingsonic/internal/middleware/opensubsonic_auth"
	smcredentialsprovider "github.com/tikhonp/openswingsonic/internal/sm_credentials_provider"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
	"github.com/tikhonp/openswingsonic/internal/util"
)

func credentialProvider(cfg *config.Config, database db.ModelsFactory) (smcredentialsprovider.SMCredentialsProvider, error) {
	switch cfg.CredentialsProvider {
	case config.CredentialsProviderTypeDatabase:
		return smcredentialsprovider.NewDBCredentialsProvider(database.AuthUsers()), nil
	case config.CredentialsProviderTypeFile:
		return smcredentialsprovider.NewUsersFileCredentialsProvider(cfg.UsersFilePath)
	default:
		return nil, errors.New("unknown credentials provider type")
	}
}

func main() {
	cfg := config.ReadConfig()

	database, err := db.Connect(cfg.DatabasePath)
	if err != nil {
		panic(err)
	}

	client := swingmusic.NewClient(cfg.SwingsonicBaseURL)

	e := echo.New()

	e.Logger.SetHeader("[${time_rfc3339}] ${level} ${message}")
	e.Debug = cfg.Debug
	e.HideBanner = true
	e.Validator = util.NewDefaultValidator()
	e.HTTPErrorHandler = util.ErrorHandler

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Never gonna give you up!")
	})

	credentialProvider, err := credentialProvider(cfg, database)
	if err != nil {
		e.Logger.Fatal("Failed to initialize credentials provider: ", err)
	}
	osauth := opensubsonicauth.NewOpenSubsonicAuth(database.AuthSessions(), client, credentialProvider)

	protected := e.Group("/rest", opensubsonicauth.Middleware(osauth))

	protected.GET("/ping.view", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	e.Logger.Fatal(e.Start(cfg.Addr))
}
