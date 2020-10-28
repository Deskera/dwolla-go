package dwolla

type Client struct {
	Auth     *auth
	Customer *customer
	Account  *account
	Business *business
	Payment  *payment
}

type Config struct {
	ClientId     string
	ClientSecret string
	Enviorment   string
}

func NewClient(config *Config) (*Client, error) {
	baseURL := getBaseURLFromEnviorment(config.Enviorment)
	authConf := &auth{
		clientId:     config.ClientId,
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

	paymentHandler := &payment{
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
	}

	return baseURL
}
