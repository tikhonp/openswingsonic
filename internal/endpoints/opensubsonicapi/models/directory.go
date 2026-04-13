package models

// https://opensubsonic.netlify.app/docs/responses/directory/

type Directory struct {
	ID    string  `json:"id" xml:"id"`
	Name  string  `json:"name" xml:"name"`
	Child []Child `json:"child,omitempty" xml:"child,omitempty"`
}

type Child struct {
	ID       string `json:"id" xml:"id"`
	Parent   string `json:"parent,omitempty" xml:"parent,omitempty"`
	IsDir    bool   `json:"isDir" xml:"isDir"`
	Title    string `json:"title" xml:"title"`
	Artist   string `json:"artist,omitempty" xml:"artist,omitempty"`
	Album    string `json:"album,omitempty" xml:"album,omitempty"`
	CoverArt string `json:"coverArt,omitempty" xml:"coverArt,omitempty"`
	Duration int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	Track    int64  `json:"track,omitempty" xml:"track,omitempty"`
	Path     string `json:"path,omitempty" xml:"path,omitempty"`
}
