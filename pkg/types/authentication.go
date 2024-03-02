package types

// Request payload type for authenticating against the API
type AuthenticateRequest struct {
	Username  string `url:"username"`
	Password  string `url:"password"`
	GrantType string `url:"grant_type"`
	Scope     string `url:"scope"`
	ClientID  string `url:"client_id"`
}

// Response returned from successful authentication against the API
type AuthenticateResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}
