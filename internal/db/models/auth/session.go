package auth

import (
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type Session struct {
	ID           int       `db:"id"`
	UserID       int       `db:"user_id"`
	SessionToken string    `db:"session_token"`
	CreatedAt    time.Time `db:"created_at"`
	ExpiresAt    time.Time `db:"expires_at"`
}

type Sessions interface {
	// GetSessionByUsername retrieves a session by the associated user's username
	GetSessionByUsername(username string) (*Session, error)

	// InsertSession creates a new session for the specified username
	InsertSession(username, sessionToken string, expiresAt time.Time) error
}

type sessions struct {
	db *sqlx.DB
}

func NewSessions(db *sqlx.DB) Sessions {
	return &sessions{db: db}
}

func (s sessions) GetSessionByUsername(username string) (*Session, error) {
	var session Session
	query := `
		SELECT s.* FROM sessions s
		JOIN swingmusic_users u ON s.user_id = u.id
		WHERE u.username = $1
		ORDER BY s.created_at DESC
		LIMIT 1
	`
	err := s.db.Get(&session, query, username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error fetching session by user ID:", err)
		}
		return nil, err
	}
	return &session, nil
}

func (s sessions) InsertSession(username, sessionToken string, expiresAt time.Time) error {
	query := `
		INSERT INTO sessions (user_id, session_token, expires_at)
		VALUES (
			(SELECT id FROM swingmusic_users WHERE username = $1),
			$2,
			$3
		)
	`
	_, err := s.db.Exec(query, username, sessionToken, expiresAt)
	return err
}
