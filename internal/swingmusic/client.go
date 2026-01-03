// Package swingmusic provides a client for interacting with the Swing Music music server.
package swingmusic

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
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

// doRequest is a helper function to perform HTTP requests and parse JSON responses.
func doRequest[Response any](c *swingMusicClientAuthed, method, url string, body io.Reader) (*Response, error) {
	// Create HTTP request
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		return nil, err
	}

	// Add authentication cookie
	request.AddCookie(c.authCookie)

	// Set headers and perform the request
	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP request failed with status: " + resp.Status)
	}

	// Decode JSON response
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("Error decoding JSON response:", err)
		return nil, err
	}
	return &response, nil
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

	return doRequest[models.Folders](c, http.MethodPost, url, bytes.NewBuffer(jsonData))
}

func (c *swingMusicClientAuthed) AllArtists() (artists *models.Artists, err error) {
	u, err := url.Parse(c.baseURL + "/getall/artists")
	if err != nil {
		log.Println("Error parsing URL:", err)
		return nil, err
	}
	query := u.Query()
	query.Set("start", "0")
	query.Set("limit", "10000")
	query.Set("sortby", "created_date")
	query.Set("reverse", "1")
	u.RawQuery = query.Encode()

	return doRequest[models.Artists](c, http.MethodGet, u.String(), nil)
}
