package btcpay

// Enums SpeedPolicy
type BTCPaySpeedPolicy string

type SpeedPolicy struct {
	HighSpeed      BTCPaySpeedPolicy
	MediumSpeed    BTCPaySpeedPolicy
	LowMediumSpeed BTCPaySpeedPolicy
	LowSpeed       BTCPaySpeedPolicy
}

func GetSpeedPolicy() *SpeedPolicy {
	return &SpeedPolicy{
		HighSpeed:      "HighSpeed",
		MediumSpeed:    "MediumSpeed",
		LowMediumSpeed: "LowMediumSpeed",
		LowSpeed:       "LowSpeed",
	}
}

// Enums InvoiceStatus
type BTCPayInvoiceStatus string

type InvoiceStatus struct {
	New        BTCPayInvoiceStatus
	Processing BTCPayInvoiceStatus
	Expired    BTCPayInvoiceStatus
	Invalid    BTCPayInvoiceStatus
	Settled    BTCPayInvoiceStatus
}

func GetInvoiceStatus() *InvoiceStatus {
	return &InvoiceStatus{
		New:        "New",
		Processing: "Processing",
		Expired:    "Expired",
		Invalid:    "Invalid",
		Settled:    "Settled",
	}
}

// Enums InvoiceAdditionalStatus
type BTCPayInvoiceAdditionalStatus string

type InvoiceAdditionalStatus struct {
	None        BTCPayInvoiceAdditionalStatus
	PaidLate    BTCPayInvoiceAdditionalStatus
	PaidPartial BTCPayInvoiceAdditionalStatus
	Marked      BTCPayInvoiceAdditionalStatus
	AddInvalid  BTCPayInvoiceAdditionalStatus
	PaidOver    BTCPayInvoiceAdditionalStatus
}

func GetInvoiceAdditionalStatus() *InvoiceAdditionalStatus {
	return &InvoiceAdditionalStatus{
		None:        "None",
		PaidLate:    "PaidLate",
		PaidPartial: "PaidPartial",
		Marked:      "Marked",
		AddInvalid:  "Invalid",
		PaidOver:    "PaidOver",
	}

}

// Enums InvoiceStatusMark
type BTCPayInvoiceStatusMark string

type InvoiceStatusMark struct {
	MarkInvalid  BTCPayInvoiceStatusMark
	MarkComplete BTCPayInvoiceStatusMark
}

func GetInvoiceStatusMark() *InvoiceStatusMark {
	return &InvoiceStatusMark{
		MarkInvalid:  "Invalid",
		MarkComplete: "Complete",
	}
}

// Enums PaymentStatus
type BTCPayPaymentStatus string

type PaymentStatus struct {
	New        BTCPayPaymentStatus
	Processing BTCPayPaymentStatus
	Expired    BTCPayPaymentStatus
	Invalid    BTCPayPaymentStatus
	Settled    BTCPayPaymentStatus
}

func GetPaymentStatus() *PaymentStatus {
	return &PaymentStatus{
		Processing: "Processing",
		Invalid:    "Invalid",
		Settled:    "Settled",
	}
}

// Enums PaymentRequestStatus
type BTCPayPaymentRequestStatus string

type PaymentRequestStatus struct {
	Pending   BTCPayInvoiceStatus
	Completed BTCPayInvoiceStatus
	Expired   BTCPayInvoiceStatus
}

func GetPaymentRequestStatus() *PaymentRequestStatus {
	return &PaymentRequestStatus{
		Pending:   "Pending",
		Completed: "Completed",
		Expired:   "Expired",
	}
}

// Enums NetworkFeeMode
type BTCPayPayoutStatus string

type PayoutStatus struct {
	AwaitingApproval BTCPayPayoutStatus
	AwaitingPayment  BTCPayPayoutStatus
	InProgress       BTCPayPayoutStatus
	Completed        BTCPayPayoutStatus
	Cancelled        BTCPayPayoutStatus
}

func GetPayoutStatus() *PayoutStatus {
	return &PayoutStatus{
		AwaitingApproval: "AwaitingApproval",
		AwaitingPayment:  "AwaitingPayment",
		InProgress:       "InProgress",
		Completed:        "Completed",
		Cancelled:        "Cancelled",
	}
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
