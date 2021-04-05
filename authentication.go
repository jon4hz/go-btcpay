package btcpay

import "fmt"

type Permission string

const (
	Unrestricted                            Permission = "unrestricted"
	BTCPayUserCanviewprofile                Permission = "btcpay.user.canviewprofile"
	BTCPayUserCanmodifyprofile              Permission = "btcpay.user.canmodifyprofile"
	BTCPayUserCanmanagenotificationsforuser Permission = "btcpay.user.canmanagenotificationsforuser"
	BTCPayUserCanviewnotificationsforuser   Permission = "btcpay.user.canviewnotificationsforuser"

	BTCPayServerCancreateuser                         Permission = "btcpay.server.cancreateuser"
	BTCPayServerCanmodifyserversettings               Permission = "btcpay.server.canmodifyserversettings"
	BTCPayServerCanuseinternallightningnode           Permission = "btcpay.server.canuseinternallightningnode"
	BTCPayServerCancreatelightninginvoiceinternalnode Permission = "btcpay.server.cancreatelightninginvoiceinternalnode"

	BTCPayStoreCanmodifystoresettings    Permission = "btcpay.store.canmodifystoresettings"
	BTCPayStoreWebhooksCanmodifywebhooks Permission = "btcpay.store.webhooks.canmodifywebhooks"
	BTCPayStoreCanviewstoresettings      Permission = "btcpay.store.canviewstoresettings"
	BTCPayStoreCancreateinvoice          Permission = "btcpay.store.cancreateinvoice"
	BTCPayStoreCanviewinvoices           Permission = "btcpay.store.canviewinvoices"
	BTCPayStoreCanmodifypaymentrequests  Permission = "btcpay.store.canmodifypaymentrequests"
	BTCPayStoreCanviewpaymentrequests    Permission = "btcpay.store.canviewpaymentrequests"
	BTCPayStoreCanuselightningnode       Permission = "btcpay.store.canuselightningnode"
	BTCPayStoreCancreatelightninginvoice Permission = "btcpay.store.cancreatelightninginvoice"
)

func (c *Client) CreateRestrictedKey(permission Permission, storeID StoreID) Permission {
	return Permission(fmt.Sprintf("%s:%s", permission, storeID))
}
