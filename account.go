package dwolla

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

type root struct {
	authHandler *auth
	baseURL     string
}

func Root(rootConfig *root) *root {
	return rootConfig
}

var accountId string

func (r *root) setupRoot() error {
	url := r.baseURL + "/"
	accept := "application/vnd.dwolla.v1.hal+json"

	token, err := r.authHandler.GetToken()
	if err != nil {
		log.Println(err)
		return nil
	}

	res, err := makeGetRequest(url, accept, nil, token)
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

func (r *root) GetAccountId() (string, error) {
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

func (r *root) GetAccountDetails() {
	id, err := r.GetAccountId()
	if err != nil {
		return
	}

	url := r.baseURL + "/accounts" + "/" + id
	accept := "application/vnd.dwolla.v1.hal+json"

	token, err := r.authHandler.GetToken()
	if err != nil {
		log.Println(err)
		return
	}

	res, err := makeGetRequest(url, accept, nil, token)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	var accountsResponse AccountDetailsResponse
	if err := json.NewDecoder(res.Body).Decode(&accountsResponse); err != nil {
		return
	}

	log.Printf("%+v\n", accountsResponse)
}

func (r *root) GetFundingSources() {
	id, err := r.GetAccountId()
	if err != nil {
		return
	}

	url := r.baseURL + "/accounts" + "/" + id + "/funding-sources"
	accept := "application/vnd.dwolla.v1.hal+json"

	token, err := r.authHandler.GetToken()
	if err != nil {
		log.Println(err)
		return
	}

	res, err := makeGetRequest(url, accept, nil, token)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	var accountsResponse AccountDetailsResponse
	if err := json.NewDecoder(res.Body).Decode(&accountsResponse); err != nil {
		return
	}

	log.Printf("%+v\n", accountsResponse)
}
