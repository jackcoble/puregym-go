package puregym

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/jackcoble/puregym-go/pkg/types"
)

const (
	AUTH_URL = "https://auth.puregym.com/connect/token"
)

// Client represents the PureGym API Client
type Client struct {
	Email string
	PIN   string

	AccessToken string
}

// Return a new Client
func NewClient(email string, pin string) *Client {
	return &Client{
		Email: email,
		PIN:   pin,
	}
}

// Authenticate against the PureGym API to return an Access Token
func (c *Client) Authenticate() (*types.AuthenticateResponse, error) {
	// Construct the URL Encoded Form request
	requestBody := types.AuthenticateRequest{
		Username:  c.Email,
		Password:  c.Email,
		GrantType: "password",
		Scope:     "pgcapi",
		ClientID:  "ro.client",
	}

	vals, _ := query.Values(requestBody)

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", AUTH_URL, bytes.NewBuffer([]byte(vals.Encode())))
	if err != nil {
		return nil, err
	}

	// Set the Request headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "PureGym/1523 CFNetwork/1312 Darwin/21.0.0")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Marshal response into JSON
	var authResponse types.AuthenticateResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return nil, err
	}

	return &authResponse, nil
}
