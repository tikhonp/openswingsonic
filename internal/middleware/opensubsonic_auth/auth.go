package opensubsonicauth

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/tikhonp/openswingsonic/internal/db/models/auth"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	smcredentialsprovider "github.com/tikhonp/openswingsonic/internal/sm_credentials_provider"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
)

type OpenSubsonicAuth interface {
	// Authentificate performs authentication based on the provided parameters.
	// It returns a session key if authentication is successful.
	Authentificate(params authParams) (sessionKey string, err error)
}

type openSubsonicAuth struct {
	sessions    auth.Sessions
	swingmusic  swingmusic.SwingMusicClient
	credentials smcredentialsprovider.SMCredentialsProvider
}

func NewOpenSubsonicAuth(
	sessions auth.Sessions,
	sm swingmusic.SwingMusicClient,
	creds smcredentialsprovider.SMCredentialsProvider,
) OpenSubsonicAuth {
	return &openSubsonicAuth{
		sessions:    sessions,
		swingmusic:  sm,
		credentials: creds,
	}
}

// validateData checks for conflicting authentication parameters.
//
// It respects the following rules:
// *) Either p or both t and s must be specified.
// **) If apiKey is specified, then none of p, t, s, nor u can be specified.
func (osa *openSubsonicAuth) validateData(p authParams) error {
	if p.APIKey != "" && (p.U != "" || p.P != "" || p.T != "" || p.S != "") {
		return &middleware.MultipleConflictingAuthMechanisms
	}
	if p.U != "" && p.P != "" && (p.APIKey != "" || p.T != "" || p.S != "") {
		return &middleware.MultipleConflictingAuthMechanisms
	}
	if p.U != "" && p.T != "" && p.S != "" && (p.APIKey != "" || p.P != "") {
		return &middleware.MultipleConflictingAuthMechanisms
	}
	return nil
}

func (osa *openSubsonicAuth) authentificateByUP(u, p string) (string, error) {
	password, err := osa.credentials.GetPasswordForUsername(u)
	if err != nil {
		return "", middleware.WrongUsernameOrPassword
	}

	// p is the password, either in clear text or hex-encoded with a “enc:” prefix.
	// check for the "enc:" prefix and decode if necessary
	if strings.HasPrefix(p, "enc:") {
		decoded, err := hex.DecodeString(p[4:])
		if err != nil {
			log.Println("Failed to decode hex password:", err)
			return "", middleware.WrongUsernameOrPassword
		}
		p = string(decoded)
	}

	if password != p {
		return "", middleware.WrongUsernameOrPassword
	}

	return osa.getSessionKeyForUser(u, password)
}

func (osa *openSubsonicAuth) authentificateByUTS(u, t, s string) (string, error) {
	password, err := osa.credentials.GetPasswordForUsername(u)
	if err != nil {
		return "", middleware.WrongUsernameOrPassword
	}

	// From doc
	//
	// Starting with API version 1.13.0, the recommended authentication scheme
	// is to send an authentication token, calculated as a one-way salted hash of the password.
	//
	// This involves two steps:
	// 1. For each REST call, generate a random string called the salt.
	//    Send this as parameter s. Use a salt length of at least six characters.
	// 2. Calculate the authentication token as follows: token = md5(password + salt).
	//    The md5() function takes a string and returns the 32-byte ASCII hexadecimal representation of the MD5 hash,
	//    using lower case characters for the hex values.
	//    The ‘+’ operator represents concatenation of the two strings.
	//    Treat the strings as UTF-8 encoded when calculating the hash. Send the result as parameter t.
	hashBytes := md5.Sum([]byte(password + s))
	expectedToken := hex.EncodeToString(hashBytes[:])
	if expectedToken != t {
		return "", middleware.WrongUsernameOrPassword
	}

	return osa.getSessionKeyForUser(u, password)
}

func (osa *openSubsonicAuth) authentificateByAPIKey(_ string) (string, error) {
	// API key authentication is not supported for now
	return "", middleware.TokenAuthNotSupported
}

func (osa *openSubsonicAuth) createNewSession(username, password string) (string, error) {
	cookie, err := osa.swingmusic.Login(username, password)
	if err != nil {
		return "", err
	}

	err = osa.sessions.InsertSession(username, cookie.Value, cookie.Expires)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

func (osa *openSubsonicAuth) getSessionKeyForUser(username, password string) (string, error) {
	session, err := osa.sessions.GetSessionByUsername(username)
	if errors.Is(err, sql.ErrNoRows) {
		return osa.createNewSession(username, password)
	} else if err != nil {
		return "", err
	}
	if session.ExpiresAt.Before(time.Now().Add(-5 * time.Minute)) {
		return osa.createNewSession(username, password)
	}
	return session.SessionToken, nil
}

func (osa *openSubsonicAuth) Authentificate(params authParams) (string, error) {
	if err := osa.validateData(params); err != nil {
		return "", err
	}

	// Detect authentication mechanism
	if params.U != "" && params.P != "" {
		// Username and password authentication
		return osa.authentificateByUP(params.U, params.P)
	} else if params.U != "" && params.T != "" && params.S != "" {
		// Token and salt authentication
		return osa.authentificateByUTS(params.U, params.T, params.S)
	} else if params.APIKey != "" {
		// API key authentication
		return osa.authentificateByAPIKey(params.APIKey)
	} else {
		return "", middleware.RequiredParametrIsMissing
	}
}
