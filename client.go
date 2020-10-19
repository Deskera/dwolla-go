package dwolla

type Client struct {
	Auth     *auth
	Costumer *customer
	Account  *account
}

type Config struct {
	ClientId     string
	ClientSecret string
	Enviorment   string
}

func NewClient(config *Config) *Client {
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
		return nil
	}
	return &Client{
		Auth:     authHandler,
		Costumer: customerHandler,
		Account:  rootHandler,
	}
}

func getBaseURLFromEnviorment(enviorment string) string {
	var baseURL string
	switch enviorment {
	case "sandbox":
		baseURL = "https://api-sandbox.dwolla.com"
	}

	return baseURL
}
