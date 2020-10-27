package dwolla

import (
	"encoding/json"
	"log"
)

type businessClassifications struct {
	authHandler *auth
	baseURL     string
}

func (c *businessClassifications) Get() (*BusinessClassificationsResponse, error) {
	url := c.baseURL + "/business-classifications"

	token, err := c.authHandler.GetToken()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := makeGetRequest(url, token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var data BusinessClassificationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
