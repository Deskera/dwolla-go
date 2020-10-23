package dwolla

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func makePostRequest(url string, header *Header, payload interface{}, token *Token) (*http.Response, error) {
	accept := "application/vnd.dwolla.v1.hal+json"
	bytesArray := new(bytes.Buffer)
	if err := json.NewEncoder(bytesArray).Encode(payload); err != nil {
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

	if header != nil {
		req.Header.Add("Idempotency-Key", header.IdempotencyKey)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	status := res.StatusCode
	if status == 400 || status == 403 || status == 404 || status == 500 || status == 401 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, errors.New(string(body))
	}

	return res, nil
}

func makeGetRequest(url string, token *Token) (*http.Response, error) {
	accept := "application/vnd.dwolla.v1.hal+json"
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

	status := res.StatusCode
	if status == 400 || status == 403 || status == 404 || status == 500 || status == 401 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, errors.New(string(body))
	}

	return res, nil
}
