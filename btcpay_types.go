package btcpay

import "net/http"

type APIKey string

type Client struct {
	URL            string
	APIKey         APIKey
	Username       string
	Password       string
	Http           *http.Client
	Store          *Store
	Invoice        *Invoice
	PaymentRequest *PaymentRequest
	Notification   *Notification
	PullPayment    *PullPayment
	Payout         *Payout
}

type APIKeyRequest struct {
	Label       string             `json:"label,omitempty"`
	Permissions []BTCPayPermission `json:"permissions,omitempty"`
}
type APIKeyResponse struct {
	APIKey      APIKey             `json:"apiKey"`
	Label       string             `json:"label"`
	Permissions []BTCPayPermission `json:"permissions"`
}

type AuthorizationRequest struct {
	Permissions           []BTCPayPermission `json:"permissions,omitempty"`
	ApplicationName       string             `json:"applicationName,omitempty"`
	Strict                bool               `json:"strict,omitempty"`
	SelectiveStores       bool               `json:"selectiveStores,omitempty"`
	Redirect              string             `json:"redirect,omitempty"`
	ApplicationIdentifier string             `json:"applicationIdentifier,omitempty"`
}

type HealthResponse struct {
	Synchronized bool `json:"synchronized"`
}

type InvoiceID string

type Invoice struct {
	Store  *Store
	Client *Client
	ID     InvoiceID
}

type InvoiceResponse struct {
	Amount           string                        `json:"amount,omitempty"`
	Currency         string                        `json:"currency,omitempty"`
	Metadata         InvoiceMetadata               `json:"metadata,omitempty"`
	Checkout         InvoiceCheckout               `json:"checkout,omitempty"`
	ID               InvoiceID                     `json:"id"`
	CheckoutLink     string                        `json:"checkoutLink"`
	CreatedTime      int64                         `json:"createdTime"`
	ExpirationTime   int64                         `json:"expirationTime"`
	MonitoringTime   int64                         `json:"monitoringTime"`
	Status           BTCPayInvoiceStatus           `json:"status"`
	AdditionalStatus BTCPayInvoiceAdditionalStatus `json:"additionalStatus"`
}

type InvoiceMetadata map[string]interface{}

type InvoiceCheckout struct {
	SpeedPolicy       BTCPaySpeedPolicy `json:"speedPolicy,omitempty"`
	PaymentMethods    []string          `json:"paymentMethods,omitempty"`
	ExpirationMinutes int               `json:"expirationMinutes,omitempty"`
	MonitoringMinutes int               `json:"monitoringMinutes,omitempty"`
	PaymentTolerance  float64           `json:"paymentTolerance,omitempty"`
	RedirectURL       string            `json:"redirectURL,omitempty"`
	DefaultLanguage   string            `json:"defaultLanguage,omitempty"`
}

type InvoiceRequest struct {
	Amount          string          `json:"amount"`
	Currency        string          `json:"currency,omitempty"`
	Metadata        InvoiceMetadata `json:"metadata,omitempty"`
	InvoiceCheckout InvoiceCheckout `json:"checkout,omitempty"`
}

type InvoiceUpdate struct {
	Metadata InvoiceMetadata `json:"metadata,omitempty"`
}

type PaymentID string

type Payment struct {
	ID           PaymentID           `json:"id"`
	ReceivedDate int64               `json:"receivedDate"`
	Value        string              `json:"value"`
	Fee          string              `json:"fee"`
	Status       BTCPayPaymentStatus `json:"status"`
	Destination  string              `json:"destination"`
}

type InvoicePaymentMethodResponse struct {
	PaymentMethod     string    `json:"paymentMethod"`
	Destination       string    `json:"destination"`
	PaymentLink       string    `json:"paymentLink,omitempty"`
	Rate              string    `json:"rate"`
	PaymentMethodPaid string    `json:"paymentMethodPaid"`
	TotalPaid         string    `json:"totalPaid"`
	Due               string    `json:"due"`
	Amount            string    `json:"amount"`
	NetworkFee        string    `json:"networkFee"`
	Payments          []Payment `json:"payments"`
}

type MarkInvoiceStatusRequest struct {
	Status BTCPayInvoiceStatusMark `json:"status"`
}

type LanguageCodesRespose struct {
	Code            string `json:"code"`
	CurrentLanguage string `json:"currentLanguage"`
}

type InvoiceCheckoutPage struct {
	Page []byte
}

type Notification struct {
	ID     NotificationID
	Client *Client
}

type NotificationID string

