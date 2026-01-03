// Package swingmusic provides a client for interacting with the Swing Music music server.
package swingmusic

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// swingMusicClient implements the SwingMusicClient interface.
type swingMusicClient struct {
	baseURL string
}

func NewClient(baseURL string) SwingMusicClient {
	return &swingMusicClient{baseURL: baseURL}
}

func (c *swingMusicClient) Login(username, password string) (cookieString string, err error) {
	url := c.baseURL + "/auth/login"

	jsonData, err := json.Marshal(LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error making POST request:", err)
		return "", err
	}

	cokkies := resp.Header.Get("Set-Cookie")
	println("Received cookies:", cokkies)

	return cokkies, nil
}
