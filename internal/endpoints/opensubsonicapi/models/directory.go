package models

// https://opensubsonic.netlify.app/docs/responses/directory/

type Directory struct {
	ID    string  `json:"id" xml:"id,attr"`
	Name  string  `json:"name" xml:"name,attr"`
	Child []Child `json:"child,omitempty" xml:"child,omitempty"`
}

type Child struct {
	ID       string `json:"id" xml:"id,attr"`
	Parent   string `json:"parent,omitempty" xml:"parent,attr,omitempty"`
	IsDir    bool   `json:"isDir" xml:"isDir,attr"`
	Title    string `json:"title" xml:"title,attr"`
	Artist   string `json:"artist,omitempty" xml:"artist,attr,omitempty"`
	Album    string `json:"album,omitempty" xml:"album,attr,omitempty"`
	CoverArt string `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	Duration int64  `json:"duration,omitempty" xml:"duration,attr,omitempty"`
	Track    int64  `json:"track,omitempty" xml:"track,attr,omitempty"`
	Path     string `json:"path,omitempty" xml:"path,attr,omitempty"`
}
