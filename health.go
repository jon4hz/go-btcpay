package btcpay

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BTCPayHealthResponse struct {
	Synchronized bool `json:"synchronized"`
}

func (c *Client) GetHealth() (*BTCPayHealthResponse, error) {
	endpoint := fmt.Sprintf("%s/api/v1/health", c.URL)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data BTCPayHealthResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *BasicClient) GetHealth() (*BTCPayHealthResponse, error) {
	endpoint := fmt.Sprintf("%s/api/v1/health", c.URL)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data BTCPayHealthResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
