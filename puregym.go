package puregym

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/jackcoble/puregym-go/pkg/types"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/oauth2"
)

const (
	AUTH_URL       = "https://auth.puregym.com/connect/token"
	GYM_API_URL    = "https://capi.puregym.com/api/v1/gyms"
	MEMBER_API_URL = "https://capi.puregym.com/api/v1/member"
)

// Client represents the PureGym API Client
type Client struct {
	email string
	pin   string

	accessToken string
	homeGymId   int
}

// Return a new Client
func NewClient(email string, pin string) (*Client, error) {
	// Ensure the Email and PIN are not empty
	if len(strings.TrimSpace(email)) == 0 {
		return nil, errors.New("email address not provided")
	}

	if len(strings.TrimSpace(pin)) == 0 {
		return nil, errors.New("pin not provided")
	}

	return &Client{
		email: email,
		pin:   pin,
	}, nil
}

// Authenticate against the PureGym API via OAuth2 to return an Access Token
func (c *Client) Authenticate(ctx context.Context) error {
	authEndpoint := oauth2.Endpoint{
		TokenURL: AUTH_URL,
	}

	conf := &oauth2.Config{
		ClientID: "ro.client",
		Scopes:   []string{"pgcapi"},
		Endpoint: authEndpoint,
	}

	token, err := conf.PasswordCredentialsToken(ctx, c.email, c.pin)
	if err != nil {
		return err
	}

	// Extract and set the Access Token for future requests
	c.accessToken = token.AccessToken

	return nil
}

// Set an access token. If a token is already obtained, it can be re-used here.
func (c *Client) SetAccessToken(token string) error {
	// Parse the provided token into a JWT format
	verifyOpt := jwt.WithVerify(false)
	validateOpt := jwt.WithValidate(false)

	_, err := jwt.ParseString(token, verifyOpt, validateOpt)
	if err != nil {
		return errors.New("unable to parse provided token: " + err.Error())
	}

	// Set the token for the client
	c.accessToken = token
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

// Get member information
func (c *Client) GetMemberInfo() (*types.MemberResponse, error) {
	// Create the GET request
	req, err := http.NewRequest("GET", MEMBER_API_URL, nil)
	if err != nil {
		return nil, err
	}

	// Set the Access token
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Marshal response into JSON
	var memberInfoResponse types.MemberResponse
	if err := json.NewDecoder(resp.Body).Decode(&memberInfoResponse); err != nil {
		return nil, err
	}

	return &memberInfoResponse, nil
}

// Set the home gym in the Client for the user
func (c *Client) SetHomeGym() error {
	memberInfo, err := c.GetMemberInfo()
	if err != nil {
		return err
	}

	c.homeGymId = memberInfo.HomeGymID
	return nil
}

// Get the attendance information for the users Home Gym, or any ID provided
func (c *Client) GetGymAttendance(gymId ...int) (*types.GymAttendanceResponse, error) {
	var gym int

	// If gym IDs provided is zero, use the Home Gym
	if len(gymId) == 0 {
		// Check home gym is set
		if c.homeGymId == 0 {
			return nil, errors.New("home gym is not set within client")
		}
		gym = c.homeGymId
	} else {
		// Use the first provided Gym ID
		gym = gymId[0]
	}

	// Construct the request URL for the Gym
	requestUrl := fmt.Sprintf("%s/%d/attendance", GYM_API_URL, gym)

	// Create the GET request
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	// Set the Access token
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Marshal response into JSON
	var gymAttendanceResponse types.GymAttendanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&gymAttendanceResponse); err != nil {
		return nil, err
	}

	return &gymAttendanceResponse, nil
}
