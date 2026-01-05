package models

type ArtistInfo struct {
	Biography      string `json:"biography,omitempty"`
	MusicBrainzID  string `json:"musicBrainzId,omitempty"`
	SmallImageURL  string `json:"smallImageUrl,omitempty"`
	MediumImageURL string `json:"mediumImageUrl,omitempty"`
	LargeImageURL  string `json:"largeImageUrl,omitempty"`
}
