// Package models contains data structures for Subsonic API responses.
package models

type SubsonicBase struct {
	Status        string `json:"status" xml:"status"`
	Version       string `json:"version" xml:"version"`
	Type          string `json:"type" xml:"type"`
	ServerVersion string `json:"serverVersion" xml:"serverVersion"`
	OpenSubsonic  bool   `json:"openSubsonic" xml:"openSubsonic"`
}

type SubsonicError struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}

func (e SubsonicError) Error() string {
	return e.Message
}
