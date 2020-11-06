package dwolla

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ErrNoID ...
var ErrNoID = errors.New("unable to extract ID")

// Currency represents the monetary currency
type Currency string

const (
	// USD is U.S. dollars
	USD Currency = "USD"
)
const (
	location = "Location"
)

type resp struct {
	Body   []byte
	Header *http.Header
}

func post(url string, header *Header, payload interface{}, token *Token) (*resp, error) {
	var bodyReader io.Reader

	if payload != nil {
		bodyBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewReader(bodyBytes)
	}

	var client http.Client
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.dwolla.v1.hal+json")
	req.Header.Set("Content-Type", "application/vnd.dwolla.v1.hal+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	if header != nil {
		req.Header.Add("Idempotency-Key", header.IdempotencyKey)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(body))

	status := res.StatusCode
	if status == 400 || status == 403 || status == 404 || status == 500 || status == 401 {
		return nil, errors.New(string(body))
	}

	return &resp{
		Body:   body,
		Header: &res.Header,
	}, nil
}

func get(url string, token *Token) (*resp, error) {
	var client http.Client
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.dwolla.v1.hal+json")
	req.Header.Set("Content-Type", "application/vnd.dwolla.v1.hal+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(body))

	status := res.StatusCode
	if status == 400 || status == 403 || status == 404 || status == 500 || status == 401 {
		return nil, errors.New(string(body))
	}

	return &resp{
		Body:   body,
		Header: &res.Header,
	}, nil
}

func parseError(err string) ([]HALError, error) {
	var result ValidationError

	if err := json.Unmarshal([]byte(err), &result); err != nil {
		return nil, err
	}

	return result.Embedded.Errors, nil
}

// ExtractIDFromLocation takes an HREF link and returns the ID at the end of the HREF.
// This is useful for processing webhooks where you have an HREF, but need
// to make calls using this SDK, which expects bare IDs.
//
// If the input HREF is malformed, or this function is unable to extract the ID,
// ErrNoID will be returned.
func ExtractIDFromLocation(location string) (string, error) {
	if location == "" {
		return "", errors.New("no location")
	}

	locationSplit := strings.Split(location, "/customers/")
	if len(locationSplit) < 1 {
		return "", errors.New("error extraction id")
	}

	return locationSplit[1], nil
}
