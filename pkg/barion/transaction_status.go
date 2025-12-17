package barion

// TransactionStatus enumerates the possible transaction states as defined by the Barion API.
type TransactionStatus string

const (
	TransactionStatusPrepared                 TransactionStatus = "Prepared"
	TransactionStatusStarted                  TransactionStatus = "Started"
	TransactionStatusSucceeded                TransactionStatus = "Succeeded"
	TransactionStatusTimeout                  TransactionStatus = "Timeout"
	TransactionStatusShopIsDeleted            TransactionStatus = "ShopIsDeleted"
	TransactionStatusShopIsClosed             TransactionStatus = "ShopIsClosed"
	TransactionStatusRejected                 TransactionStatus = "Rejected"
	TransactionStatusRejectedByShop           TransactionStatus = "RejectedByShop"
	TransactionStatusStorno                   TransactionStatus = "Storno"
	TransactionStatusReserved                 TransactionStatus = "Reserved"
	TransactionStatusDeleted                  TransactionStatus = "Deleted"
	TransactionStatusExpired                  TransactionStatus = "Expired"
	TransactionStatusAuthorized               TransactionStatus = "Authorized"
	TransactionStatusReversed                 TransactionStatus = "Reversed"
	TransactionStatusInvalidPaymentRecord     TransactionStatus = "InvalidPaymentRecord"
	TransactionStatusPaymentTimeout           TransactionStatus = "PaymentTimeOut"
	TransactionStatusInvalidPaymentStatus     TransactionStatus = "InvalidPaymentStatus"
	TransactionStatusSenderOrRecipientInvalid TransactionStatus = "PaymentSenderOrRecipientIsInvalid"
	TransactionStatusUnknown                  TransactionStatus = "Unknown"
)
