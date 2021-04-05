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

func (c *Client) GetHealth(ctx context.Context) (*HealthResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/health", c.URL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes HealthResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}
