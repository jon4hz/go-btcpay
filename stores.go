package btcpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type StoreID string

type NetworkFeeMode string

const (
	MultiplePaymentsOnly NetworkFeeMode = "MultiplePaymentsOnly"
	Always               NetworkFeeMode = "Always"
	Never                NetworkFeeMode = "Never"
)

type StoreResponse struct {
	Name                         string         `json:"name"`
	Website                      string         `json:"website"`
	InvoiceExpiration            int64          `json:"invoiceExpiration"`
	MonitoringExpiration         int64          `json:"monitoringExpiration"`
	SpeedPolicy                  SpeedPolicy    `json:"speedPolicy"`
	LightningDescriptionTemplate string         `json:"lightningDescriptionTemplate,omitempty"`
	PaymentTolerance             float64        `json:"paymentTolerance"`
	AnyoneCanCreateInvoice       bool           `json:"anyoneCanCreateInvoice"`
	RequiresRefundEmail          bool           `json:"requiresRefundEmail"`
	LightningAmountInSatoshi     bool           `json:"lightningAmountInSatoshi"`
	LightningPrivateRouteHints   bool           `json:"lightningPrivateRouteHints"`
	OnChainWithLnInvoiceFallback bool           `json:"onChainWithLnInvoiceFallback"`
	RedirectAutomatically        bool           `json:"redirectAutomatically"`
	ShowRecommendedFee           bool           `json:"showRecommendedFee"`
	RecommendedFeeBlockTarget    int32          `json:"recommendedFeeBlockTarget"`
	DefaultLang                  string         `json:"defaultLang"`
	CustomLogo                   string         `json:"customLogo,omitempty"`
	CustomCSS                    string         `json:"customCSS,omitempty"`
	HtmlTitle                    string         `json:"htmlTitle,omitempty"`
	NetworkFeeMode               NetworkFeeMode `json:"networkFeeMode"`
	PayJoinEnabled               bool           `json:"payJoinEnabled"`
	DefaultPaymentMethod         string         `json:"defaultPaymentMethod"`
	ID                           StoreID        `json:"id"`
}

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
