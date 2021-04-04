package btcpay

// WIP

type AuthorizationRequest struct {
	Permissions           []string
	ApplicationName       string
	Strict                bool
	SelectiveStores       bool
	Redirect              string
	ApplicationIdentifier string
}
