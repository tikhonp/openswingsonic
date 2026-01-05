// Package swingmusic provides a client for interacting with the Swing Music music server.
package swingmusic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strconv"

	"github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

var (
	ErrNotFound = errors.New("requested resource not found")
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

func (c *swingMusicClient) GetAlbumImageURL(albumHash string, size ImageSize) string {
	if size == ImageSizeLarge {
		return fmt.Sprintf("%s/img/thumbnail/%s.webp", c.baseURL, albumHash)
	}
	return fmt.Sprintf("%s/img/thumbnail/%s/%s.webp", c.baseURL, size, albumHash)
}

func (c *swingMusicClient) GetArtistImageURL(artistHash string, size ImageSize) string {
	if size == ImageSizeLarge {
		return fmt.Sprintf("%s/img/artist/%s.webp", c.baseURL, artistHash)
	}
	return fmt.Sprintf("%s/img/artist/%s/%s.webp", c.baseURL, size, artistHash)
}

func (c *swingMusicClient) GetThumbnailByID(thumbnailID string) (string, io.ReadCloser, error) {
	url := fmt.Sprintf("%s/img/thumbnail/%s", c.baseURL, thumbnailID)
	resp, err := http.Get(url)
	if err != nil {
		return "", nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("failed to get thumbnail, status: %s", resp.Status)
	}
	return resp.Header.Get("Content-Type"), resp.Body, nil
}

// swingMusicClientAuthed implements the SwingMusicClientAuthed interface.
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
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		return nil, err
	}

	request.AddCookie(c.authCookie)
	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, err
	}
	if resp.StatusCode >= 400 {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println("Error decoding JSON response:", err)
		return nil, err
	}
	return &response, nil
}

