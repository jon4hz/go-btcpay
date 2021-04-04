package btcpay

import (
	"testing"
)

// change url to valid
func TestValidHealthReponseWithBasicClient(t *testing.T) {
	client := NewBasicClient("https://docs.btcpayserver.com", "", "")
	_, err := client.GetHealth()
	if err != nil {
		t.Error("Error while getting the health status: ", err)
	}

}

// change url to valid
func TestValidHealthReponseWithAPIClient(t *testing.T) {
	client := NewClient("https://docs.btcpayserver.com", "")
	_, err := client.GetHealth()
	if err != nil {
		t.Error("Error while getting the health status: ", err)
	}

}

func TestInvalidHealthReponseWithBasicClient(t *testing.T) {
	client := NewBasicClient("https://docs.btcpayserver.com", "", "")
	_, err := client.GetHealth()
	if err == nil {
		// change error message
		t.Error("Error while getting the health status: ", err)
	}

}
func TestInvalidHealthReponseWithAPIClient(t *testing.T) {
	client := NewClient("https://docs.btcpayserver.com", "")
	_, err := client.GetHealth()
	if err == nil {
		// change error message
		t.Error("Error while getting the health status: ", err)
	}

}

func TestInvalidRequestHealthReponseWithBasicClient(t *testing.T) {
	client := NewBasicClient("", "", "")
	_, err := client.GetHealth()
	if err == nil {
		// change error message
		t.Error("Error while getting the health status: ", err)
	}

}
func TestInvalidRequestHealthReponseWithAPIClient(t *testing.T) {
	client := NewClient("", "")
	_, err := client.GetHealth()
	if err == nil {
		// change error message
		t.Error("Error while getting the health status: ", err)
	}

}
