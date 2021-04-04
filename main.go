package btcpay

func main() {
	var client BTCPayClient
	client = NewClient("https://docs.btcpayserver.org/", "")
	client.GetHealth()
}
