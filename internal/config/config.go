// Package config provides env configuration settings for the application.
package config

import (
	"log"
	"os"
)

type CredentialsProviderType string

const (
	// CredentialsProviderTypeFile indicates that credentials are provided via a file.
	CredentialsProviderTypeFile CredentialsProviderType = "file"
	// CredentialsProviderTypeDatabase indicates that credentials are provided via a database.
	CredentialsProviderTypeDatabase CredentialsProviderType = "database"
	// CredentialsProviderTypeEnv indicates that credentials are provided via environment variables.
	CredentialsProviderTypeEnv CredentialsProviderType = "env"
)

type Config struct {

	// Debug enables debug mode.
	Debug bool

	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	Addr string

	// DatabasePath specifies the path to the SQLite database file.
	DatabasePath string

	// SwingMusicBaseURL specifies the base URL for the swing music server.
	SwingMusicBaseURL string

	// PublicSwingMusicURL specifies the public URL for swing music avaliable to clients.
	PublicSwingMusicURL string

	// CredentialsProvider specifies the type of credentials provider to use.
	CredentialsProvider CredentialsProviderType

	// UsersFilePath specifies the path to the users file. Optional.
	// User if users file provider is used.
	UsersFilePath string

	// JSONLog enables JSON formatted logging.
	JSONLog bool
}

func ReadConfig() *Config {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		log.Println("WARNING: LISTEN_ADDR not set")
	}
	databasePath := os.Getenv("DATABASE_PATH")
	if databasePath == "" {
		log.Fatalln("FATAL: DATABASE_PATH not set")
	}
	swingmusicBaseURL := os.Getenv("SWINGMUSIC_BASE_URL")
	if swingmusicBaseURL == "" {
		log.Fatalln("FATAL: SWINGMUSIC_BASE_URL not set")
	}
	publicSwingMusicURL := os.Getenv("PUBLIC_SWINGMUSIC_URL")
	if publicSwingMusicURL == "" {
		log.Println("WARNING: PUBLIC_SWINGMUSIC_URL not set, defaulting to SWINGMUSIC_BASE_URL")
		publicSwingMusicURL = swingmusicBaseURL
	}
	credentialsProvider := CredentialsProviderType(os.Getenv("CRED_PROVIDER"))
	if credentialsProvider == "" {
		log.Println("WARNING: CRED_PROVIDER not set, defaulting to 'database'")
		credentialsProvider = CredentialsProviderTypeDatabase
	}
	usersFilePath := os.Getenv("USERS_FILE_PATH")
	if credentialsProvider == CredentialsProviderTypeFile && usersFilePath == "" {
		log.Fatalln("FATAL: USERS_FILE_PATH not set, required when CRED_PROVIDER is 'file'")
	}
	return &Config{
		Debug:               os.Getenv("DEBUG") == "true",
		Addr:                addr,
		DatabasePath:        databasePath,
		SwingMusicBaseURL:   swingmusicBaseURL,
		PublicSwingMusicURL: publicSwingMusicURL,
		CredentialsProvider: credentialsProvider,
		UsersFilePath:       usersFilePath,
		JSONLog:             os.Getenv("JSON_LOG") == "true",
	}
}
