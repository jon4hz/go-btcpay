package btcpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) RevokeAPIKey(apiKey *APIKey, ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys/%s", c.URL, *apiKey)
	req, err := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return 0, err
	}
	_, statusCode, err := c.doRequest(req)
	if err != nil {
		return statusCode, err
	}

	return statusCode, nil
}

type APIKeyResponse struct {
	APIKey      APIKey             `json:"apiKey"`
	Label       string             `json:"label"`
	Permissions []BTCPayPermission `json:"permissions"`
}

func (c *Client) GetCurrentAPIKey(ctx context.Context) (*APIKeyResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys/current", c.URL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes APIKeyResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (c *Client) RevokeCurrentAPIKey(ctx context.Context) (*APIKeyResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys/current", c.URL)
	req, err := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var data APIKeyResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, 0, err
	}
	return &data, statusCode, nil
}

type APIKeyRequest struct {
	Label       string             `json:"label,omitempty"`
	Permissions []BTCPayPermission `json:"permissions,omitempty"`
}

func (c *Client) CreateAPIKey(apiKeyRequest *APIKeyRequest, ctx context.Context) (*APIKeyResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys", c.URL)
	dataReq, err := json.Marshal(apiKeyRequest)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes APIKeyResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}
