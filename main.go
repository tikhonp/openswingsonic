package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tikhonp/openswingsonic/internal/config"
	"github.com/tikhonp/openswingsonic/internal/db"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi"
	opensubsonicauth "github.com/tikhonp/openswingsonic/internal/middleware/opensubsonic_auth"
	smcredentialsprovider "github.com/tikhonp/openswingsonic/internal/sm_credentials_provider"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
	"github.com/tikhonp/openswingsonic/internal/util"
)

func credentialProvider(cfg *config.Config, database db.ModelsFactory) (smcredentialsprovider.SMCredentialsProvider, error) {
	log.Println("Using credentials provider type:", cfg.CredentialsProvider)
	switch cfg.CredentialsProvider {
	case config.CredentialsProviderTypeDatabase:
		return smcredentialsprovider.NewDBCredentialsProvider(database.AuthUsers()), nil
	case config.CredentialsProviderTypeFile:
		return smcredentialsprovider.NewUsersFileCredentialsProvider(cfg.UsersFilePath, database.AuthUsers())
	default:
		return nil, errors.New("unknown credentials provider type")
	}
}

func main() {
	log.Printf("OpenSwingSonic: version \"%s\"", util.AppVersion)

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
	opensubsonicapi.ConfigureOpenSubsonicRoutes(e.Group("/rest"), osauth, client)

	e.Logger.Fatal(e.Start(cfg.Addr))
}
