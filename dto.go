package dwolla

type ReceiveOnlyCustomer struct {
	FirstName     string       `json:"firstName"`
	LastName      string       `json:"lastName"`
	Email         string       `json:"email"`
	Type          CustomerType `json:"type"`
	CorrelationID string       `json:"correlationId"`
}

type VerifiedCustomer struct {
	FirstName              string       `json:"firstName"`
	LastName               string       `json:"lastName"`
	Email                  string       `json:"email"`
	Type                   CustomerType `json:"type"`
	BusinessName           string       `json:"businessName"`
	CorrelationID          string       `json:"correlationId"`
	SSN                    string       `json:"ssn"`
	DateOfBirth            string       `json:"dateOfBirth"`
	PostalCode             string       `json:"postalCode"`
	State                  string       `json:"state"`
	City                   string       `json:"city"`
	Address1               string       `json:"address1"`
	Address2               string       `json:"address2"`
	BusinessType           BusinessType `json:"businessType"`
	DoingBusinessAs        string       `json:"doingBusinessAs"`
	BusinessClassification string       `json:"businessClassification"`
	EIN                    string       `json:"ein"`
	Website                string       `json:"website"`
	Phone                  string       `json:"phone"`
}

type UnverifiedCustomer struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	BusinessName  string `json:"businessName"`
	CorrelationID string `json:"correlationId"`
}

type Customer struct {
	ID            string       `json:"id"`
	CorrelationID string       `json:"correlationId"`
	Location      string       `json:"location"`
	FirstName     string       `json:"firstName"`
	LastName      string       `json:"lastName"`
	Email         string       `json:"email"`
	Type          CustomerType `json:"type"`
	Status        string       `json:"status"`
	Created       bool         `json:"created"`
	BusinessName  string       `json:"businessName"`
	RawResponse   string       `json:"rawResponse"`
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

type SelfLink struct {
	Self Link `json:"self"`
}

type DestinationLink struct {
	Destination Link `json:"destination"`
}

type SourceLink struct {
	Source Link `json:"source,omitempty"`
}

type MassPayment struct {
	Links         SourceLink        `json:"_links,omitempty"`
	ID            string            `json:"id,omitempty"`
	Items         []MassPaymentItem `json:"items,omitempty"`
	Status        MassPaymentStatus `json:"status,omitempty"`
	CorrelationID string            `json:"correlationId,omitempty"`
	Metadata      interface{}       `json:"metadata,omitempty"`
	Location      string            `json:"-"`
}

// MassPaymentItem is a dwolla mass payment item
type MassPaymentItem struct {
	Links         DestinationLink       `json:"_links,omitempty"`
	Amount        Amount                `json:"amount,omitempty"`
	Status        MassPaymentItemStatus `json:"status,omitempty"`
	Metadata      interface{}           `json:"metadata,omitempty"`
	CorrelationID string                `json:"correlationId,omitempty"`
	Embedded      Embedded              `json:"_embedded,omitempty,omitempty"`
}

type Amount struct {
	Value    string   `json:"value"`
	Currency Currency `json:"currency"`
}

type Header struct {
	IdempotencyKey string
}

type UpdateMassPayment struct {
	Status MassPaymentStatus `json:"status"`
}

type IndustryClassification struct {
	Name     string `json:"name"`
	DwollaId string `json:"id"`
}

type BusinessClassification struct {
	Name     string `json:"name"`
	DwollaID string `json:"id"`
	Embedded struct {
		IndustryClassifications []IndustryClassification `json:"industry-classifications"`
	} `json:"_embedded"`
}

type BusinessClassificationsResponse struct {
	Embedded struct {
		BusinessClassifications []BusinessClassification `json:"business-classifications"`
	} `json:"_embedded"`
}

// FundingSourceBalance is a funding source balance
type FundingSourceBalance struct {
	Balance     Amount `json:"balance"`
	LastUpdated string `json:"lastUpdated"`
}

type MassPaymentResponse struct {
	Links    Links             `json:"_links"`
	ID       string            `json:"id"`
	Status   MassPaymentStatus `json:"status"`
	Created  string            `json:"created"`
	Metadata interface{}       `json:"metadata"`
}

type WebhookSubscription struct {
	Id      string `json:"id"`
	Url     string `json:"url"`
	Created string `json:"created"`
}

type WebhookSubscriptionsResponse struct {
	Embedded struct {
		Subscriptions []WebhookSubscription `json:"webhook-subscriptions"`
	} `json:"_embedded"`
}

type WebhookSubscriptionRequest struct {
	URL    string `json:"url"`
	Secret string `json:"secret"`
}
