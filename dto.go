package dwolla

import "time"

type Token struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	ExpiresIn   float64 `json:"expires_in"`
	CreatedAt   time.Time
}

type ReceiveOnlyCustomer struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	Type          string `json:"type"`
	BusinessName  string `json:"businessName"`
	CorrelationId string `json:"correlationId"`
	CustomerId    string `json:"-"`
}

type VerifiedCustomer struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	Type          string `json:"type"`
	BusinessName  string `json:"businessName"`
	CorrelationId string `json:"correlationId"`
	SSN           string `json:"ssn"`
	DateOfBirth   string `json:"dateOfBirth"`
	PostalCode    string `json:"postalCode"`
	State         string `json:"state"`
	City          string `json:"city"`
	Address1      string `json:"address1"`
	CustomerId    string `json:"-"`
}

type UnverifiedCustomer struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	BusinessName  string `json:"businessName"`
	CorrelationId string `json:"correlationId"`
	CustomerId    string `json:"-"`
}

type Customer struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	Created      string `json:"created"`
	BusinessName string `json:"businessName"`
}

type CustomersResponse struct {
	Embedded struct {
		Customers []Customer `json:"customers"`
	} `json:"_embedded"`
}

type RootResponse struct {
	Links struct {
		Account link `json:"account"`
	} `json:"_links"`
}

type AccountDetailsResponse struct {
	Links struct {
		Self           link `json:"self"`
		Receive        link `json:"receive"`
		FundingSources link `json:"funding-sources"`
		Transfers      link `json:"transfers"`
		Customers      link `json:"customers"`
		Send           link `json:"send"`
	} `json:"_links"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Funding struct {
	Id              string   `json:"id"`
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
	RoutingNumber   string `json:"routingNumber"`
	AccountNumber   string `json:"accountNumber"`
	BankAccountType string `json:"bankAccountType"`
	Name            string `json:"name"`
	PlaidToken      string `json:"plaidToken"`
}

type PlaidFundingSourceRequest struct {
	PlaidToken string `json:"plaidToken"`
	Name       string `json:"name"`
}

type FundingSourcesResponse struct {
	Links struct {
		Self link `json:"self"`
	} `json:"_links"`

	Embedded struct {
		FundingSources []Funding `json:"funding-sources"`
	} `json:"_embedded"`
}

type link struct {
	Href         string `json:"href"`
	LinkType     string `json:"type"`
	ResourceType string `json:"resourceType"`
}
