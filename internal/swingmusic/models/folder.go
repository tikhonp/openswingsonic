package models

type FolderRequest struct {
	Folder            string `json:"folder"`
	Start             int64  `json:"start"`
	Limit             int64  `json:"limit"`
	TracksOnly        bool   `json:"tracks_only"`
	Sorttracksby      string `json:"sorttracksby"`
	TracksortReverse  bool   `json:"tracksort_reverse"`
	Sortfoldersby     string `json:"sortfoldersby"`
	FoldersortReverse bool   `json:"foldersort_reverse"`
}

type Folders struct {
	Folders []Folder `json:"folders"`
	Path    string   `json:"path"`
	Total   int64    `json:"total"`
	Tracks  []Track  `json:"tracks"`
}

type Folder struct {
	IsSym      bool   `json:"is_sym"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Trackcount int64  `json:"trackcount"`
}

type Track struct {
	Album        string   `json:"album"`
	Albumartists []Artist `json:"albumartists"`
	Albumhash    string   `json:"albumhash"`
	Artisthashes []string `json:"artisthashes"`
	Artists      []Artist `json:"artists"`
	Bitrate      int64    `json:"bitrate"`
	Duration     int64    `json:"duration"`
	Explicit     bool     `json:"explicit"`
	Extra        Extra    `json:"extra"`
	Filepath     string   `json:"filepath"`
	Folder       string   `json:"folder"`
	Image        string   `json:"image"`
	IsFavorite   bool     `json:"is_favorite"`
	Title        string   `json:"title"`
	Trackhash    string   `json:"trackhash"`
	Weakhash     string   `json:"weakhash"`
	Disc         int64    `json:"disc"`
	Track        int64    `json:"track"`
}

type Artist struct {
	Artisthash string `json:"artisthash"`
	Name       string `json:"name"`
}

type Extra struct {
	Artist     []string `json:"artist"`
	Bitdepth   int64    `json:"bitdepth"`
	Channels   int64    `json:"channels"`
	Composer   []string `json:"composer"`
	Filesize   int64    `json:"filesize"`
	Genre      []string `json:"genre"`
	Hashinfo   Hashinfo `json:"hashinfo"`
	Label      []string `json:"label"`
	Samplerate int64    `json:"samplerate"`
	TrackTotal int64    `json:"track_total"`
}

type Hashinfo struct {
	Algo   string `json:"algo"`
	Format string `json:"format"`
}

type DirBrowserResponse struct {
	Folders []DirBrowserItem `json:"folders"`
}

type DirBrowserItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type FolderTrackResponse struct {
	Tracks []Track `json:"tracks"`
}
