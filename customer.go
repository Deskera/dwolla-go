package dwolla

import (
	"errors"
	"log"
	"strings"
)

type customer struct {
	authHandler *auth
	baseURL     string
}

func Customer(customerConfig *customer) *customer {
	return customerConfig
}

func (c *customer) CreateUnverifiedCostumer(unverifiedCostumer UnverifiedCustomer) (*UnverifiedCustomer, error) {
	url := c.baseURL + "/customers"
	accept := "application/vnd.dwolla.v1.hal+json"

	token, err := c.authHandler.GetToken()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := makePostRequest(url, accept, unverifiedCostumer, token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	customerId := resp.Header.Get("Location")
	unverifiedCostumer.CustomerId = customerId

	return &unverifiedCostumer, nil
}

func (c *customer) CreateReceiveOnlyCostumer(receiveOnlyCostumer ReceiveOnlyCustomer) (*ReceiveOnlyCustomer, error) {
	url := c.baseURL + "/customers"
	accept := "application/vnd.dwolla.v1.hal+json"

	token, err := c.authHandler.GetToken()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := makePostRequest(url, accept, receiveOnlyCostumer, token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	customerId := resp.Header.Get("Location")
	receiveOnlyCostumer.CustomerId = customerId

	return &receiveOnlyCostumer, nil
}

func getCustomerId(location string) (string, error) {
	if location == "" {
		return "", errors.New("no location")
	}

	locationSplit := strings.Split(location, "/customers/")
	if len(locationSplit) < 1 {
		return "", errors.New("error extraction id")
	}

	customerId := locationSplit[1]
	return customerId, nil
}
