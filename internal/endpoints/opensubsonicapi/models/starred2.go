package models

// https://opensubsonic.netlify.app/docs/responses/starred2/

type Starred2 struct {
	Artist []ArtistID3 `json:"artist" xml:"artist"`
	Album  []AlbumID3  `json:"album" xml:"album"`
	Song   []Song      `json:"song" xml:"song"`
}
