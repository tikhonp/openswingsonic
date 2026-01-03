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

	// UpsertUser inserts a new user or updates the password if the user already exists
	UpsertUser(username, password string) error

	// BatchUpsertUsers inserts or updates multiple users in a single operation
	BatchUpsertUsers(users map[string]string) error
}

type users struct {
	db *sqlx.DB
}

func NewUsers(db *sqlx.DB) Users {
	return &users{db: db}
}

func (u *users) GetUserByUsername(username string) (*User, error) {
	var user User
	err := u.db.Get(&user, "SELECT * FROM swingmusic_users WHERE username=$1", username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error fetching user by username:", err)
		}
		return nil, err
	}
	return &user, nil
}

func (u *users) UpsertUser(username, password string) error {
	query := `
		INSERT INTO swingmusic_users (username, password),
		VALUES ($1, $2)
		ON CONFLICT (username) DO UPDATE SET password = EXCLUDED.password
	`
	_, err := u.db.Exec(query, username, password)
	return err
}

func (u *users) BatchUpsertUsers(users map[string]string) error {
	tx, err := u.db.Beginx()
	if err != nil {
		return err
	}
	stmt, err := tx.Preparex(`
		INSERT INTO swingmusic_users (username, password)
		VALUES ($1, $2)
		ON CONFLICT (username) DO UPDATE SET password = EXCLUDED.password
	`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for username, password := range users {
		_, err := stmt.Exec(username, password)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
