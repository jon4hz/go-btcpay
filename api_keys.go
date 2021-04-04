package btcpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RevokeAPIKeyRequest struct {
	APIKey string `json:"apiKey"`
}

func (c *Client) RevokeAPIKey(apiKey *RevokeAPIKeyRequest, ctx context.Context) error {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys/%s", c.URL, *apiKey)
	dataReq, err := json.Marshal(apiKey)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("DELETE", endpoint, bytes.NewBuffer(dataReq))
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

type APIKeyResponse struct {
	APIKey      *RevokeAPIKeyRequest
	Label       string   `json:"label"`
	Permissions []string `json:"permissions"` // maybe change to struct?
}

func (c *Client) GetCurrentAPIKey(ctx context.Context) (*APIKeyResponse, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys/current", c.URL)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var dataRes APIKeyResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, err
	}
	return &dataRes, nil
}

func (c *Client) RevokeCurrentAPIKey(ctx context.Context) (*APIKeyResponse, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys/current", c.URL)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data APIKeyResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type CreateAPIKeyRequest struct {
	Label       string   `json:"label,omitempty"`
	Permissions []string `json:"permissions,omitempty"` // maybe change to struct?
}

func (c *Client) CreateAPIKey(apiKeyRequest *CreateAPIKeyRequest, ctx context.Context) (*APIKeyResponse, error) {
	endpoint := fmt.Sprintf("%s/api/v1/api-keys", c.URL)
	dataReq, err := json.Marshal(apiKeyRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var dataRes APIKeyResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, err
	}
	return &dataRes, nil
}
