// Package models contains data structures for Subsonic API responses.
package models

type SubsonicBase struct {
	Status        string `json:"status" xml:"status,attr"`
	Version       string `json:"version" xml:"version,attr"`
	Type          string `json:"type" xml:"type,attr,omitempty"`
	ServerVersion string `json:"serverVersion" xml:"serverVersion,attr,omitempty"`
	OpenSubsonic  bool   `json:"openSubsonic" xml:"openSubsonic,attr,omitempty"`
}

type SubsonicError struct {
	Code    int    `json:"code" xml:"code,attr"`
	Message string `json:"message" xml:"message,attr,omitempty"`
}

func (e SubsonicError) Error() string {
	return e.Message
}
