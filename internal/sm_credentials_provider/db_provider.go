package smcredentialsprovider

import (
	"database/sql"
	"errors"

	"github.com/tikhonp/openswingsonic/internal/db/models/auth"
)

type dbCredentialsProvider struct {
	users auth.Users
}

func NewDBCredentialsProvider(users auth.Users) *dbCredentialsProvider {
	return &dbCredentialsProvider{users: users}
}

func (p *dbCredentialsProvider) GetPasswordForUsername(username string) (string, error) {
	user, err := p.users.GetUserByUsername(username)
	if errors.Is(err, sql.ErrNoRows) {
		return "", ErrUserNotFound
	}
	if err != nil {
		return "", err
	}
	return user.Password, nil
}
