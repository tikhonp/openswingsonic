package smcredentialsprovider

import (
	"os"
	"strings"

	"github.com/tikhonp/openswingsonic/internal/db/models/auth"
)

type usersFileCredentialsProvider struct {
	users map[string]string
}

// ReadUsersFile reads the users file and returns a map of username to password.
// The users file should have lines in the format: username:password
func readUsersFile(usersFilePath string) (map[string]string, error) {
	content, err := os.ReadFile(usersFilePath)
	if err != nil {
		return nil, err
	}

	users := make(map[string]string)
	lines := strings.SplitSeq(string(content), "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue // or return an error if the format is strict
		}
		username := strings.TrimSpace(parts[0])
		password := strings.TrimSpace(parts[1])
		users[username] = password
	}

	return users, nil
}

// NewUsersFileCredentialsProvider creates a new instance of usersFileCredentialsProvider
// that reads credentials from the specified users file.
//
// The users file should have lines in the format: username:password
func NewUsersFileCredentialsProvider(usersFilePath string, users auth.Users) (SMCredentialsProvider, error) {
	usersMap, err := readUsersFile(usersFilePath)
	if err != nil {
		return nil, err
	}
	err = users.BatchUpsertUsers(usersMap)
	if err != nil {
		return nil, err
	}
	return &usersFileCredentialsProvider{users: usersMap}, nil
}

func (p *usersFileCredentialsProvider) GetPasswordForUsername(username string) (string, error) {
	password, exists := p.users[username]
	if !exists {
		return "", ErrUserNotFound
	}
	return password, nil
}
