package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/tikhonp/openswingsonic/internal/config"
	"github.com/tikhonp/openswingsonic/internal/db"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi"
	swmiddleware "github.com/tikhonp/openswingsonic/internal/middleware"
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
	case config.CredentialsProviderTypeEnv:
		return smcredentialsprovider.NewEnvCredentialsProvider(database.AuthUsers())
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

	client := swingmusic.NewClient(cfg.SwingMusicBaseURL, cfg.PublicSwingMusicURL)

	e := echo.New()

	// Echo v5 dropped the Echo.Debug/HideBanner fields and replaced the custom
	// logger with *slog.Logger. Debug mode is now expressed by letting the default
	// HTTP error handler expose the underlying error; the banner is suppressed via
	// StartConfig where the server is started below.
	e.HTTPErrorHandler = echo.DefaultHTTPErrorHandler(cfg.Debug)
	e.Validator = util.NewDefaultValidator()

	e.Use(middleware.RequestLoggerWithConfig(
		util.GetRequestLoggerConfig(cfg),
	))
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Never gonna give you up!")
	})

	e.Pre(swmiddleware.CutViewSuffix)
	credentialProvider, err := credentialProvider(cfg, database)
	if err != nil {
		log.Fatal("Failed to initialize credentials provider: ", err)
	}
	osauth := opensubsonicauth.NewOpenSubsonicAuth(database.AuthSessions(), client, credentialProvider)
	opensubsonicapi.ConfigureOpenSubsonicRoutes(e.Group("/rest"), osauth, client)

	// e.Start always prints the startup banner in v5, so start through StartConfig
	// to keep HideBanner. The signal context mirrors e.Start's graceful shutdown.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	startConfig := echo.StartConfig{Address: cfg.Addr, HideBanner: true}
	if err := startConfig.Start(ctx, e); err != nil {
		log.Fatal(err)
	}
}
