package models

// https://opensubsonic.netlify.app/docs/responses/starred2/

type Starred2 struct {
	Artist []ArtistID3 `json:"artist"`
	Album  []AlbumID3  `json:"album"`
	Song   []Song      `json:"song"`
}
