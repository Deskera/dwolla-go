package dwolla

import (
	"encoding/json"
	"errors"

	"github.com/jinzhu/copier"
)

// CustomerType is the customer's type
type CustomerType string

const (
	// CustomerTypeBusiness is when the customer is a business
	CustomerTypeBusiness CustomerType = "business"
	// CustomerTypePersonal is when the customer is an individual
	CustomerTypePersonal CustomerType = "personal"
	// CustomerTypeReceiveOnly is when the customer can only receive funds
	CustomerTypeReceiveOnly CustomerType = "receive-only"
	// CustomerTypeUnverified is when the customer is unverified
	CustomerTypeUnverified CustomerType = "unverified"
)

// BusinessType is the type of business setup
type BusinessType string

const (
	LLC                BusinessType = "llc"
	Patnership         BusinessType = "partnership"
	Corporation        BusinessType = "corporation"
	SoleProprietorship BusinessType = "soleProprietorship"
)

// CustomerStatus is the customer's status
type CustomerStatus string

const (
	// CustomerStatusDeactivated is when the customer has been deactivated
	CustomerStatusDeactivated CustomerStatus = "deactivated"
	// CustomerStatusDocument is when the customer needs verification document
	CustomerStatusDocument CustomerStatus = "document"
	// CustomerStatusReactivated is when a deactivated customer is reactivated
	CustomerStatusReactivated CustomerStatus = "reactivated"
	// CustomerStatusRetry is when the customer needs to retry verification
	CustomerStatusRetry CustomerStatus = "retry"
	// CustomerStatusSuspended is when the customer has been suspended
	CustomerStatusSuspended CustomerStatus = "suspended"
	// CustomerStatusUnverified is when the customer is unverified
	CustomerStatusUnverified CustomerStatus = "unverified"
	// CustomerStatusVerified is when the customer is verified
	CustomerStatusVerified CustomerStatus = "verified"
)

type customer struct {
	authHandler *auth
	baseURL     string
}

//CustomerHandler generates the customer handler on client initialisation.
func CustomerHandler(customerConfig *customer) *customer {
	return customerConfig
}

func (c *customer) CreateCustomer(customer *CustomerRequest) (*Customer, error) {
	url := c.baseURL + "/customers"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	var customerResp Customer
	if err := copier.Copy(&customerResp, &customer); err != nil {
		return nil, err
	}

	resp, err := post(url, nil, customer, token)
	if err != nil {
		return nil, err
	}

	customerLocation := resp.Header.Get(location)
	customerID, err := ExtractIDFromLocation(customerLocation)
	if err != nil {
		return nil, err
	}

	customerResp.Location = customerLocation
	customerResp.ID = customerID
	customerResp.RawResponse = string(resp.Body)
	customerResp.Created = true

	return &customerResp, nil
}

func (c *customer) AddFundingSourceForCustomerPlaid(plaidToken, customerID, fundingSourceName string) (string, error) {
	url := c.baseURL + "/customers/" + customerID + "/funding-sources"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return "", err
	}

	fundingSourceReq := &PlaidFundingSourceRequest{
		PlaidToken: plaidToken,
		Name:       fundingSourceName,
	}

	resp, err := post(url, nil, fundingSourceReq, token)
	if err != nil {
		return "", err
	}

	fundingSourceLink := resp.Header.Get(location)

	return fundingSourceLink, nil
}

func (c *customer) GetFundingSourcesForCustomer(customerID string) (*FundingSourcesResponse, error) {
	url := c.baseURL + "/customers/" + customerID + "/funding-sources"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := get(url, token)
	if err != nil {
		return nil, err
	}

	var fundingSourceResp FundingSourcesResponse
	if err := json.Unmarshal(resp.Body, &fundingSourceResp); err != nil {
		return nil, err
	}

	return &fundingSourceResp, nil
}

func (c *customer) AddFundingSourceForCustomer(customerID string, fundingSourceReq *FundingSourceRequest) (string, error) {
	url := c.baseURL + "/customers/" + customerID + "/funding-sources"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return "", err
	}

	resp, err := post(url, nil, fundingSourceReq, token)
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

	resp, err := get(url, token)
	if err != nil {
		return nil, err
	}

	var customersResponse CustomersResponse
	if err := json.Unmarshal(resp.Body, &customersResponse); err != nil {
		return nil, err
	}

	return &customersResponse, nil
}

func (c *customer) CutomerErrorHandler(errMsg error) (string, error) {
	errorArr, err := parseError(errMsg.Error())
	if err != nil {
		return "", err
	}

	var errorMessage string

	for _, dwollaError := range errorArr {
		if dwollaError.Code == "Duplicate" {
			switch dwollaError.Path {
			case "/correlationId":
				errorMessage = "duplicate_correlationId"
			case "/email":
				errorMessage = "duplicate_email"
			}

			return dwollaError.Links.About.Href, errors.New(errorMessage)
		}
	}

	return "", errMsg
}
