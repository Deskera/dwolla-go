package dwolla

import (
	"encoding/json"
)

type transfer struct {
	authHandler *auth
	baseURL     string
}

func (c *transfer) GetTransfersById(transferId string) (*TransferResponse, *Raw, error) {
	url := c.baseURL + "/transfers/" + transferId

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var transferResponse TransferResponse
	if err := json.Unmarshal(resp.Body, &transferResponse); err != nil {
		return nil, raw, err
	}

	return &transferResponse, raw, nil
}
