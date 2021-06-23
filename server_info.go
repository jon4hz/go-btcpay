package btcpay

import (
	"context"
	"fmt"
)

func (c *Client) GetServerInfo(ctx context.Context) (*ServerInfoResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/server/info", c.URL)
	var dataRes ServerInfoResponse
	statusCode, err := c.doRequest(ctx, endpoint, "GET", &dataRes, nil)
	if err != nil {
		return nil, statusCode, err
	}
	return &dataRes, statusCode, nil
}
