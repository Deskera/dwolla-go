package dwolla

import (
	"encoding/json"
)

type business struct {
	authHandler *auth
	baseURL     string
}

func (c *business) GetBusinessClassification() (*BusinessClassificationsResponse, *Raw, error) {
	url := c.baseURL + "/business-classifications"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var data BusinessClassificationsResponse
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, raw, err
	}

	return &data, raw, nil
}
