package btcpay

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Synchronized bool `json:"synchronized"`
}

func (c *Client) GetHealth() (*HealthResponse, error) {
	endpoint := fmt.Sprintf("%s/api/v1/health", c.URL)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data HealthResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
