package models

type AlbumResponse struct {
	Copyright     string                      `json:"copyright"`
	Extra         AlbumExtra                  `json:"extra"`
	Info          AlbumInfo                   `json:"info"`
	MoreFrom      map[string][]AlbumShortInfo `json:"more_from"`
	OtherVersions []AlbumShortInfo            `json:"other_versions"`
	Stats         []AlbumStat                 `json:"stats"`
	Tracks        []Track                     `json:"tracks"`
}

type AlbumExtra struct {
	AvgBitrate int64 `json:"avg_bitrate"`
	TrackTotal int64 `json:"track_total"`
}

type AlbumInfo struct {
	Score        float64  `json:"_score"`
	AlbumArtists []Artist `json:"albumartists"`
	AlbumHash    string   `json:"albumhash"`
	ArtistHashes []string `json:"artisthashes"`
	BaseTitle    string   `json:"base_title"`
	Color        string   `json:"color"`
	CreatedDate  int64    `json:"created_date"`
	Date         int64    `json:"date"`
	Duration     int64    `json:"duration"`
	Extra        any      `json:"extra"` // usually empty object
	FavUserIDs   []string `json:"fav_userids"`
	GenreHashes  string   `json:"genrehashes"`
	Genres       []Genre  `json:"genres"`
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	IsFavorite   bool     `json:"is_favorite"`
	LastPlayed   int64    `json:"lastplayed"`
	OGTitle      string   `json:"og_title"`
	PathHash     string   `json:"pathhash"`
	PlayCount    int64    `json:"playcount"`
	PlayDuration int64    `json:"playduration"`
	Title        string   `json:"title"`
	TrackCount   int64    `json:"trackcount"`
	Type         string   `json:"type"`
	Versions     []any    `json:"versions"`
	WeakHash     string   `json:"weakhash"`
}

type Genre struct {
	GenreHash string `json:"genrehash"`
	Name      string `json:"name"`
}

type AlbumShortInfo struct {
	Score        float64  `json:"_score"`
	AlbumArtists []Artist `json:"albumartists"`
	AlbumHash    string   `json:"albumhash"`
	Color        string   `json:"color"`
	Date         int64    `json:"date"`
	Image        string   `json:"image"`
	PathHash     string   `json:"pathhash"`
	Title        string   `json:"title"`
	Type         string   `json:"type"`
	Versions     []any    `json:"versions"`
}

type AlbumStat struct {
	CSSClass string  `json:"cssclass"`
	Image    *string `json:"image"`
	Text     string  `json:"text"`
	Value    string  `json:"value"`
}

type Albums struct {
	Items []AlbumShortInfo `json:"items"`
	Total int64            `json:"total"`
}
