package btcpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ServerInfoResponse struct {
	Version                 string             `json:"version"`
	Onion                   string             `json:"onion"`
	SupportedPaymentMethods []string           `json:"supportedPaymentMethods"`
	FullySynched            bool               `json:"fullySynched"`
	SyncStatus              []ServerSyncStatus `json:"syncStatus"`
}

type ServerSyncStatus struct {
	CryptoCode      string                `json:"cryptoCode"`
	NodeInformation ServerNodeInformation `json:"nodeInformation,omitempty"`
	ChainHeight     int64                 `json:"chainHeight"`
	SyncHeight      int64                 `json:"syncHeight,omitempty"`
}

type ServerNodeInformation struct {
	Headers              int64   `json:"headers"`
	Blocks               int64   `json:"blocks"`
	VerificationProgress float64 `json:"verificationProgress"`
}

func (c *Client) GetServerInfo(ctx context.Context) (*ServerInfoResponse, error) {
	endpoint := fmt.Sprintf("%s/api/v1/server/info", c.URL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var dataRes ServerInfoResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, err
	}
	return &dataRes, nil
}
