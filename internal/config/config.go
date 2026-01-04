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

	// SwingsonicBaseURL specifies the base URL for the Swingsonic server.
	SwingsonicBaseURL string

	// CredentialsProvider specifies the type of credentials provider to use.
	CredentialsProvider CredentialsProviderType

	// UsersFilePath specifies the path to the users file. Optional.
	// User if users file provider is used.
	UsersFilePath string
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
	swingsonicBaseURL := os.Getenv("SWINGSONIC_BASE_URL")
	if swingsonicBaseURL == "" {
		log.Fatalln("FATAL: SWINGSONIC_BASE_URL not set")
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
		SwingsonicBaseURL:   swingsonicBaseURL,
		CredentialsProvider: credentialsProvider,
		UsersFilePath:       usersFilePath,
	}
}
