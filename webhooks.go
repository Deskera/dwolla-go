package dwolla

import "encoding/json"

type webhook struct {
	authHandler *auth
	baseURL     string
}

func (c *webhook) List() (*WebhookSubscriptionsResponse, error) {
	url := c.baseURL + "/webhook-subscriptions"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := get(url, token)
	if err != nil {
		return nil, err
	}

	var data WebhookSubscriptionsResponse
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *webhook) Create(endpoint string, secret string) error {
	url := c.baseURL + "/webhook-subscriptions"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return err
	}
	subscription := WebhookSubscriptionRequest{
		URL:    endpoint,
		Secret: secret,
	}

	_, err = post(url, nil, subscription, token)
	if err != nil {
		return err
	}

	return nil
}

func (c *webhook) Delete(id string) error {
	url := c.baseURL + "/webhook-subscriptions/" + id

	token, err := c.authHandler.GetToken()
	if err != nil {
		return err
	}

	_, err = makeDeleteRequest(url, token)
	if err != nil {
		return err
	}

	return nil
}
