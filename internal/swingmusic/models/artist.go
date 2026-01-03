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
