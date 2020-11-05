package dwolla

import (
	"encoding/json"
	"errors"
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
		return nil
	}

	res, err := get(url, token)
	if err != nil {
		return nil
	}

	var root RootResponse
	if err := json.Unmarshal(res.Body, &root); err != nil {
		return nil
	}

	accountId = root.Links.Account.Href
	return nil
}

func (a *account) GetAccountID() (string, error) {
	if accountId == "" {
		return "", errors.New("no accountId")
	}

	accountIDSplit := strings.Split(accountId, "/accounts/")
	if len(accountIDSplit) < 1 {
		return "", errors.New("error extraction id")
	}

	return accountIDSplit[1], nil
}

func (a *account) GetAccountDetails() (*AccountDetailsResponse, error) {
	id, err := a.GetAccountID()
	if err != nil {
		return nil, err
	}

	url := a.baseURL + "/accounts" + "/" + id

	token, err := a.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := get(url, token)
	if err != nil {
		return nil, err
	}

	var accountsResponse AccountDetailsResponse
	if err := json.Unmarshal(res.Body, &accountsResponse); err != nil {
		return nil, err
	}

	return &accountsResponse, nil
}

func (a *account) GetFundingSources() (*FundingSourcesResponse, error) {
	id, err := a.GetAccountID()
	if err != nil {
		return nil, err
	}

	url := a.baseURL + "/accounts" + "/" + id + "/funding-sources"

	token, err := a.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := get(url, token)
	if err != nil {
		return nil, err
	}

	var accountsResponse FundingSourcesResponse
	if err := json.Unmarshal(res.Body, &accountsResponse); err != nil {
		return nil, err
	}

	return &accountsResponse, nil
}
