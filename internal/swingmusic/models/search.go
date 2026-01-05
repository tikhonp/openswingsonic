package models

type SearchedTracks struct {
	More    bool    `json:"more"`
	Results []Track `json:"results"`
}

type SearchedAlbums struct {
	More    bool             `json:"more"`
	Results []AlbumShortInfo `json:"results"`
}

type SearchedArtists struct {
	More    bool         `json:"more"`
	Results []ArtistItem `json:"results"`
}

type SearchedAll struct {
	Albums    []AlbumShortInfo `json:"albums"`
	Artists   []ArtistItem     `json:"artists"`
	TopResult TopResult        `json:"top_result"`
	Tracks    []Track          `json:"tracks"`
}

type TopResult struct {
	Albumcount int64  `json:"albumcount"`
	Artisthash string `json:"artisthash"`
	Color      string `json:"color"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	Trackcount int64  `json:"trackcount"`
	Type       string `json:"type"`
}
