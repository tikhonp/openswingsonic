package swingmusic

// SwingMusicClient defines the interface for interacting with the Swing Music server.
type SwingMusicClient interface {
	// Login performs user authentication and returns the session cookie string.
	Login(username, password string) (cookieString string, err error)
}
