// Package models contains data structures for Subsonic API responses.
package models

type SubsonicBase struct {
	Status        string `json:"status"`
	Version       string `json:"version"`
	Type          string `json:"type"`
	ServerVersion string `json:"serverVersion"`
	OpenSubsonic  bool   `json:"openSubsonic"`
}

type SubsonicResponse struct {
	Contents SubsonicBase `json:"subsonic-response"`
}

type SubsonicError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SubsonicErrorResponse struct {
	Contents struct {
		SubsonicBase
		Error SubsonicError `json:"error"`
	} `json:"subsonic-response"`
}
