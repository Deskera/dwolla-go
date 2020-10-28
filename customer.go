package dwolla

import (
	"encoding/json"
	"errors"
	"strings"
)

type customer struct {
	authHandler *auth
	baseURL     string
}

func CustomerHandler(customerConfig *customer) *customer {
	return customerConfig
}

func (c *customer) CreateVerifiedCustomer(verifiedCustomer *VerifiedCustomer) (*VerifiedCustomer, error) {
	url := c.baseURL + "/customers"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := makePostRequest(url, nil, verifiedCustomer, token)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	customerLocation := resp.Header.Get(location)
	verifiedCustomer.CustomerLocation = customerLocation

	customerId, err := getCustomerId(customerLocation)
	if err != nil {
		return nil, err
	}

	verifiedCustomer.CustomerId = customerId

	return verifiedCustomer, nil
}

func (c *customer) CreateUnverifiedCustomer(unverifiedCustomer *UnverifiedCustomer) (*UnverifiedCustomer, error) {
	url := c.baseURL + "/customers"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := makePostRequest(url, nil, unverifiedCustomer, token)
	if err != nil {
		return nil, err
	}

	customerLocation := resp.Header.Get(location)
	unverifiedCustomer.CustomerLocation = customerLocation

	customerId, err := getCustomerId(customerLocation)
	if err != nil {
		return nil, err
	}

	unverifiedCustomer.CustomerId = customerId

	return unverifiedCustomer, nil
}

func (c *customer) CreateReceiveOnlyCustomer(receiveOnlyCustomer *ReceiveOnlyCustomer) (*ReceiveOnlyCustomer, error) {
	url := c.baseURL + "/customers"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := makePostRequest(url, nil, receiveOnlyCustomer, token)
	if err != nil {
		return nil, err
	}

	customerLocation := resp.Header.Get(location)
	receiveOnlyCustomer.CustomerLocation = customerLocation

	customerId, err := getCustomerId(customerLocation)
	if err != nil {
		return nil, err
	}

	receiveOnlyCustomer.CustomerId = customerId

	return receiveOnlyCustomer, nil
}

func (c *customer) AddFundingSourceForCustomerPlaid(plaidToken, customerId, fundingSourceName string) (string, error) {
	url := c.baseURL + "/customers/" + customerId + "/funding-sources"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return "", err
	}

	fundingSourceReq := &PlaidFundingSourceRequest{
		PlaidToken: plaidToken,
		Name:       fundingSourceName,
	}

	resp, err := makePostRequest(url, nil, fundingSourceReq, token)
	if err != nil {
		return "", err
	}

	fundingSourceLink := resp.Header.Get(location)

	return fundingSourceLink, nil
}

func (c *customer) GetFundingSourcesForCustomer(customerId string) (*FundingSourcesResponse, error) {
	url := c.baseURL + "/customers/" + customerId + "/funding-sources"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := makeGetRequest(url, token)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var fundingSourceResp FundingSourcesResponse
	if err := json.NewDecoder(resp.Body).Decode(&fundingSourceResp); err != nil {
		return nil, err
	}

	return &fundingSourceResp, nil
}

func (c *customer) AddFundingSourceForCustomer(customerId string, fundingSourceReq *FundingSourceRequest) (string, error) {
	url := c.baseURL + "/customers/" + customerId + "/funding-sources"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return "", err
	}

	resp, err := makePostRequest(url, nil, fundingSourceReq, token)
	if err != nil {
		return "", err
	}

	fundingSourceLocation := resp.Header.Get(location)

	return fundingSourceLocation, nil
}

func (c *customer) GetCustomers() (*CustomersResponse, error) {
	url := c.baseURL + "/customers"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := makeGetRequest(url, token)
	if err != nil {
		return nil, err
	}

	var customersResponse CustomersResponse
	if err := json.NewDecoder(resp.Body).Decode(&customersResponse); err != nil {
		return nil, err
	}

	return &customersResponse, nil
}

// func (c *customer) GetCustomerById(customerLocation string)

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
