package dwolla

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/jinzhu/copier"
	"io"
	"log"
	"mime/multipart"
	"strings"
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

func (c *customer) CreateCustomer(customer *CustomerRequest) (*Customer, *Raw, error) {
	url := c.baseURL + "/customers"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	var customerResp Customer
	if err := copier.Copy(&customerResp, &customer); err != nil {
		return nil, nil, err
	}

	resp, raw, err := post(url, nil, customer, token)
	if err != nil {
		return &customerResp, raw, err
	}

	customerLocation := resp.Header.Get(location)
	customerID, err := ExtractIDFromLocation(customerLocation)
	if err != nil {
		return nil, raw, err
	}

	customerResp.Location = customerLocation
	customerResp.ID = customerID

	return &customerResp, raw, nil
}

func (c *customer) UpdateCustomer(verifiedCustomerID string, customer *CustomerRequest) (*Customer, *Raw, error) {
	url := c.baseURL + "/customers/" + verifiedCustomerID

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	var customerResp Customer
	if err := copier.Copy(&customerResp, &customer); err != nil {
		return nil, nil, err
	}

	_, raw, err := post(url, nil, customer, token)
	if err != nil {
		return &customerResp, raw, err
	}

	return &customerResp, raw, nil
}

func (c *customer) AddFundingSourceForCustomerPlaid(plaidToken, customerID, fundingSourceName string) (string, *Raw, error) {
	url := c.baseURL + "/customers/" + customerID + "/funding-sources"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return "", nil, err
	}

	fundingSourceReq := &PlaidFundingSourceRequest{
		PlaidToken: plaidToken,
		Name:       fundingSourceName,
	}

	resp, raw, err := post(url, nil, fundingSourceReq, token)
	if err != nil {
		return "", raw, err
	}

	fundingSourceLink := resp.Header.Get(location)

	return fundingSourceLink, raw, nil
}

func (c *customer) GetFundingSourcesForCustomer(customerID string) (*FundingSourcesResponse, *Raw, error) {
	url := c.baseURL + "/customers/" + customerID + "/funding-sources"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var fundingSourceResp FundingSourcesResponse
	if err := json.Unmarshal(resp.Body, &fundingSourceResp); err != nil {
		return nil, raw, err
	}

	return &fundingSourceResp, raw, nil
}

func (c *customer) AddFundingSourceForCustomer(customerID string, fundingSourceReq *FundingSourceRequest) (string, *Raw, error) {
	url := c.baseURL + "/customers/" + customerID + "/funding-sources"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return "", nil, err
	}

	resp, raw, err := post(url, nil, fundingSourceReq, token)
	if err != nil {
		return "", raw, err
	}

	fundingSourceLocation := resp.Header.Get(location)

	return fundingSourceLocation, raw, nil
}

func (c *customer) GetCustomers() (*CustomersResponse, *Raw, error) {
	url := c.baseURL + "/customers"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var customersResponse CustomersResponse
	if err := json.Unmarshal(resp.Body, &customersResponse); err != nil {
		return nil, raw, err
	}

	return &customersResponse, raw, nil
}

func (c *customer) GetCustomer(verifiedCustomerID string) (*Customer, *Raw, error) {
	url := c.baseURL + "/customers/" + verifiedCustomerID

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var customersResponse Customer
	if err := json.Unmarshal(resp.Body, &customersResponse); err != nil {
		return nil, raw, err
	}

	return &customersResponse, raw, nil
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

func (c *customer) AddBeneficialOwner(verifiedCustomerID string, beneficialOwner *BeneficialOwnerRequest) (*BeneficialOwnerRequest, *Raw, error) {
	url := c.baseURL + "/customers/" + verifiedCustomerID + "/beneficial-owners"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := post(url, nil, beneficialOwner, token)
	if err != nil {
		return nil, raw, err
	}

	massPaymentLocation := resp.Header.Get(location)
	massPaymentID, err := ExtractIDFromLocation(massPaymentLocation)
	if err != nil {
		return nil, raw, err
	}

	beneficialOwner.Location = massPaymentLocation
	beneficialOwner.ID = massPaymentID

	return beneficialOwner, raw, nil
}

//A beneficial owner’s information can only be updated if their verification status is incomplete
func (c *customer) UpdateBeneficialOwner(beneficialOwnerID string, beneficialOwner *BeneficialOwnerRequest) (*Raw, error) {
	url := c.baseURL + "/beneficial-owners/" + beneficialOwnerID

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, raw, err := post(url, nil, beneficialOwner, token)
	if err != nil {
		return raw, err
	}

	log.Println(string(resp.Body))
	return raw, nil
}

func (c *customer) GetBeneficialOwnerById(beneficialOwnerID string) (*BeneficialOwner, *Raw, error) {
	url := c.baseURL + "/beneficial-owners/" + beneficialOwnerID

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var data BeneficialOwner
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, raw, err
	}

	return &data, raw, nil
}

