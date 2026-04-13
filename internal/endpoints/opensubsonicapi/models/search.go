package models

// SearchResult returned by search endpoint.
type SearchResult struct {
	Artist []Artist   `json:"artist,omitempty" xml:"artist,omitempty"`
	Album  []AlbumID3 `json:"album,omitempty" xml:"album,omitempty"`
	Song   []Song     `json:"song,omitempty" xml:"song,omitempty"`
}

// SearchResult2 returned by search2 endpoint.
type SearchResult2 struct {
	Artist []Artist   `json:"artist,omitempty" xml:"artist,omitempty"`
	Album  []AlbumID3 `json:"album,omitempty" xml:"album,omitempty"`
	Song   []Song     `json:"song,omitempty" xml:"song,omitempty"`
}

// SearchResult3 returned by search3 endpoint.
type SearchResult3 struct {
	Artist []ArtistID3 `json:"artist,omitempty" xml:"artist,omitempty"`
	Album  []AlbumID3  `json:"album,omitempty" xml:"album,omitempty"`
	Song   []Song      `json:"song,omitempty" xml:"song,omitempty"`
}
