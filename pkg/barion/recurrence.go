package barion

// RecurrenceType represents the recurring payment type as defined by the Barion API.
type RecurrenceType string

const (
	MerchantInitiatedPayment RecurrenceType = "MerchantInitiatedPayment"
	OneClickPayment          RecurrenceType = "OneClickPayment"
	RecurringPayment         RecurrenceType = "RecurringPayment"
)

type RecurrenceResult string

const (
	RecurrenceResultNone                          RecurrenceResult = "None"
	RecurrenceResultSuccessful                    RecurrenceResult = "Successful"
	RecurrenceResultFailed                        RecurrenceResult = "Failed"
	RecurrenceResultNotFound                      RecurrenceResult = "NotFound"
	RecurrenceResultThreeDSAuthenticationRequired RecurrenceResult = "ThreeDSAuthenticationRequired"
)
