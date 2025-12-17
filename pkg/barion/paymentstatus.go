package barion

// PaymentStatus enumerates the possible payment states as defined by the Barion API.
type PaymentStatus string

const (
	Prepared           PaymentStatus = "Prepared"
	Started            PaymentStatus = "Started"
	InProgress         PaymentStatus = "InProgress"
	Waiting            PaymentStatus = "Waiting"
	Reserved           PaymentStatus = "Reserved"
	Authorized         PaymentStatus = "Authorized"
	Canceled           PaymentStatus = "Canceled"
	Succeeded          PaymentStatus = "Succeeded"
	Failed             PaymentStatus = "Failed"
	PartiallySucceeded PaymentStatus = "PartiallySucceeded"
	Expired            PaymentStatus = "Expired"
)
