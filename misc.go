package btcpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type LanguageCodesRespose struct {
	Code            string `json:"code"`
	CurrentLanguage string `json:"currentLanguage"`
}

func (c *Client) GetLanguageCodes(ctx context.Context) ([]*LanguageCodesRespose, int, error) {
	endpoint := fmt.Sprintf("%s/misc/lang", c.URL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*LanguageCodesRespose
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}

	return dataRes, statusCode, nil
}
