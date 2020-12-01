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

var accountID string

func (a *account) setupRoot() (*Raw, error) {
	url := a.baseURL + "/"

	token, err := a.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	res, raw, err := get(url, token)
	if err != nil {
		return raw, err
	}

	var root RootResponse
	if err := json.Unmarshal(res.Body, &root); err != nil {
		return raw, err
	}

	accountID = root.Links.Account.Href
	return raw, nil
}

func (a *account) GetAccountID() (string, error) {
	if accountID == "" {
		return "", errors.New("no accountId")
	}

	accountIDSplit := strings.Split(accountID, "/accounts/")
	if len(accountIDSplit) < 1 {
		return "", errors.New("error extraction id")
	}

	return accountIDSplit[1], nil
}

func (a *account) GetAccountDetails() (*AccountDetailsResponse, *Raw, error) {
	id, err := a.GetAccountID()
	if err != nil {
		return nil, nil, err
	}

	url := a.baseURL + "/accounts" + "/" + id

	token, err := a.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	res, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var accountsResponse AccountDetailsResponse
	if err := json.Unmarshal(res.Body, &accountsResponse); err != nil {
		return nil, raw, err
	}

	return &accountsResponse, raw, nil
}

func (a *account) GetFundingSources() (*FundingSourcesResponse, *Raw, error) {
	id, err := a.GetAccountID()
	if err != nil {
		return nil, nil, err
	}

	url := a.baseURL + "/accounts" + "/" + id + "/funding-sources"

	token, err := a.authHandler.GetToken()
	if err != nil {
		return nil, nil, err
	}

	res, raw, err := get(url, token)
	if err != nil {
		return nil, raw, err
	}

	var accountsResponse FundingSourcesResponse
	if err := json.Unmarshal(res.Body, &accountsResponse); err != nil {
		return nil, raw, err
	}

	return &accountsResponse, raw, nil
}
