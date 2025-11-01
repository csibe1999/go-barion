package barion

type RecurrenceType int

const (
	MerchantInitiatedPayment RecurrenceType = 0
	OneClickPayment          RecurrenceType = 10
	RecurringPayment         RecurrenceType = 20
)

type RecurrenceResult string

const (
	RecurrenceResultNone                          RecurrenceResult = "None"
	RecurrenceResultSuccessful                    RecurrenceResult = "Successful"
	RecurrenceResultFailed                        RecurrenceResult = "Failed"
	RecurrenceResultNotFound                      RecurrenceResult = "NotFound"
	RecurrenceResultThreeDSAuthenticationRequired RecurrenceResult = "ThreeDSAuthenticationRequired"
)
