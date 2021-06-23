package btcpay

import (
	"context"
	"fmt"
)

func (c *Client) GetHealth(ctx context.Context) (*HealthResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/health", c.URL)
	var dataRes HealthResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
