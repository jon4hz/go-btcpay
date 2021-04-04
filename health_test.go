package btcpay

import (
	"testing"
)

// change url to valid
func TestValidHealthReponse(t *testing.T) {
	client := NewClient("https://docs.btcpayserver.com", "")
	_, err := client.GetHealth()
	if err != nil {
		t.Error("Error while getting the health status: ", err)
	}

}

func TestInvalidHealthReponse(t *testing.T) {
	client := NewClient("https://docs.btcpayserver.com", "")
	_, err := client.GetHealth()
	if err == nil {
		// change error message
		t.Error("Error while getting the health status: ", err)
	}

}

func TestInvalidRequestHealthReponse(t *testing.T) {
	client := NewClient("", "")
	_, err := client.GetHealth()
	if err == nil {
		// change error message
		t.Error("Error while getting the health status: ", err)
	}

}
