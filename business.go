package dwolla

import (
	"encoding/json"
)

type business struct {
	authHandler *auth
	baseURL     string
}

func (c *business) GetBusinessClassification() (*BusinessClassificationsResponse, error) {
	url := c.baseURL + "/business-classifications"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := get(url, token)
	if err != nil {
		return nil, err
	}

	var data BusinessClassificationsResponse
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
