package btcpay

import "fmt"

// Enums BTCPayPermission
type BTCPayPermission string

type Permission struct {
	Unrestricted                      BTCPayPermission
	UserCanviewprofile                BTCPayPermission
	UserCanmodifyprofile              BTCPayPermission
	UserCanmanagenotificationsforuser BTCPayPermission
	UserCanviewnotificationsforuser   BTCPayPermission

	ServerCancreateuser                         BTCPayPermission
	ServerCanmodifyserversettings               BTCPayPermission
	ServerCanuseinternallightningnode           BTCPayPermission
	ServerCancreatelightninginvoiceinternalnode BTCPayPermission

	StoreCanmodifystoresettings    BTCPayPermission
	StoreWebhooksCanmodifywebhooks BTCPayPermission
	StoreCanviewstoresettings      BTCPayPermission
	StoreCancreateinvoice          BTCPayPermission
	StoreCanviewinvoices           BTCPayPermission
	StoreCanmodifypaymentrequests  BTCPayPermission
	StoreCanviewpaymentrequests    BTCPayPermission
	StoreCanuselightningnode       BTCPayPermission
	StoreCancreatelightninginvoice BTCPayPermission
	CustomPermission               BTCPayPermission
}

func GetPermission() *Permission {
	return &Permission{
		Unrestricted: "unrestricted",

		UserCanviewprofile:                "btcpay.user.canviewprofile",
		UserCanmodifyprofile:              "btcpay.user.canmodifyprofile",
		UserCanmanagenotificationsforuser: "btcpay.user.canmanagenotificationsforuser",
		UserCanviewnotificationsforuser:   "btcpay.user.canviewnotificationsforuser",

		ServerCancreateuser:                         "btcpay.server.cancreateuser",
		ServerCanmodifyserversettings:               "btcpay.server.canmodifyserversettings",
		ServerCanuseinternallightningnode:           "btcpay.server.canuseinternallightningnode",
		ServerCancreatelightninginvoiceinternalnode: "btcpay.server.cancreatelightninginvoiceinternalnode",

		StoreCanmodifystoresettings:    "btcpay.store.canmodifystoresettings",
		StoreWebhooksCanmodifywebhooks: "btcpay.store.webhooks.canmodifywebhooks",
		StoreCanviewstoresettings:      "btcpay.store.canviewstoresettings",
		StoreCancreateinvoice:          "btcpay.store.cancreateinvoice",
		StoreCanviewinvoices:           "btcpay.store.canviewinvoices",
		StoreCanmodifypaymentrequests:  "btcpay.store.canmodifypaymentrequests",
		StoreCanviewpaymentrequests:    "btcpay.store.canviewpaymentrequests",
		StoreCanuselightningnode:       "btcpay.store.canuselightningnode",
		StoreCancreatelightninginvoice: "btcpay.store.cancreatelightninginvoice",
	}
}

func CreateCustomPermission(permission BTCPayPermission, storeID StoreID) BTCPayPermission {
	return BTCPayPermission(fmt.Sprintf("%s:%s", permission, storeID))
}
