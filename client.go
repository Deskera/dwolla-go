package dwolla

//Client is the dwolla client
type Client struct {
	Auth     *auth
	Customer *customer
	Account  *account
	Business *business
	Payment  *massPayment
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
	authConf := &auth{
		clientId:     config.ClientKey,
		clientSecret: config.ClientSecret,
		baseURL:      baseURL,
	}
	authHandler := AuthHandler(authConf)
	authHandler.FetchToken()

	customerConf := &customer{
		authHandler: authHandler,
		baseURL:     baseURL,
	}
	customerHandler := CustomerHandler(customerConf)

	rootConf := &account{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	rootHandler := AccountHandler(rootConf)
	err := rootHandler.setupRoot()
	if err != nil {
		return nil, err
	}

	businessClassifications := &business{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	paymentHandler := &massPayment{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	return &Client{
		Auth:     authHandler,
		Customer: customerHandler,
		Account:  rootHandler,
		Business: businessClassifications,
		Payment:  paymentHandler,
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
