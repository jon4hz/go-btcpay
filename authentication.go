package btcpay

import "fmt"

type Authentication string

const (
	Unrestricted                            Authentication = "unrestricted"
	BTCPayUserCanviewprofile                Authentication = "btcpay.user.canviewprofile"
	BTCPayUserCanmodifyprofile              Authentication = "btcpay.user.canmodifyprofile"
	BTCPayUserCanmanagenotificationsforuser Authentication = "btcpay.user.canmanagenotificationsforuser"
	BTCPayUserCanviewnotificationsforuser   Authentication = "btcpay.user.canviewnotificationsforuser"

	BTCPayServerCancreateuser                         Authentication = "btcpay.server.cancreateuser"
	BTCPayServerCanmodifyserversettings               Authentication = "btcpay.server.canmodifyserversettings"
	BTCPayServerCanuseinternallightningnode           Authentication = "btcpay.server.canuseinternallightningnode"
	BTCPayServerCancreatelightninginvoiceinternalnode Authentication = "btcpay.server.cancreatelightninginvoiceinternalnode"

	BTCPayStoreCanmodifystoresettings    Authentication = "btcpay.store.canmodifystoresettings"
	BTCPayStoreWebhooksCanmodifywebhooks Authentication = "btcpay.store.webhooks.canmodifywebhooks"
	BTCPayStoreCanviewstoresettings      Authentication = "btcpay.store.canviewstoresettings"
	BTCPayStoreCancreateinvoice          Authentication = "btcpay.store.cancreateinvoice"
	BTCPayStoreCanviewinvoices           Authentication = "btcpay.store.canviewinvoices"
	BTCPayStoreCanmodifypaymentrequests  Authentication = "btcpay.store.canmodifypaymentrequests"
	BTCPayStoreCanviewpaymentrequests    Authentication = "btcpay.store.canviewpaymentrequests"
	BTCPayStoreCanuselightningnode       Authentication = "btcpay.store.canuselightningnode"
	BTCPayStoreCancreatelightninginvoice Authentication = "btcpay.store.cancreatelightninginvoice"
)

func (c *Client) CreateRestrictedKey(permission Authentication, apistoreID string) Authentication {
	return Authentication(fmt.Sprintf("%s:%s", permission, apistoreID))
}
