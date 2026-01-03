package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/tikhonp/openswingsonic/internal/db/models/auth"
)

type ModelsFactory interface {
	AuthUsers() auth.Users
	AuthSessions() auth.Sessions
}

type modelsFactory struct {
	users    auth.Users
	sessions auth.Sessions
}

func newModelsFactory(db *sqlx.DB) ModelsFactory {
	return &modelsFactory{
		users:    auth.NewUsers(db),
		sessions: auth.NewSessions(db),
	}
}

func (f *modelsFactory) AuthUsers() auth.Users {
	return f.users
}

func (f *modelsFactory) AuthSessions() auth.Sessions {
	return f.sessions
}
