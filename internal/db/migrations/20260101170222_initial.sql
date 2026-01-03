-- +goose Up
-- +goose StatementBegin
CREATE TABLE swingmusic_users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(256) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER REFERENCES swingmusic_users(id),
    session_token VARCHAR(512) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE swingmusic_users;
DROP TABLE sessions;
-- +goose StatementEnd
