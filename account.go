package dwolla

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

type account struct {
	authHandler *auth
	baseURL     string
}

func AccountHandler(accountConfig *account) *account {
	return accountConfig
}

var accountId string

func (a *account) setupRoot() error {
	url := a.baseURL + "/"

	token, err := a.authHandler.GetToken()
	if err != nil {
		log.Println(err)
		return nil
	}

	res, err := makeGetRequest(url, nil, token)
	if err != nil {
		log.Println(err)
		return nil
	}

	defer res.Body.Close()

	var root RootResponse
	if err := json.NewDecoder(res.Body).Decode(&root); err != nil {
		return nil
	}

	accountId = root.Links.Account.Href
	return nil
}

func (a *account) GetAccountId() (string, error) {
	if accountId == "" {
		return "", errors.New("no accountId")
	}

	accountIdSplit := strings.Split(accountId, "/accounts/")
	if len(accountIdSplit) < 1 {
		return "", errors.New("error extraction id")
	}

	accountIdFetched := accountIdSplit[1]
	return accountIdFetched, nil
}

func (a *account) GetAccountDetails() (*AccountDetailsResponse, error) {
	id, err := a.GetAccountId()
	if err != nil {
		return nil, err
	}

	url := a.baseURL + "/accounts" + "/" + id

	token, err := a.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := makeGetRequest(url, nil, token)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var accountsResponse AccountDetailsResponse
	if err := json.NewDecoder(res.Body).Decode(&accountsResponse); err != nil {
		return nil, err
	}

	log.Printf("%+v\n", accountsResponse)

	return &accountsResponse, nil
}

func (a *account) GetFundingSources() (*FundingSourcesResponse, error) {
	id, err := a.GetAccountId()
	if err != nil {
		return nil, err
	}

	url := a.baseURL + "/accounts" + "/" + id + "/funding-sources"

	token, err := a.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := makeGetRequest(url, nil, token)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var accountsResponse FundingSourcesResponse
	if err := json.NewDecoder(res.Body).Decode(&accountsResponse); err != nil {
		return nil, err
	}

	return &accountsResponse, nil
}
