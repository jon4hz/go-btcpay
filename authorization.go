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
	Permissions           []Permission `json:"permissions,omitempty"`
	ApplicationName       string       `json:"applicationName,omitempty"`
	Strict                bool         `json:"strict,omitempty"`
	SelectiveStores       bool         `json:"selectiveStores,omitempty"`
	Redirect              string       `json:"redirect,omitempty"`
	ApplicationIdentifier string       `json:"applicationIdentifier,omitempty"`
}

func (c *Client) Authorize(authRequest *AuthorizationRequest, ctx context.Context) error {
	endpoint := fmt.Sprintf("%s/api/v1/authorize", c.URL)
	dataReq, err := json.Marshal(authRequest)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}
