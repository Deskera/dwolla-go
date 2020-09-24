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

type UnverifiedCustomer struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	BusinessName  string `json:"businessName"`
	CorrelationId string `json:"correlationId"`
	CustomerId    string `json:"-"`
}

type CustomerResponse struct {
}

type CustomersResponse struct {
}

type RootResponse struct {
	Links struct {
		Account struct {
			Href string `json:"href"`
		} `json:"account"`
	} `json:"_links"`
}

type AccountDetailsResponse struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Receive struct {
			Href string `json:"href"`
		} `json:"receive"`
		FundingSources struct {
			Href string `json:"href"`
		} `json:"funding-sources"`
		Transfers struct {
			Href string `json:"href"`
		} `json:"transfers"`
		Customers struct {
			Href string `json:"href"`
		} `json:"customers"`
		Send struct {
			Href string `json:"href"`
		} `json:"send"`
	} `json:"_links"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FundingSourcesResponse struct {
}
