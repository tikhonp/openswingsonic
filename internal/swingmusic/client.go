// Package swingmusic provides a client for interacting with the Swing Music music server.
package swingmusic

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"slices"

	"github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

// swingMusicClient implements the SwingMusicClient interface.
type swingMusicClient struct {
	baseURL string
}

func NewClient(baseURL string) SwingMusicClient {
	return &swingMusicClient{baseURL: baseURL}
}

func (c *swingMusicClient) Login(username, password string) (authCookie *http.Cookie, err error) {
	url := c.baseURL + "/auth/login"

	jsonData, err := json.Marshal(models.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error making POST request:", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("login failed with status: " + resp.Status)
	}

	cookies := resp.Cookies()
	indxOfAuthCookie := slices.IndexFunc(cookies, func(c *http.Cookie) bool {
		return c.Name == "access_token_cookie"
	})
	if indxOfAuthCookie == -1 {
		return nil, errors.New("authentication cookie not found in response")
	}
	return cookies[indxOfAuthCookie], nil
}

type swingMusicClientAuthed struct {
	*swingMusicClient
	authCookie *http.Cookie
}

func (c *swingMusicClient) GetAuthed(authCookie string) SwingMusicClientAuthed {
	return &swingMusicClientAuthed{
		swingMusicClient: c,
		authCookie: &http.Cookie{
			Name:  "access_token_cookie",
			Value: authCookie,
		},
	}
}

// FolderContents retrieves the contents of the specified folder.
// note "$home" can be used to refer the root music folder.
func (c *swingMusicClientAuthed) FolderContents(folder string) (folders *models.Folders, err error) {
	url := c.baseURL + "/folder"

	jsonData, err := json.Marshal(models.FolderRequest{
		Folder:            folder,
		FoldersortReverse: false,
		Limit:             0,
		Sortfoldersby:     "lastmod",
		Sorttracksby:      "default",
		Start:             0,
		TracksOnly:        false,
		TracksortReverse:  false,
	})
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		return nil, err
	}
	request.AddCookie(c.authCookie)
	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("Error making POST request:", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("folder contents request failed with status: " + resp.Status)
	}

	var folderResponse models.Folders
	err = json.NewDecoder(resp.Body).Decode(&folderResponse)
	if err != nil {
		log.Println("Error decoding JSON response:", err)
		return nil, err
	}

	return &folderResponse, nil
}
