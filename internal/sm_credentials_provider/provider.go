// Package smcredentialsprovider provides different implementations for retrieving
// credentials from swingmusic
package smcredentialsprovider

import "errors"

// ErrUserNotFound is returned when a username is not found in the credentials provider
var ErrUserNotFound = errors.New("user not found")

// SMCredentialsProvider is an interface for retrieving passwords for given usernames
// I want do it as interface to allow implement it by providing creds in .env, file, db or other ways
type SMCredentialsProvider interface {
	// GetPasswordForUsername retrieves the password for the given username
	GetPasswordForUsername(username string) (string, error)
}
