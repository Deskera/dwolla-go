package dwolla_test

import (
	"testing"

	"github.com/Deskera/dwolla-go"
)

func TestNewClient(t *testing.T) {
	conf := &dwolla.Config{
		ClientId:     "<DWOLLA_KEY>",
		ClientSecret: "<DWOLLA_SECRET>",
		Enviorment:   "sandbox",
	}
	client := dwolla.NewClient(conf)
	id, err := client.Root.GetAccountId()
	if err != nil {
		t.Log(err)
	}
	t.Log("id: ", id)
}
