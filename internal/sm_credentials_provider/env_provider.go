package smcredentialsprovider

import (
	"errors"
	"fmt"
	"os"

	"github.com/tikhonp/openswingsonic/internal/db/models/auth"
)

type envCredentialsProvider struct {
	users map[string]string
}

func readUsersFromEnv() map[string]string {
	users := make(map[string]string)
	var i = 0
	for {
		usernameEnv := fmt.Sprintf("OSM_USER_%d_USERNAME", i)
		passwordEnv := fmt.Sprintf("OSM_USER_%d_PASSWORD", i)
		username := os.Getenv(usernameEnv)
		password := os.Getenv(passwordEnv)
		if username == "" || password == "" {
			break
		}
		users[username] = password
		i++
	}
	return users
}

// NewEnvCredentialsProvider creates a new instance of envCredentialsProvider
// that reads credentials from environment variables.
//
// The expected environment variables are in the format:
// OSM_USER_0_USERNAME, OSM_USER_0_PASSWORD
// OSM_USER_1_USERNAME, OSM_USER_1_PASSWORD
// and so on.
func NewEnvCredentialsProvider(users auth.Users) (*envCredentialsProvider, error) {
	usersMap := readUsersFromEnv()
	if len(usersMap) == 0 {
		return nil, errors.New("no users found in environment variables")
	}
	err := users.BatchUpsertUsers(usersMap)
	if err != nil {
		return nil, err
	}
	return &envCredentialsProvider{
		users: usersMap,
	}, nil
}

func (p *envCredentialsProvider) GetPasswordForUsername(username string) (string, error) {
	password, exists := p.users[username]
	if !exists {
		return "", ErrUserNotFound
	}
	return password, nil
}
