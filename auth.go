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
	clientId     string
	clientSecret string
	baseURL      string
}

func AuthHandler(authConfig *auth) *auth {
	return authConfig
}

var token Token

func (a *auth) FetchToken() (*Token, error) {
	var client http.Client
	url := a.baseURL + "/token"
	payload := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	data := a.clientId + ":" + a.clientSecret

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	req.Header.Add("Authorization", "Basic "+sEnc)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&token); err != nil {
		log.Println(err)
		return nil, err
	}

	token.CreatedAt = time.Now()
	return &token, nil
}

func (a *auth) GetToken() (*Token, error) {
	if &token != nil && isValid(&token) {
		log.Println("Token is Valid")
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
