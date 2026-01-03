// Package models contains data structures for swing music client API.
package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