type NotificationResponse struct {
	ID          NotificationID `json:"id"`
	Body        string         `json:"body"`
	Link        string         `json:"link"`
	CreatedTime int64          `json:"createdTime"`
	Seen        bool           `json:"seen"`
}

type UpdateNotification struct {
	Seen bool `json:"seen"`
}

type PaymentRequestID string

type PaymentRequest struct {
	Store  *Store
	Client *Client
	ID     PaymentRequestID
}

type PaymentRequestRequest struct {
	Amount                    float64 `json:"amount"`
	Title                     string  `json:"title"`
	Currency                  string  `json:"currency"`
	Email                     string  `json:"email,omitempty"`
	Description               string  `json:"description,omitempty"`
	ExpiryDate                int64   `json:"expiryDate,omitempty"`
	EmbeddedCSS               string  `json:"embeddedCSS,omitempty"`
	CustomCSSLink             string  `json:"customCSSLink,omitempty"`
	AllowCustomPaymentAmounts bool    `json:"allowCustomPaymentAmounts,omitempty"`
}

type PaymentRequestResponse struct {
	ID       PaymentRequestID           `json:"id"`
	Status   BTCPayPaymentRequestStatus `json:"status"`
	Created  string                     `json:"created"`
	Archived bool                       `json:"archived"`
	PaymentRequestRequest
}

type PullPaymentID string

type PullPayment struct {
	Store  *Store
	Client *Client
	ID     PullPaymentID
}

type PullPaymentResponse struct {
	ID       PullPaymentID `json:"id"`
	Name     string        `json:"name"`
	Currency string        `json:"currency"`
	Amount   string        `json:"amount"`
	Period   int64         `json:"period,omitempty"`
	Archived bool          `json:"archived"`
	ViewLink string        `json:"viewLink"`
}

type PayoutID string

type Payout struct {
	Store  *Store
	Client *Client
	ID     PayoutID
}

type PullPaymentRequest struct {
	Name           string   `json:"name,omitempty"`
	Amount         string   `json:"amount"`
	Currency       string   `json:"currency"`
	Period         int64    `json:"period,omitempty"`
	StartsAt       int64    `json:"startsAt,omitempty"`
	ExpiresAt      int64    `json:"expiresAt,omitempty"`
	PaymentMethods []string `json:"paymentMethods"`
}

type PayoutResponse struct {
	ID                  PayoutID           `json:"id"`
	Revision            int64              `json:"revision"`
	PullPaymentID       PullPaymentID      `json:"pullPaymentId"`
	Date                string             `json:"date"`
	Destination         string             `json:"destination"`
	Amount              string             `json:"amount"`
	PaymentMethod       string             `json:"paymentMethod"`
	PaymentMethodAmount string             `json:"paymentMethodAmount"`
	State               BTCPayPayoutStatus `json:"state"`
}

type PayoutApproveRequest struct {
	Revision int64  `json:"revision"`
	RateRule string `json:"rateRule,omitempty"`
}

type PayoutRequest struct {
	Destination   string `json:"destination"`
	Amount        string `json:"amount"`
	PaymentMethod string `json:"paymentMethod"`
}

type ServerInfoResponse struct {
	Version                 string             `json:"version"`
	Onion                   string             `json:"onion"`
	SupportedPaymentMethods []string           `json:"supportedPaymentMethods"`
	FullySynched            bool               `json:"fullySynched"`
	SyncStatus              []ServerSyncStatus `json:"syncStatus"`
}

type ServerSyncStatus struct {
	CryptoCode      string                `json:"cryptoCode"`
	NodeInformation ServerNodeInformation `json:"nodeInformation,omitempty"`
	ChainHeight     int64                 `json:"chainHeight"`
	SyncHeight      int64                 `json:"syncHeight,omitempty"`
}

type ServerNodeInformation struct {
	Headers              int64   `json:"headers"`
	Blocks               int64   `json:"blocks"`
	VerificationProgress float64 `json:"verificationProgress"`
}

type StoreID string

type Store struct {
	ID     StoreID
	Client *Client
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

type UserID string

type UserResponse struct {
	ID                        UserID   `json:"id"`
	Email                     string   `json:"email"`
	EmailConfirmed            bool     `json:"emailConfirmed"`
	RequiresEmailConfirmation bool     `json:"requiresEmailConfirmation"`
	Created                   int64    `json:"created,omitempty"`
	Roles                     []string `json:"roles"`
}

type UserRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	IsAdministrator bool   `json:"isAdministrator"`
}
