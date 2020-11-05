package dwolla_test

import (
	"testing"

	"github.com/Deskera/dwolla-go"
)

func TestNewClient(t *testing.T) {
	conf := &dwolla.Config{
		ClientKey:    "",
		ClientSecret: "",
		Enviorment:   "sandbox",
	}
	client, err := dwolla.NewClient(conf)
	if err != nil {
		t.Log(err)
	}
	id, err := client.Account.GetAccountID()
	if err != nil {
		t.Log(err)
	}
	t.Log("id: ", id)

	accountDetails, err := client.Account.GetAccountDetails()
	if err != nil {
		t.Log(err)
	}
	t.Logf("account: %+v\n", accountDetails)

	fundingSources, err := client.Account.GetFundingSources()
	if err != nil {
		t.Log(err)
	}

	t.Logf("funding sources: %+v\n", fundingSources)

	customers, err := client.Customer.GetCustomers()
	if err != nil {
		t.Log(err)
	}

	t.Logf("Customers: %+v\n", customers)
}
