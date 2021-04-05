package btcpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Synchronized bool `json:"synchronized"`
}

func (c *Client) GetHealth(ctx context.Context) (*HealthResponse, error) {
	endpoint := fmt.Sprintf("%s/api/v1/health", c.URL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var dataRes HealthResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, err
	}
	return &dataRes, nil
}
