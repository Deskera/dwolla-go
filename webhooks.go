package dwolla

import "encoding/json"

type webhook struct {
	authHandler *auth
	baseURL     string
}

func (c *webhook) List() (*WebhookSubscriptionsResponse, *Raw, error) {
	url := c.baseURL + "/webhook-subscriptions"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var data WebhookSubscriptionsResponse
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, raw, err
	}

	return &data, raw, nil
}

func (c *webhook) Create(endpoint string, secret string) (*Raw, error) {
	url := c.baseURL + "/webhook-subscriptions"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}
	subscription := WebhookSubscriptionRequest{
		URL:    endpoint,
		Secret: secret,
	}

	_, raw, err := post(url, nil, subscription, token)
	if err != nil {
		return raw, err
	}

	return raw, nil
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
