package dwolla

// Raw represents the actual request and response sent/received by dwolla
type Raw struct {
	Endpoint string
	Request  string
	Response string
	Status   int
}

// Controller is a controller of a business
type Controller struct {
	FirstName   string   `json:"firstName,omitempty"`
	LastName    string   `json:"lastName,omitempty"`
	Title       string   `json:"title,omitempty"`
	DateOfBirth string   `json:"dateOfBirth,omitempty"`
	SSN         string   `json:"ssn,omitempty"`
	Address     Address  `json:"address,omitempty"`
	Passport    Passport `json:"passport,omitempty"`
}

// Address represents a street address
type Address struct {
	Address1            string `json:"address1"`
	Address2            string `json:"address2,omitempty"`
	Address3            string `json:"address3,omitempty"`
	City                string `json:"city"`
	StateProvinceRegion string `json:"stateProvinceRegion"`
	PostalCode          string `json:"postalCode,omitempty"`
	Country             string `json:"country"`
}

// Passport represents a passport
type Passport struct {
	Number  string `json:"number"`
	Country string `json:"country"`
}

// Customer is a dwolla customer
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
}

// CustomerRequest is a customer create/update request
type CustomerRequest struct {
	FirstName              string             `json:"firstName,omitempty"`
	LastName               string             `json:"lastName,omitempty"`
	Email                  string             `json:"email,omitempty"`
	CorrelationID          string             `json:"correlationId"`
	IPAddress              string             `json:"ipAddress,omitempty"`
	Type                   CustomerType       `json:"type,omitempty"`
	Status                 CustomerStatus     `json:"status,omitempty"`
	DateOfBirth            string             `json:"dateOfBirth,omitempty"`
	SSN                    string             `json:"ssn,omitempty"`
	Phone                  string             `json:"phone,omitempty"`
	Address1               string             `json:"address1,omitempty"`
	Address2               string             `json:"address2,omitempty"`
	City                   string             `json:"city,omitempty"`
	State                  string             `json:"state,omitempty"`
	PostalCode             string             `json:"postalCode,omitempty"`
	BusinessClassification string             `json:"businessClassification,omitempty"`
	BusinessType           BusinessType       `json:"businessType,omitempty"`
	BusinessName           string             `json:"businessName,omitempty"`
	DoingBusinessAs        string             `json:"doingBusinessAs,omitempty"`
	EIN                    string             `json:"ein,omitempty"`
	Website                string             `json:"website,omitempty"`
	Controller             *ControllerRequest `json:"controller,omitempty"`
}

// ControllerRequest is a controller of a business create/update request
type ControllerRequest struct {
	FirstName   string    `json:"firstName,omitempty"`
	LastName    string    `json:"lastName,omitempty"`
	Title       string    `json:"title,omitempty"`
	DateOfBirth string    `json:"dateOfBirth,omitempty"`
	SSN         string    `json:"ssn,omitempty"`
	Address     Address   `json:"address,omitempty"`
	Passport    *Passport `json:"passport,omitempty"`
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

// MassPayment is a dwolla mass payment
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
	Embedded      Embedded              `json:"_embedded,omitempty"`
}

// MassPaymentItems is a collection of mass payment items
type MassPaymentItems struct {
	Embedded map[string][]MassPaymentItem `json:"_embedded"`
	Total    int                          `json:"total"`
}

// Amount stores the amount object required by dwolla
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
	DwollaID string `json:"id"`
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
	ID      string `json:"id"`
	URL     string `json:"url"`
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

type UpdateWebhookSubscriptionRequest struct {
	Pause    bool `json:"pause"`
}
