package btcpay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type StoreID string

type Store struct {
	ID     StoreID
	Client *Client
}

// Enums NetworkFeeMode
type BTCPayNetworkFeeMode string

type NetworkFeeMode struct {
	MultiplePaymentsOnly BTCPayNetworkFeeMode
	Always               BTCPayNetworkFeeMode
	Never                BTCPayNetworkFeeMode
}

func GetNetworkFeeMode() *NetworkFeeMode {
	return &NetworkFeeMode{
		MultiplePaymentsOnly: "MultiplePaymentsOnly",
		Always:               "Always",
		Never:                "Never",
	}
}

type StoreResponse struct {
	Name                         string               `json:"name"`
	Website                      string               `json:"website"`
	InvoiceExpiration            int64                `json:"invoiceExpiration"`
	MonitoringExpiration         int64                `json:"monitoringExpiration"`
	SpeedPolicy                  BTCPaySpeedPolicy    `json:"speedPolicy"`
	LightningDescriptionTemplate string               `json:"lightningDescriptionTemplate,omitempty"`
	PaymentTolerance             float64              `json:"paymentTolerance"`
	AnyoneCanCreateInvoice       bool                 `json:"anyoneCanCreateInvoice"`
	RequiresRefundEmail          bool                 `json:"requiresRefundEmail"`
	LightningAmountInSatoshi     bool                 `json:"lightningAmountInSatoshi"`
	LightningPrivateRouteHints   bool                 `json:"lightningPrivateRouteHints"`
	OnChainWithLnInvoiceFallback bool                 `json:"onChainWithLnInvoiceFallback"`
	RedirectAutomatically        bool                 `json:"redirectAutomatically"`
	ShowRecommendedFee           bool                 `json:"showRecommendedFee"`
	RecommendedFeeBlockTarget    int32                `json:"recommendedFeeBlockTarget"`
	DefaultLang                  string               `json:"defaultLang"`
	CustomLogo                   string               `json:"customLogo,omitempty"`
	CustomCSS                    string               `json:"customCSS,omitempty"`
	HtmlTitle                    string               `json:"htmlTitle,omitempty"`
	NetworkFeeMode               BTCPayNetworkFeeMode `json:"networkFeeMode"`
	PayJoinEnabled               bool                 `json:"payJoinEnabled"`
	DefaultPaymentMethod         string               `json:"defaultPaymentMethod"`
	ID                           StoreID              `json:"id"`
}

// View information about the available stores
func (c *Client) GetStores(ctx context.Context) ([]*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores", c.URL)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes []*StoreResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}

	return dataRes, statusCode, nil
}

type StoreRequest struct {
	Name                         string               `json:"name"`
	Website                      string               `json:"website,omitempty"`
	InvoiceExpiration            int64                `json:"invoiceExpiration,omitempty"`
	MonitoringExpiration         int64                `json:"monitoringExpiration,omitempty"`
	SpeedPolicy                  BTCPaySpeedPolicy    `json:"speedPolicy,omitempty"`
	LightningDescriptionTemplate string               `json:"lightningDescriptionTemplate,omitempty"`
	PaymentTolerance             float64              `json:"paymentTolerance,omitempty"`
	AnyoneCanCreateInvoice       bool                 `json:"anyoneCanCreateInvoice,omitempty"`
	RequiresRefundEmail          bool                 `json:"requiresRefundEmail,omitempty"`
	LightningAmountInSatoshi     bool                 `json:"lightningAmountInSatoshi,omitempty"`
	LightningPrivateRouteHints   bool                 `json:"lightningPrivateRouteHints,omitempty"`
	OnChainWithLnInvoiceFallback bool                 `json:"onChainWithLnInvoiceFallback,omitempty"`
	RedirectAutomatically        bool                 `json:"redirectAutomatically,omitempty"`
	ShowRecommendedFee           bool                 `json:"showRecommendedFee,omitempty"`
	RecommendedFeeBlockTarget    int32                `json:"recommendedFeeBlockTarget,omitempty"`
	DefaultLang                  string               `json:"defaultLang,omitempty"`
	CustomLogo                   string               `json:"customLogo,omitempty"`
	CustomCSS                    string               `json:"customCSS,omitempty"`
	HtmlTitle                    string               `json:"htmlTitle,omitempty"`
	NetworkFeeMode               BTCPayNetworkFeeMode `json:"networkFeeMode,omitempty"`
	PayJoinEnabled               bool                 `json:"payJoinEnabled,omitempty"`
	DefaultPaymentMethod         string               `json:"defaultPaymentMethod,omitempty"`
}

