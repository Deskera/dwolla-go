package dwolla

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func makePostRequest(url, accept string, payload interface{}, token *Token) (*http.Response, error) {
	bytesArray := new(bytes.Buffer)
	if err := json.NewEncoder(bytesArray).Encode(payload); err != nil {
		log.Println(err)
		return nil, err
	}

	var client http.Client
	req, err := http.NewRequest(http.MethodPost, url, bytesArray)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", accept)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func makeGetRequest(url, accept string, payload interface{}, token *Token) (*http.Response, error) {
	var client http.Client
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", accept)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
