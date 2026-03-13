package models

// User represents the authenticated Swing Music user.
type User struct {
	Extra    any      `json:"extra"`
	ID       int64    `json:"id"`
	Image    *string  `json:"image"`
	Roles    []string `json:"roles"`
	Username string   `json:"username"`
}
