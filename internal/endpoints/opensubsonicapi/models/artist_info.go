package models

type ArtistInfo struct {
	Biography      string `json:"biography,omitempty" xml:"biography,omitempty"`
	MusicBrainzID  string `json:"musicBrainzId,omitempty" xml:"musicBrainzId,omitempty"`
	SmallImageURL  string `json:"smallImageUrl,omitempty" xml:"smallImageUrl,omitempty"`
	MediumImageURL string `json:"mediumImageUrl,omitempty" xml:"mediumImageUrl,omitempty"`
	LargeImageURL  string `json:"largeImageUrl,omitempty" xml:"largeImageUrl,omitempty"`
}
