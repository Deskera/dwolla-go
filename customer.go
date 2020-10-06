package dwolla

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

type customer struct {
	authHandler *auth
	baseURL     string
}

func CustomerHandler(customerConfig *customer) *customer {
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

func (c *customer) GetCustomers() (*CustomersResponse, error) {
	url := c.baseURL + "/customers"
	accept := "application/vnd.dwolla.v1.hal+json"

	token, err := c.authHandler.GetToken()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := makeGetRequest(url, accept, nil, token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var customersResponse CustomersResponse
	if err := json.NewDecoder(resp.Body).Decode(&customersResponse); err != nil {
		return nil, err
	}

	return &customersResponse, nil

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
