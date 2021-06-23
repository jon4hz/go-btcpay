package btcpay

import (
	"context"
	"fmt"
)

func (c *Client) RevokeAPIKey(ctx context.Context, apiKey *APIKey) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys/%s", c.URL, *apiKey)
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", nil, nil)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (c *Client) GetCurrentAPIKey(ctx context.Context) (*APIKeyResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys/current", c.URL)
	var dataRes APIKeyResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (c *Client) RevokeCurrentAPIKey(ctx context.Context) (*APIKeyResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys/current", c.URL)
	var dataRes APIKeyResponse
	statusCode, err := c.doRequest(ctx, endpoint, "DELETE", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}

func (c *Client) CreateAPIKey(ctx context.Context, apiKeyRequest *APIKeyRequest) (*APIKeyResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys", c.URL)
	var dataRes APIKeyResponse
	statusCode, err := c.doRequest(ctx, endpoint, "POST", &dataRes, apiKeyRequest)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
