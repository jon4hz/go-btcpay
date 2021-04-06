package btcpay

import (
	"context"
	"testing"

	"github.com/jon4hz/go-btcpay/examples/config"
)

// change url to valid
func TestValidHealthReponseWithBasicClient(t *testing.T) {
	config, err := config.ReadConf("examples/config/config.yml")
	if err != nil {
		t.Error("Error while reading the config")
	}
	client := NewBasicClient(config.BTCPay.URL, config.BTCPay.Username, config.BTCPay.Password)
	_, _, err = client.GetHealth(context.Background())
	if err != nil {
		t.Error("Error while getting the health status: ", err)
	}

}

func TestInvalidHealthReponse(t *testing.T) {
	client := NewClient("https://docs.btcpayserver.com", "")
	req, _, err := client.GetHealth(context.Background())
	if err == nil {
		t.Error("Should have thrown an error because URL is invalid: expected nil, got ", req)
	}

}
