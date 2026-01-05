package models

type Artists struct {
	Items []ArtistItem `json:"items"`
	Total int64        `json:"total"`
}

type ArtistItem struct {
	Artisthash string `json:"artisthash"`
	Color      string `json:"color"`
	HelpText   string `json:"help_text"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	Type       string `json:"type"`
}

type ArtistResponse struct {
	Albums ArtistAlbumsResponse `json:"albums"`
	Artist ArtistDetail         `json:"artist"`
	Stats  []AlbumStat          `json:"stats"`
	Tracks []Track              `json:"tracks"`
}

type ArtistDetail struct {
	AlbumCount int64   `json:"albumcount"`
	ArtistHash string  `json:"artisthash"`
	Color      string  `json:"color"`
	Duration   int64   `json:"duration"`
	Genres     []Genre `json:"genres"`
	Image      string  `json:"image"`
	IsFavorite bool    `json:"is_favorite"`
	Name       string  `json:"name"`
	TrackCount int64   `json:"trackcount"`
	Type       string  `json:"type"`
}

type ArtistAlbumsResponse struct {
	Albums        []AlbumShortInfo `json:"albums"`
	Appearances   []AlbumShortInfo `json:"appearances"`
	ArtistName    string           `json:"artistname"`
	Compilations  []AlbumShortInfo `json:"compilations"`
	SinglesAndEPs []AlbumShortInfo `json:"singles_and_eps"`
}
