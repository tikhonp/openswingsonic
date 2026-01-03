package models

// https://opensubsonic.netlify.app/docs/responses/musicfolders/

type MusicFolders struct {
	MusicFolder []MusicFolder `json:"musicFolder"`
}

type MusicFolder struct {
	ID   int64   `json:"id"`
	Name *string `json:"name,omitempty"`
}
