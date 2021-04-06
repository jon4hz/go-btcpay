package btcpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// WIP, needs testing

type AuthorizationRequest struct {
	Permissions           []BTCPayPermission `json:"permissions,omitempty"`
	ApplicationName       string             `json:"applicationName,omitempty"`
	Strict                bool               `json:"strict,omitempty"`
	SelectiveStores       bool               `json:"selectiveStores,omitempty"`
	Redirect              string             `json:"redirect,omitempty"`
	ApplicationIdentifier string             `json:"applicationIdentifier,omitempty"`
}

func (c *Client) Authorize(ctx context.Context, authRequest *AuthorizationRequest) (int, error) {
	endpoint := fmt.Sprintf("%s/api-keys/authorize", c.URL)
	dataReq, err := json.Marshal(authRequest)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return 0, err
	}
	_, statusCode, err := c.doRequest(req)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
