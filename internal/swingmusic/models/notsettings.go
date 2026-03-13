package models

type NotSettings struct {
	Finished                  bool                `json:"_finished"`
	ArtistSeparators          []string            `json:"artistSeparators"`
	ArtistSplitIgnoreList     []string            `json:"artistSplitIgnoreList"`
	CleanAlbumTitle           bool                `json:"cleanAlbumTitle"`
	EnablePeriodicScans       bool                `json:"enablePeriodicScans"`
	EnablePlugins             bool                `json:"enablePlugins"`
	EnableWatchdog            bool                `json:"enableWatchdog"`
	ExcludeDirs               []string            `json:"excludeDirs"`
	ExtractFeaturedArtists    bool                `json:"extractFeaturedArtists"`
	GenreSeparators           []string            `json:"genreSeparators"`
	LastfmAPIKey              string              `json:"lastfmApiKey"`
	LastfmAPISecret           string              `json:"lastfmApiSecret"`
	LastfmSessionKey          string              `json:"lastfmSessionKey"`
	MergeAlbums               bool                `json:"mergeAlbums"`
	Plugins                   []NotSettingsPlugin `json:"plugins"`
	RemoveProdBy              bool                `json:"removeProdBy"`
	RemoveRemasterInfo        bool                `json:"removeRemasterInfo"`
	RootDirs                  []string            `json:"rootDirs"`
	ScanInterval              int                 `json:"scanInterval"`
	ServerID                  string              `json:"serverId"`
	ShowAlbumsAsSingles       bool                `json:"showAlbumsAsSingles"`
	ShowPlaylistsInFolderView bool                `json:"showPlaylistsInFolderView"`
	UsersOnLogin              bool                `json:"usersOnLogin"`
	Version                   string              `json:"version"`
}

type NotSettingsPlugin struct {
	Active   bool                     `json:"active"`
	Extra    NotSettingsPluginExtra   `json:"extra"`
	Name     string                   `json:"name"`
	Settings NotSettingsPluginSetting `json:"settings"`
}

type NotSettingsPluginExtra struct {
	Description string `json:"description"`
}

type NotSettingsPluginSetting struct {
	AutoDownload    bool `json:"auto_download"`
	OverideUnsynced bool `json:"overide_unsynced"`
}
