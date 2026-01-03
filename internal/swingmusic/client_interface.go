package swingmusic

import (
	"net/http"

	"github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

// SwingMusicClient defines the interface for interacting with the Swing Music server.
type SwingMusicClient interface {
	// Login performs user authentication and returns the session cookie.
	Login(username, password string) (authCookie *http.Cookie, err error)

	// GetAuthed returns an authenticated client using the provided authentication cookie.
	GetAuthed(authCookie string) SwingMusicClientAuthed
}

// SwingMusicClientAuthed defines the interface for an authenticated Swing Music client.
type SwingMusicClientAuthed interface {
	// GetFolderContents retrieves the contents of the specified folder.
	FolderContents(folder string) (folders *models.Folders, err error)
}
