package models

type AlbumInfo struct {
	Notes          string `json:"notes,omitempty"`
	MusicBrainzID  string `json:"musicBrainzId,omitempty"`
	SmallImageURL  string `json:"smallImageUrl,omitempty"`
	MediumImageURL string `json:"mediumImageUrl,omitempty"`
	LargeImageURL  string `json:"largeImageUrl,omitempty"`
}
