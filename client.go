package dwolla

//Client is the dwolla client
type Client struct {
	Auth                 *auth
	Customer             *customer
	Account              *account
	Business             *business
	Payment              *massPayment
	WebhookSubscriptions *webhook
	Transfer             *transfer
	FundingSource        *fundingSource
}

//Config ...
type Config struct {
	ClientKey    string
	ClientSecret string
	Enviorment   string
}

//NewClient setups a new dwolla client
func NewClient(config *Config) (*Client, error) {
	baseURL := getBaseURLFromEnviorment(config.Enviorment)
	authHandler := &auth{
		clientID:     config.ClientKey,
		clientSecret: config.ClientSecret,
		baseURL:      baseURL,
	}

	authHandler.FetchToken()

	customerHandler := &customer{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	rootHandler := &account{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	_, err := rootHandler.setupRoot()
	if err != nil {
		return nil, err
	}

	businessClassificationsHandler := &business{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	paymentHandler := &massPayment{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	webhookHandler := &webhook{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	transferHandler := &transfer{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	fundingSourceHandler := &fundingSource{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	return &Client{
		Auth:                 authHandler,
		Customer:             customerHandler,
		Account:              rootHandler,
		Business:             businessClassificationsHandler,
		Payment:              paymentHandler,
		WebhookSubscriptions: webhookHandler,
		Transfer:             transferHandler,
		FundingSource:        fundingSourceHandler,
	}, nil
}

func getBaseURLFromEnviorment(enviorment string) string {
	var baseURL string
	switch enviorment {
	case "sandbox":
		baseURL = "https://api-sandbox.dwolla.com"
	case "production":
		baseURL = "https://api.dwolla.com"
	}

	return baseURL
}
