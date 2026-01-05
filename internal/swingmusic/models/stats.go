package models

type Stats struct {
	Dates string `json:"dates"`
	Stats []Stat `json:"stats"`
}

type Stat struct {
	Cssclass string  `json:"cssclass"`
	Image    *string `json:"image,omitempty"`
	Text     string  `json:"text"`
	Value    string  `json:"value"`
}

type TopAlbums struct {
	Albums    []AlbumElementStat `json:"albums"`
	Scrobbles Scrobbles          `json:"scrobbles"`
}

type Scrobbles struct {
	Dates string `json:"dates"`
	Text  string `json:"text"`
	Trend string `json:"trend"`
}

type TopArtists struct {
	Artists   []ArtistElementStat `json:"artists"`
	Scrobbles Scrobbles           `json:"scrobbles"`
}

type ArtistElementStat struct {
	Artisthash string          `json:"artisthash"`
	Color      string          `json:"color"`
	Extra      ExtraStatArtist `json:"extra"`
	HelpText   string          `json:"help_text"`
	Image      string          `json:"image"`
	Name       string          `json:"name"`
	Trend      TrendClass      `json:"trend"`
	Type       string          `json:"type"`
}

type ExtraStatArtist struct {
	Playcount int64 `json:"playcount"`
}

type TrendClass struct {
	IsNew bool   `json:"is_new"`
	Trend string `json:"trend"`
}

type AlbumElementStat struct {
	Score        int64         `json:"_score"`
	Albumartists []Albumartist `json:"albumartists"`
	Albumhash    string        `json:"albumhash"`
	Color        *string       `json:"color"`
	Date         int64         `json:"date"`
	HelpText     string        `json:"help_text"`
	Image        string        `json:"image"`
	Pathhash     string        `json:"pathhash"`
	Title        string        `json:"title"`
	Trend        TrendClass    `json:"trend"`
	Type         string        `json:"type"`
	Versions     []any         `json:"versions"`
}

type Albumartist struct {
	Artisthash string `json:"artisthash"`
	Name       string `json:"name"`
}

type TopTracks struct {
	Scrobbles Scrobbles   `json:"scrobbles"`
	Tracks    []TrackStat `json:"tracks"`
}

type TrackStat struct {
	Album        string     `json:"album"`
	Albumartists []Artist   `json:"albumartists"`
	Albumhash    string     `json:"albumhash"`
	Artisthashes []string   `json:"artisthashes"`
	Artists      []Artist   `json:"artists"`
	Bitrate      int64      `json:"bitrate"`
	Duration     int64      `json:"duration"`
	Explicit     bool       `json:"explicit"`
	Extra        ExtraStat  `json:"extra"`
	Filepath     string     `json:"filepath"`
	Folder       string     `json:"folder"`
	HelpText     string     `json:"help_text"`
	Image        string     `json:"image"`
	IsFavorite   bool       `json:"is_favorite"`
	Title        string     `json:"title"`
	Trackhash    string     `json:"trackhash"`
	Trend        TrendClass `json:"trend"`
	Weakhash     string     `json:"weakhash"`
}

type ExtraStat struct {
	Albumartistsort           []string `json:"albumartistsort,omitempty"`
	Artist                    []string `json:"artist"`
	Artistsort                []string `json:"artistsort,omitempty"`
	Barcode                   []string `json:"barcode,omitempty"`
	Bitdepth                  int64    `json:"bitdepth"`
	CatalogNumber             []string `json:"catalog_number,omitempty"`
	Channels                  int64    `json:"channels"`
	Composer                  []string `json:"composer"`
	DiscTotal                 *int64   `json:"disc_total,omitempty"`
	Filesize                  int64    `json:"filesize"`
	Genre                     []string `json:"genre"`
	Hashinfo                  Hashinfo `json:"hashinfo"`
	Isrc                      []string `json:"isrc,omitempty"`
	Label                     []string `json:"label,omitempty"`
	Media                     []string `json:"media,omitempty"`
	MusicbrainzAlbumartistid  []string `json:"musicbrainz_albumartistid,omitempty"`
	MusicbrainzAlbumid        []string `json:"musicbrainz_albumid,omitempty"`
	MusicbrainzArtistid       []string `json:"musicbrainz_artistid,omitempty"`
	MusicbrainzReleasegroupid []string `json:"musicbrainz_releasegroupid,omitempty"`
	MusicbrainzReleasetrackid []string `json:"musicbrainz_releasetrackid,omitempty"`
	MusicbrainzTrackid        []string `json:"musicbrainz_trackid,omitempty"`
	Originaldate              []string `json:"originaldate,omitempty"`
	Originalyear              []string `json:"originalyear,omitempty"`
	Releasecountry            []string `json:"releasecountry,omitempty"`
	Releasestatus             []string `json:"releasestatus,omitempty"`
	Releasetype               []string `json:"releasetype,omitempty"`
	Samplerate                int64    `json:"samplerate"`
	Script                    []string `json:"script,omitempty"`
	TrackTotal                int64    `json:"track_total"`
	CompatibleBrands          []string `json:"compatible_brands,omitempty"`
	Encoder                   []string `json:"encoder,omitempty"`
	MajorBrand                []string `json:"major_brand,omitempty"`
	MinorVersion              []string `json:"minor_version,omitempty"`
	Performer                 []string `json:"performer,omitempty"`
	Releasetime               []string `json:"releasetime,omitempty"`
	SortAlbum                 []string `json:"sort_album,omitempty"`
	SortAlbumArtist           []string `json:"sort_album_artist,omitempty"`
	SortArtist                []string `json:"sort_artist,omitempty"`
	SortComposer              []string `json:"sort_composer,omitempty"`
	SortName                  []string `json:"sort_name,omitempty"`
	Upc                       []string `json:"upc,omitempty"`
}

type LogTrackRequest struct {
	Duration  int64  `json:"duration"`
	Source    string `json:"source"`
	Timestamp int64  `json:"timestamp"`
	Trackhash string `json:"trackhash"`
}