// create a new store
func (c *Client) CreateStore(ctx context.Context, storeRequest *StoreRequest) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores", c.URL)
	dataReq, err := json.Marshal(storeRequest)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes StoreResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

// View information about the specified store
func (c *Client) GetStore(ctx context.Context, storeID *StoreID) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", c.URL, *storeID)
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes StoreResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) GetStore(ctx context.Context) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", s.Client.URL, s.ID)
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, nil)
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes StoreResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

type StoreUpdate struct {
	Name                         string               `json:"name,omitempty"`
	Website                      string               `json:"website,omitempty"`
	InvoiceExpiration            int64                `json:"invoiceExpiration,omitempty"`
	MonitoringExpiration         int64                `json:"monitoringExpiration,omitempty"`
	SpeedPolicy                  BTCPaySpeedPolicy    `json:"speedPolicy,omitempty"`
	LightningDescriptionTemplate string               `json:"lightningDescriptionTemplate,omitempty"`
	PaymentTolerance             float64              `json:"paymentTolerance,omitempty"`
	AnyoneCanCreateInvoice       bool                 `json:"anyoneCanCreateInvoice,omitempty"`
	RequiresRefundEmail          bool                 `json:"requiresRefundEmail,omitempty"`
	LightningAmountInSatoshi     bool                 `json:"lightningAmountInSatoshi,omitempty"`
	LightningPrivateRouteHints   bool                 `json:"lightningPrivateRouteHints,omitempty"`
	OnChainWithLnInvoiceFallback bool                 `json:"onChainWithLnInvoiceFallback,omitempty"`
	RedirectAutomatically        bool                 `json:"redirectAutomatically,omitempty"`
	ShowRecommendedFee           bool                 `json:"showRecommendedFee,omitempty"`
	RecommendedFeeBlockTarget    int32                `json:"recommendedFeeBlockTarget,omitempty"`
	DefaultLang                  string               `json:"defaultLang,omitempty"`
	CustomLogo                   string               `json:"customLogo,omitempty"`
	CustomCSS                    string               `json:"customCSS,omitempty"`
	HtmlTitle                    string               `json:"htmlTitle,omitempty"`
	NetworkFeeMode               BTCPayNetworkFeeMode `json:"networkFeeMode,omitempty"`
	PayJoinEnabled               bool                 `json:"payJoinEnabled,omitempty"`
	DefaultPaymentMethod         string               `json:"defaultPaymentMethod,omitempty"`
	ID                           StoreID              `json:"id,omitempty"`
}

func (c *Client) UpdateStore(ctx context.Context, storeID *StoreID, storeUpdate *StoreUpdate) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", c.URL, *storeID)
	dataReq, err := json.Marshal(storeUpdate)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes StoreResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

func (s *Store) UpdateStore(ctx context.Context, storeUpdate *StoreUpdate) (*StoreResponse, int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", s.Client.URL, s.ID)
	dataReq, err := json.Marshal(storeUpdate)
	if err != nil {
		return nil, 0, err
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", endpoint, bytes.NewBuffer(dataReq))
	if err != nil {
		return nil, 0, err
	}
	bytes, statusCode, err := s.Client.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}
	var dataRes StoreResponse
	err = json.Unmarshal(bytes, &dataRes)
	if err != nil {
		return nil, 0, err
	}
	return &dataRes, statusCode, nil
}

// Removes the specified store. If there is another user with access, only your access will be removed.
func (c *Client) RemoveStore(ctx context.Context, storeID *StoreID) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", c.URL, *storeID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return 0, err
	}
	_, statusCode, err := c.doRequest(req)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (s *Store) RemoveStore(ctx context.Context) (int, error) {
	endpoint := fmt.Sprintf("%s/api/v1/stores/%s", s.Client.URL, s.ID)
	req, err := http.NewRequestWithContext(ctx, "DELETE", endpoint, nil)
	if err != nil {
		return 0, err
	}
	_, statusCode, err := s.Client.doRequest(req)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
