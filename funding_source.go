package dwolla

import "encoding/json"

type fundingSource struct {
	authHandler *auth
	baseURL     string
}

// FundingSourceStatus is a funding source's status
type FundingSourceStatus string

// FundingSourceType is the funding source type
type FundingSourceType string

// FundingSourceBankAccountType is a dwolla bank account type enum
type FundingSourceBankAccountType string

const (
	// FundingSourceBankAccountTypeChecking is a checking bank account
	FundingSourceBankAccountTypeChecking FundingSourceBankAccountType = "checking"
	// FundingSourceBankAccountTypeSavings is a savings bank account
	FundingSourceBankAccountTypeSavings FundingSourceBankAccountType = "savings"
)

const (
	// FundingSourceStatusUnverified is when the funding source is unverified
	FundingSourceStatusUnverified FundingSourceStatus = "unverified"
	// FundingSourceStatusVerified is when the funding source is verified
	FundingSourceStatusVerified FundingSourceStatus = "verified"
)

const (
	// FundingSourceTypeBank is when the funding source is a bank account
	FundingSourceTypeBank FundingSourceType = "bank"
	// FundingSourceTypeBalance is when the funding source is a dwolla balance
	FundingSourceTypeBalance FundingSourceType = "balance"
)

type Funding struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Status          string   `json:"status"`
	Type            string   `json:"type"`
	BankAccountType string   `json:"bankAccountType"`
	Created         string   `json:"created"`
	Removed         bool     `json:"removed"`
	Channels        []string `json:"channels"`
	BankName        string   `json:"bankName"`
}

type FundingSourceRequest struct {
	RoutingNumber   string                       `json:"routingNumber"`
	AccountNumber   string                       `json:"accountNumber"`
	BankAccountType FundingSourceBankAccountType `json:"bankAccountType"`
	Name            string                       `json:"name"`
	PlaidToken      string                       `json:"plaidToken"`
}

type PlaidFundingSourceRequest struct {
	PlaidToken string `json:"plaidToken"`
	Name       string `json:"name"`
}

// FundingSource is a dwolla funding source
type FundingSource struct {
	Resource
	ID              string                       `json:"id"`
	Status          FundingSourceStatus          `json:"status"`
	Type            FundingSourceType            `json:"type"`
	BankAccountType FundingSourceBankAccountType `json:"bankAccountType"`
	Name            string                       `json:"name"`
	Created         string                       `json:"created"`
	Balance         Amount                       `json:"balance"`
	Removed         bool                         `json:"removed"`
	Channels        []string                     `json:"channels"`
	BankName        string                       `json:"bankName"`
	Fingerprint     string                       `json:"fingerprint"`
}

type FundingSourcesResponse struct {
	Links    SelfLink `json:"_links"`
	Embedded struct {
		FundingSources []Funding `json:"funding-sources"`
	} `json:"_embedded"`
}

func (f *fundingSource) GetFundingSourcesBalance(fundingSourceID string) (*FundingSourceBalance, *Raw, error) {
	url := f.baseURL + "/funding-sources/" + fundingSourceID + "/balance"

	token, err := f.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	resp, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var fundingSourceResp FundingSourceBalance
	if err := json.Unmarshal(resp.Body, &fundingSourceResp); err != nil {
		return nil, raw, err
	}

	return &fundingSourceResp, raw, nil
}