func (c *swingMusicClientAuthed) Album(albumHash string, albumLimit int) (*models.AlbumResponse, error) {
	url := c.baseURL + "/album"
	reqBody := models.AlbumRequest{
		AlbumHash:  albumHash,
		AlbumLimit: albumLimit,
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	return doRequest[models.AlbumResponse](c, http.MethodPost, url, bytes.NewBuffer(jsonData))
}

func (c *swingMusicClientAuthed) AlbumOtherVersions(albumHash, ogAlbumTitle string) (*[]models.AlbumShortInfo, error) {
	url := c.baseURL + "/album/other-versions"
	jsonData, err := json.Marshal(models.AlbumOtherVersionsRequest{
		AlbumHash:    albumHash,
		OgAlbumTitle: ogAlbumTitle,
	})
	if err != nil {
		return nil, err
	}
	return doRequest[[]models.AlbumShortInfo](c, http.MethodPost, url, bytes.NewBuffer(jsonData))
}

func (c *swingMusicClientAuthed) AlbumTracks(albumHash string) (*[]models.Track, error) {
	url := c.baseURL + "/album/" + albumHash + "/tracks"
	return doRequest[[]models.Track](c, http.MethodGet, url, nil)
}

func (c *swingMusicClientAuthed) Artist(artistHash string) (*models.ArtistResponse, error) {
	u, err := url.Parse(c.baseURL + "/artist/" + artistHash)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	// query.Set("tracklimit", strconv.Itoa(trackLimit))
	// query.Set("albumlimit", strconv.Itoa(albumLimit))
	query.Set("all", "true")
	u.RawQuery = query.Encode()

	return doRequest[models.ArtistResponse](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) ArtistAlbums(artistHash string) (*models.ArtistAlbumsResponse, error) {
	u, err := url.Parse(c.baseURL + "/artist/" + artistHash + "/albums")
	if err != nil {
		return nil, err
	}
	query := u.Query()
	// query.Set("albumlimit", strconv.Itoa(albumLimit))
	query.Set("all", "true")
	u.RawQuery = query.Encode()

	return doRequest[models.ArtistAlbumsResponse](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) ArtistTracks(artistHash string) (*[]models.Track, error) {
	url := c.baseURL + "/artist/" + artistHash + "/tracks"
	return doRequest[[]models.Track](c, http.MethodGet, url, nil)
}

func (c *swingMusicClientAuthed) FolderContents(folder string) (*models.Folders, error) {
	url := c.baseURL + "/folder"
	jsonData, err := json.Marshal(models.FolderRequest{
		Folder:            folder,
		FoldersortReverse: false,
		Limit:             1000000, // large number to get all folders
		Sortfoldersby:     "name",
		Sorttracksby:      "default",
		Start:             0,
		TracksOnly:        false,
		TracksortReverse:  false,
	})
	if err != nil {
		return nil, err
	}
	return doRequest[models.Folders](c, http.MethodPost, url, bytes.NewBuffer(jsonData))
}

func (c *swingMusicClientAuthed) FolderDirBrowser(folder string) (*models.DirBrowserResponse, error) {
	url := c.baseURL + "/folder/dir-browser"
	reqBody := map[string]string{"folder": folder}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	return doRequest[models.DirBrowserResponse](c, http.MethodPost, url, bytes.NewBuffer(jsonData))
}

func (c *swingMusicClientAuthed) FolderTracksAll(path string) (*models.FolderTrackResponse, error) {
	u, err := url.Parse(c.baseURL + "/folder/tracks/all")
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("path", path)
	u.RawQuery = query.Encode()
	return doRequest[models.FolderTrackResponse](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) AllArtists(sortBy string, reverse bool) (*models.Artists, error) {
	u, err := url.Parse(c.baseURL + "/getall/artists")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("start", "0")
	q.Set("limit", "1000000") // large number to get all artists
	q.Set("sortby", sortBy)
	if reverse {
		q.Set("reverse", "1")
	} else {
		q.Set("reverse", "0")
	}
	u.RawQuery = q.Encode()

	return doRequest[models.Artists](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) AllAlbums(sortBy string, reverse bool) (*models.Albums, error) {
	u, err := url.Parse(c.baseURL + "/getall/albums")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("start", "0")
	q.Set("limit", "1000000") // large number to get all tracks
	q.Set("sortby", sortBy)
	if reverse {
		q.Set("reverse", "1")
	} else {
		q.Set("reverse", "0")
	}
	u.RawQuery = q.Encode()

	return doRequest[models.Albums](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) SearchTracks(query string, limit int) (*models.SearchedTracks, error) {
	u, err := url.Parse(c.baseURL + "/search/")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("q", query)
	q.Set("start", "0")
	q.Set("itemtype", "tracks")
	u.RawQuery = q.Encode()

	return doRequest[models.SearchedTracks](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) SearchAlbums(query string, limit int) (*models.SearchedAlbums, error) {
	u, err := url.Parse(c.baseURL + "/search/")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("q", query)
	q.Set("start", "0")
	q.Set("itemtype", "albums")
	u.RawQuery = q.Encode()

	return doRequest[models.SearchedAlbums](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) SearchArtists(query string, limit int) (*models.SearchedArtists, error) {
	u, err := url.Parse(c.baseURL + "/search/")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("q", query)
	q.Set("start", "0")
	q.Set("itemtype", "artists")
	u.RawQuery = q.Encode()

	return doRequest[models.SearchedArtists](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) SearchAll(query string, limit int) (*models.SearchedAll, error) {
	u, err := url.Parse(c.baseURL + "/search/top")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("limit", strconv.Itoa(limit))
	q.Set("q", query)
	q.Set("start", "0")
	u.RawQuery = q.Encode()

	return doRequest[models.SearchedAll](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) Stats() (*models.Stats, error) {
	url := c.baseURL + "/logger/stats"
	return doRequest[models.Stats](c, http.MethodGet, url, nil)
}

func (c *swingMusicClientAuthed) TopTracks(duration, orderBy string, limit int) (*models.TopTracks, error) {
	u, err := url.Parse(c.baseURL + "/logger/top-tracks")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("duration", duration)
	q.Set("orderBy", orderBy)
	q.Set("limit", strconv.Itoa(limit))
	u.RawQuery = q.Encode()

	return doRequest[models.TopTracks](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) TopAlbums(duration, orderBy string, limit int) (*models.TopAlbums, error) {
	u, err := url.Parse(c.baseURL + "/logger/top-albums")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("duration", duration)
	q.Set("orderBy", orderBy)
	q.Set("limit", strconv.Itoa(limit))
	u.RawQuery = q.Encode()
	return doRequest[models.TopAlbums](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) TopArtists(duration, orderBy string, limit int) (*models.TopArtists, error) {
	u, err := url.Parse(c.baseURL + "/logger/top-artists")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("duration", duration)
	q.Set("orderBy", orderBy)
	q.Set("limit", strconv.Itoa(limit))
	u.RawQuery = q.Encode()

	return doRequest[models.TopArtists](c, http.MethodGet, u.String(), nil)
}

func (c *swingMusicClientAuthed) LogTrack(req *models.LogTrackRequest) error {
	url := c.baseURL + "/logger/track/log"

	jsonData, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, err = doRequest[any](c, http.MethodPost, url, bytes.NewBuffer(jsonData))
	return err
}

func (c *swingMusicClientAuthed) Stream(trackhash, filepath, rangeHeader string) (*models.StreamedFileHeaders, io.ReadCloser, error) {
	u, err := url.Parse(fmt.Sprintf("%s/file/%s/legacy", c.baseURL, trackhash))
	if err != nil {
		return nil, nil, err
	}
	q := u.Query()
	q.Set("filepath", filepath)
	q.Set("container", "flac")
	q.Set("quality", "original")
	u.RawQuery = q.Encode()

	request, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	request.AddCookie(c.authCookie)
	if rangeHeader != "" {
		request.Header.Set("Range", rangeHeader)
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode >= 400 {
		if resp.StatusCode == http.StatusNotFound {
			return nil, nil, ErrNotFound
		}
		return nil, nil, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}
	headers := &models.StreamedFileHeaders{
		ContentType:        resp.Header.Get("Content-Type"),
		ContentDisposition: resp.Header.Get("Content-Disposition"),
		ContentRange:       resp.Header.Get("Content-Range"),
	}
	if resp.ContentLength != -1 {
		headers.ContentLength = int(resp.ContentLength)
	}
	return headers, resp.Body, nil
}
