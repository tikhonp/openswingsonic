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

func (e SubsonicError) Error() string {
	return e.Message
}

type SubsonicErrorResponseContents struct {
	SubsonicBase
	Error SubsonicError `json:"error"`
}

type SubsonicErrorResponse struct {
	Contents SubsonicErrorResponseContents `json:"subsonic-response"`
}
