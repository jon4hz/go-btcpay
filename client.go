package btcpay

// WIP (webhooks?)

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewClient(url string, apiKey APIKey) *Client {
	client := &Client{
		URL:    url,
		APIKey: apiKey,
		Http:   &http.Client{},
	}
	client.Store = &Store{Client: client}
	client.Invoice = &Invoice{Client: client, Store: client.Store}
	client.PaymentRequest = &PaymentRequest{Client: client, Store: client.Store}
	client.PullPayment = &PullPayment{Client: client, Store: client.Store}
	client.Payout = &Payout{Client: client, Store: client.Store}
	client.Notification = &Notification{Client: client}
	return client
}

func NewBasicClient(url, username, password string) *Client {
	client := &Client{
		URL:      url,
		Username: username,
		Password: password,
		Http:     &http.Client{},
	}
	client.Store = &Store{Client: client}
	client.Invoice = &Invoice{Client: client, Store: client.Store}
	client.PaymentRequest = &PaymentRequest{Client: client, Store: client.Store}
	client.PullPayment = &PullPayment{Client: client, Store: client.Store}
	client.Payout = &Payout{Client: client, Store: client.Store}
	client.Notification = &Notification{Client: client}
	return client
}

func setHeaders(c *Client, r *http.Request) {
	if len(c.Username) > 0 && len(c.Password) > 0 {
		r.SetBasicAuth(c.Username, c.Password)
	} else if len(c.APIKey) > 0 {
		r.Header.Set("Authorization", fmt.Sprintf("token %s", c.APIKey))
	}
	r.Header.Set("Content-Type", "application/json")
}

func setQueryParam(endpoint *string, params []map[string]interface{}) {
	for _, param := range params {
		for i := range param {
			*endpoint = fmt.Sprintf("%s?%s=%v", *endpoint, i, param[i])
		}
	}
}

func (c *Client) doRequest(ctx context.Context, endpoint, method string, expRes interface{}, reqData interface{}, opts ...map[string]interface{}) (int, error) {

	var dataReq []byte
	var err error

	if reqData != nil {
		dataReq, err = json.Marshal(reqData)
		if err != nil {
			return 0, err
		}
	}

	if len(opts) > 0 && len(opts[0]) > 0 {
		setQueryParam(&endpoint, opts)
	}
	fmt.Println(endpoint)
	req, err := http.NewRequestWithContext(ctx, method, endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return 0, err
	}

	setHeaders(c, req)

	resp, err := c.Http.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	switch resp.StatusCode {
	case 200:
		if expRes != nil {
			err = json.Unmarshal(body, expRes)
			if err != nil {
				return 0, err
			}
		}
		return resp.StatusCode, nil

	default:
		return resp.StatusCode, fmt.Errorf("%s", body)
	}
}

func (c *Client) doSimpleRequest(req *http.Request) ([]byte, int, error) {

	setHeaders(c, req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	switch resp.StatusCode {
	case 200:
		return body, resp.StatusCode, nil
	default:
		return nil, resp.StatusCode, fmt.Errorf("%s", body)
	}
}
