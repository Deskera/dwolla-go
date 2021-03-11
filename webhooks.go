package dwolla

import (
	"encoding/json"
	"time"
)

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

func (c *webhook) Create(endpoint string, secret string) (*WebhookSubscription, *Raw, error) {
	url := c.baseURL + "/webhook-subscriptions"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	var webhookSubscription WebhookSubscription
	subscription := WebhookSubscriptionRequest{
		URL:    endpoint,
		Secret: secret,
	}

	resp, raw, err := post(url, nil, subscription, token)
	if err != nil {
		return nil, raw, err
	}

	webhookLocation := resp.Header.Get(location)
	webhookID, err := ExtractIDFromLocation(webhookLocation)
	if err != nil {
		return nil, raw, err
	}

	webhookSubscription.ID = webhookID
	webhookSubscription.URL = endpoint
	webhookSubscription.Created = time.Now().String()

	return &webhookSubscription, raw, nil
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
