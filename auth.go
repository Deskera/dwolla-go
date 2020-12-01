package dwolla

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type auth struct {
	clientID     string
	clientSecret string
	baseURL      string
}

// Token is a dwolla auth token
type Token struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	ExpiresIn   float64 `json:"expires_in"`
	CreatedAt   time.Time
}

var token Token

func (a *auth) FetchToken() (*Token, error) {
	var client http.Client
	url := a.baseURL + "/token"
	payload := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}

	data := a.clientID + ":" + a.clientSecret

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	req.Header.Add("Authorization", "Basic "+sEnc)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}

	token.CreatedAt = time.Now()
	return &token, nil
}

func (a *auth) GetToken() (*Token, error) {
	if &token != nil && isValid(&token) {
		return &token, nil
	}

	log.Println("Invalid token, fetching new one")
	token, err := a.FetchToken()
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *auth) SetToken(tokenInput Token) {
	token = tokenInput
}

func isValid(token *Token) bool {
	timeSinceCreation := time.Since(token.CreatedAt)
	if timeSinceCreation.Seconds() > token.ExpiresIn {
		return false
	}

	return true
}

// Expired returns true if token has expired
func (t *Token) Expired() bool {
	return time.Since(t.CreatedAt) > time.Duration(t.ExpiresIn)*time.Second
}
