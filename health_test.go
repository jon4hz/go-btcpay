package btcpay

import (
	"testing"
)

func TestHealthReponse(t *testing.T) {
	client := NewBasicClient("https://docs.btcpayserver.com", "", "")
	_, err := client.GetHealth()
	if err != nil {
		t.Error("Error while getting the health status: ", err)
	}

}
