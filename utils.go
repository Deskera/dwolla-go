package dwolla

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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
	location   = "Location"
	xRequestId = "X-Request-Id"
)

type resp struct {
	Body   []byte
	Header *http.Header
}

func post(url string, header *Header, payload interface{}, token *Token) (*resp, *Raw, error) {
	var bodyReader io.Reader
	var bodyBytes []byte

	if payload != nil {
		var err error
		bodyBytes, err = json.Marshal(payload)
		if err != nil {
			return nil, nil, err
		}

		bodyReader = bytes.NewReader(bodyBytes)
	}

	var client http.Client
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/vnd.dwolla.v1.hal+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	if header != nil && header.ContentType != "" {
		req.Header.Add("Content-Type", header.ContentType)
	} else {
		req.Header.Set("Content-Type", "application/vnd.dwolla.v1.hal+json")
	}

	if header != nil && header.IdempotencyKey != "" {
		req.Header.Add("Idempotency-Key", header.IdempotencyKey)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	status := res.StatusCode

	raw := &Raw{
		Endpoint: url,
		Request:  string(bodyBytes),
		Response: string(body),
		Status:   status,
		XRequestId: res.Header.Get(xRequestId),
	}

	if status == 400 || status == 403 || status == 404 || status == 500 || status == 401 {
		return nil, raw, errors.New(string(body))
	}

	return &resp{
		Body:   body,
		Header: &res.Header,
	}, raw, nil
}

func get(url string, token *Token) (*resp, *Raw, error) {
	var client http.Client
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/vnd.dwolla.v1.hal+json")
	req.Header.Set("Content-Type", "application/vnd.dwolla.v1.hal+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	status := res.StatusCode

	raw := &Raw{
		Endpoint: url,
		Request:  "",
		Response: string(body),
		Status:   status,
	}

	if status == 400 || status == 403 || status == 404 || status == 500 || status == 401 {
		return nil, raw, errors.New(string(body))
	}

	return &resp{
		Body:   body,
		Header: &res.Header,
	}, raw, nil
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
	lastIDX := strings.LastIndex(location, "/")
	if lastIDX < 0 {
		return "", ErrNoID
	}

	return location[lastIDX+1:], nil
}

func makeDeleteRequest(url string, token *Token) (*http.Response, error) {
	accept := "application/vnd.dwolla.v1.hal+json"
	var client http.Client
	req, err := http.NewRequest(http.MethodDelete, url, nil)
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

func upload(url string, header *Header, payload []byte, token *Token) (*resp, *Raw, error) {

	var client http.Client
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/vnd.dwolla.v1.hal+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	req.Header.Add("Content-Type", header.ContentType)

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	status := res.StatusCode

	raw := &Raw{
		Endpoint: url,
		Request:  "",
		Response: res.Header.Get(location),
		Status:   status,
		XRequestId: res.Header.Get(xRequestId),
	}

	if status == 400 || status == 403 || status == 404 || status == 500 || status == 401 {
		return nil, raw, errors.New(string(body))
	}

	return &resp{
		Body:   body,
		Header: &res.Header,
	}, raw, nil
}