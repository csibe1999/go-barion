package barion

// PaymentStatus enumerates the possible payment states as defined by the Barion API.
type PaymentStatus string

const (
	PaymentStatusPrepared           PaymentStatus = "Prepared"
	PaymentStatusStarted            PaymentStatus = "Started"
	PaymentStatusInProgress         PaymentStatus = "InProgress"
	PaymentStatusWaiting            PaymentStatus = "Waiting"
	PaymentStatusReserved           PaymentStatus = "Reserved"
	PaymentStatusAuthorized         PaymentStatus = "Authorized"
	PaymentStatusCanceled           PaymentStatus = "Canceled"
	PaymentStatusSucceeded          PaymentStatus = "Succeeded"
	PaymentStatusFailed             PaymentStatus = "Failed"
	PaymentStatusPartiallySucceeded PaymentStatus = "PartiallySucceeded"
	PaymentStatusExpired            PaymentStatus = "Expired"
	PaymentStatusError              PaymentStatus = "Error"
)
