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
	Links AccountLink `json:"_links"`
}

type AccountLink struct {
	Account Link `json:"account"`
}

type AccountDetailsResponse struct {
	Links struct {
		Self           Link `json:"self"`
		Receive        Link `json:"receive"`
		FundingSources Link `json:"funding-sources"`
		Transfers      Link `json:"transfers"`
		Customers      Link `json:"customers"`
		Send           Link `json:"send"`
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
	Links    SelfLink `json:"_links"`
	Embedded struct {
		FundingSources []Funding `json:"funding-sources"`
	} `json:"_embedded"`
}

type SelfLink struct {
	Self Link `json:"self"`
}

type Link struct {
	Href         string `json:"href"`
	LinkType     string `json:"type"`
	ResourceType string `json:"resourceType"`
}

type MassPayment struct {
	Links         SourceLink    `json:"_links"`
	Items         []PaymentItem `json:"items"`
	Status        PaymentStatus `json:"status"`
	CorrelationId string        `json:"correlationId"`
	Location      string        `json:"-"`
}

type SourceLink struct {
	Source Link `json:"source"`
}

type DestinationLink struct {
	Destination Link `json:"destination"`
}
type PaymentItem struct {
	Links         DestinationLink `json:"_links"`
	Amount        Amount          `json:"amount"`
	Metadata      interface{}     `json:"metadata"`
	CorrelationId string          `json:"correlationId"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Header struct {
	IdempotencyKey string
}

type UpdateMassPayment struct {
	Status PaymentStatus `json:"status"`
}

type IndustryClassification struct {
	Name     string `json:"name"`
	DwollaId string `json:"id"`
}

type BusinessClassification struct {
	Name     string `json:"name"`
	DwollaId string `json:"id"`
	Embedded struct {
		IndustryClassifications []IndustryClassification `json:"industry-classifications"`
	} `json:"_embedded"`
}

type BusinessClassificationsResponse struct {
	Embedded struct {
		BusinessClassifications []BusinessClassification `json:"business-classifications"`
	} `json:"_embedded"`
}
