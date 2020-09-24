package dwolla

type client struct {
	Auth     *auth
	Costumer *customer
	Root     *root
}

type Config struct {
	ClientId     string
	ClientSecret string
	Enviorment   string
}

func NewClient(config *Config) *client {
	baseURL := getBaseURLFromEnviorment(config.Enviorment)
	authConf := &auth{
		clientId:     config.ClientId,
		clientSecret: config.ClientSecret,
		baseURL:      baseURL,
	}
	authHandler := Auth(authConf)
	authHandler.FetchToken()

	customerConf := &customer{
		authHandler: authHandler,
		baseURL:     baseURL,
	}
	customerHandler := Customer(customerConf)

	rootConf := &root{
		authHandler: authHandler,
		baseURL:     baseURL,
	}

	rootHandler := Root(rootConf)
	err := rootHandler.setupRoot()
	if err != nil {
		return nil
	}
	return &client{
		Auth:     authHandler,
		Costumer: customerHandler,
		Root:     rootHandler,
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
