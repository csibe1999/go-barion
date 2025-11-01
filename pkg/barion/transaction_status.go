package barion

import (
	"encoding/json"
	"fmt"
)

// TransactionStatus enumerates the possible transaction states as defined by the Barion API.
type TransactionStatus int

const (
	TransactionStatusPrepared                 TransactionStatus = 0
	TransactionStatusStarted                  TransactionStatus = 1
	TransactionStatusSucceeded                TransactionStatus = 2
	TransactionStatusTimeout                  TransactionStatus = 3
	TransactionStatusShopIsDeleted            TransactionStatus = 4
	TransactionStatusShopIsClosed             TransactionStatus = 5
	TransactionStatusRejected                 TransactionStatus = 6
	TransactionStatusRejectedByShop           TransactionStatus = 12
	TransactionStatusStorno                   TransactionStatus = 13
	TransactionStatusReserved                 TransactionStatus = 14
	TransactionStatusDeleted                  TransactionStatus = 15
	TransactionStatusExpired                  TransactionStatus = 16
	TransactionStatusAuthorized               TransactionStatus = 17
	TransactionStatusReversed                 TransactionStatus = 18
	TransactionStatusInvalidPaymentRecord     TransactionStatus = 210
	TransactionStatusPaymentTimeout           TransactionStatus = 211
	TransactionStatusInvalidPaymentStatus     TransactionStatus = 212
	TransactionStatusSenderOrRecipientInvalid TransactionStatus = 213
	TransactionStatusUnknown                  TransactionStatus = 255
)

var transactionStatusToString = map[TransactionStatus]string{
	TransactionStatusPrepared:                 "Prepared",
	TransactionStatusStarted:                  "Started",
	TransactionStatusSucceeded:                "Succeeded",
	TransactionStatusTimeout:                  "Timeout",
	TransactionStatusShopIsDeleted:            "ShopIsDeleted",
	TransactionStatusShopIsClosed:             "ShopIsClosed",
	TransactionStatusRejected:                 "Rejected",
	TransactionStatusRejectedByShop:           "RejectedByShop",
	TransactionStatusStorno:                   "Storno",
	TransactionStatusReserved:                 "Reserved",
	TransactionStatusDeleted:                  "Deleted",
	TransactionStatusExpired:                  "Expired",
	TransactionStatusAuthorized:               "Authorized",
	TransactionStatusReversed:                 "Reversed",
	TransactionStatusInvalidPaymentRecord:     "InvalidPaymentRecord",
	TransactionStatusPaymentTimeout:           "PaymentTimeOut",
	TransactionStatusInvalidPaymentStatus:     "InvalidPaymentStatus",
	TransactionStatusSenderOrRecipientInvalid: "PaymentSenderOrRecipientIsInvalid",
	TransactionStatusUnknown:                  "Unknown",
}

var stringToTransactionStatus = map[string]TransactionStatus{
	"Prepared":                          TransactionStatusPrepared,
	"Started":                           TransactionStatusStarted,
	"Succeeded":                         TransactionStatusSucceeded,
	"Timeout":                           TransactionStatusTimeout,
	"ShopIsDeleted":                     TransactionStatusShopIsDeleted,
	"ShopIsClosed":                      TransactionStatusShopIsClosed,
	"Rejected":                          TransactionStatusRejected,
	"RejectedByShop":                    TransactionStatusRejectedByShop,
	"Storno":                            TransactionStatusStorno,
	"Reserved":                          TransactionStatusReserved,
	"Deleted":                           TransactionStatusDeleted,
	"Expired":                           TransactionStatusExpired,
	"Authorized":                        TransactionStatusAuthorized,
	"Reversed":                          TransactionStatusReversed,
	"InvalidPaymentRecord":              TransactionStatusInvalidPaymentRecord,
	"PaymentTimeOut":                    TransactionStatusPaymentTimeout,
	"InvalidPaymentStatus":              TransactionStatusInvalidPaymentStatus,
	"PaymentSenderOrRecipientIsInvalid": TransactionStatusSenderOrRecipientInvalid,
	"Unknown":                           TransactionStatusUnknown,
}

// String converts an enum value to its documented representation.
func (s TransactionStatus) String() string {
	if val, ok := transactionStatusToString[s]; ok {
		return val
	}
	return fmt.Sprintf("TransactionStatus(%d)", s)
}

// MarshalJSON serialises the status as a string value.
func (s TransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// UnmarshalJSON deserialises the status from a documented string value.
func (s *TransactionStatus) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return fmt.Errorf("TransactionStatus should be a string, got %s", data)
	}
	if val, ok := stringToTransactionStatus[value]; ok {
		*s = val
		return nil
	}
	return fmt.Errorf("%s does not belong to TransactionStatus values", value)
}

// TransactionStatusValues returns every supported transaction status.
func TransactionStatusValues() []TransactionStatus {
	return []TransactionStatus{
		TransactionStatusPrepared,
		TransactionStatusStarted,
		TransactionStatusSucceeded,
		TransactionStatusTimeout,
		TransactionStatusShopIsDeleted,
		TransactionStatusShopIsClosed,
		TransactionStatusRejected,
		TransactionStatusRejectedByShop,
		TransactionStatusStorno,
		TransactionStatusReserved,
		TransactionStatusDeleted,
		TransactionStatusExpired,
		TransactionStatusAuthorized,
		TransactionStatusReversed,
		TransactionStatusInvalidPaymentRecord,
		TransactionStatusPaymentTimeout,
		TransactionStatusInvalidPaymentStatus,
		TransactionStatusSenderOrRecipientInvalid,
		TransactionStatusUnknown,
	}
}
