package puregym

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/jackcoble/puregym-go/pkg/types"
)

const (
	AUTH_URL = "https://auth.puregym.com/connect/token"
)

// Client represents the PureGym API Client
type Client struct {
	email string
	pin   string

	accessToken string
}

// Return a new Client
func NewClient(email string, pin string) *Client {
	return &Client{
		email: email,
		pin:   pin,
	}
}

// Authenticate against the PureGym API to return an Access Token
func (c *Client) Authenticate() error {
	// Construct the URL Encoded Form request
	requestBody := types.AuthenticateRequest{
		Username:  c.email,
		Password:  c.pin,
		GrantType: "password",
		Scope:     "pgcapi",
		ClientID:  "ro.client",
	}

	vals, err := query.Values(requestBody)
	if err != nil {
		return err
	}

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", AUTH_URL, bytes.NewBuffer([]byte(vals.Encode())))
	if err != nil {
		return err
	}

	// Set the Request headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "PureGym/1523 CFNetwork/1312 Darwin/21.0.0")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Marshal response into JSON
	var authResponse types.AuthenticateResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return err
	}

	// Extract and set the Access Token for future requests
	c.accessToken = authResponse.AccessToken

	return nil
}

// Returns the access token that was set during authentication.
// The client will error if no access token is set.
func (c *Client) GetAccessToken() (string, error) {
	// Check if access token is set
	if len(strings.TrimSpace(c.accessToken)) == 0 {
		return "", errors.New("no access token set")
	}

	return c.accessToken, nil
}
