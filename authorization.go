package btcpay

type BTCPayAuthorizationRequest struct {
	permissions           []string
	applicationName       string
	strict                bool
	selectiveStores       bool
	redirect              string
	applicationIdentifier string
}
