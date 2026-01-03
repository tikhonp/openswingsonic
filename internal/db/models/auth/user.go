// Package auth provides models for authentication.
package auth

import (
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int       `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type Users interface {
	// GetUserByUsername retrieves a user by their username
	GetUserByUsername(username string) (*User, error)
}

type users struct {
	db *sqlx.DB
}

func NewUsers(db *sqlx.DB) Users {
	return &users{db: db}
}

func (u *users) GetUserByUsername(username string) (*User, error) {
	var user User
	err := u.db.Get(&user, "SELECT * FROM swingmusic_users WHERE username=?", username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error fetching user by username:", err)
		}
		return nil, err
	}
	return &user, nil
}