func (c *customer) GetAllBeneficialOwners(verifiedCustomerID string) (*BeneficialOwnersResponse, *Raw, error) {
	url := c.baseURL + "/customers/" + verifiedCustomerID + "/beneficial-owners"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var data BeneficialOwnersResponse
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, raw, err
	}

	return &data, raw, nil
}

//A removed beneficial owner cannot be retrieved after being removed.
func (c *customer) DeleteBeneficialOwnerById(beneficialOwnerID string) error {
	url := c.baseURL + "/beneficial-owners/" + beneficialOwnerID

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

func (c *customer) RetrieveBeneficialOwnershipStatus(verifiedCustomerID string) (*BeneficialOwnershipStatusResponse, *Raw, error) {
	url := c.baseURL + "/customers/" + verifiedCustomerID + "/beneficial-ownership"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var data BeneficialOwnershipStatusResponse
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, raw, err
	}

	return &data, raw, nil
}

func (c *customer) CertifyBeneficialOwnership(verifiedCustomerID string, certifyBeneficialOwnership CertifyBeneficialOwnershipReq) (*BeneficialOwnershipStatusResponse, *Raw, error) {
	url := c.baseURL + "/customers/" + verifiedCustomerID + "/beneficial-ownership"

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := post(url, nil, certifyBeneficialOwnership, token)
	if err != nil {
		return nil, raw, err
	}

	var data BeneficialOwnershipStatusResponse
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, raw, err
	}

	return &data, raw, nil
}

//File types supported:
//Personal IDs - .jpg, .jpeg or .png.
//Business Documents - .jpg, .jpeg, .png, or .pdf.
//Files must be no larger than 10MB in size.
func (c *customer) UploadVerificationDocument(beneficialOwnerID, documentType, identity string, fileReq FileRequest) (string, *Raw, error) {
	if fileReq.FileHeader == nil || fileReq.File == nil {
		log.Println("Invalid file object.")
		return "", nil, errors.New("invalid file object")
	}

	var url string
	if identity == "customer" {
		url = c.baseURL + "/customers/" + beneficialOwnerID + "/documents"
	} else {
		url = c.baseURL + "/beneficial-owners/" + beneficialOwnerID + "/documents"
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, err := writer.CreateFormField("documentType")
	if err != nil {
	}
	_, err = io.Copy(fw, strings.NewReader(documentType))
	if err != nil {
		return "", nil, err
	}
	fw, err = writer.CreateFormFile("file", fileReq.FileHeader.Filename)
	if err != nil {
	}
	_, err = io.Copy(fw, fileReq.File)
	if err != nil {
		return "", nil, err
	}
	// Close multipart writer.
	writer.Close()

	token, err := c.authHandler.GetToken()
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	header := &Header{
		ContentType: writer.FormDataContentType(),
	}

	resp, raw, err := upload(url, header, body.Bytes(), token)
	if err != nil {
		log.Println(err)
		return "", raw, err
	}

	documentLocation := resp.Header.Get(location)

	log.Println(string(resp.Body))
	return documentLocation, raw, nil
}

func (c *customer) GetDocumentById(documentId string) (*Document, *Raw, error) {
	url := c.baseURL + "/documents/" + documentId

	token, err := c.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var data Document
	if err := json.Unmarshal(resp.Body, &data); err != nil {
		return nil, raw, err
	}

	return &data, raw, nil
}