package barion

//go:generate enumer -type=PaymentStatus -json
type PaymentStatus int

const (
	Prepared           PaymentStatus = 10
	Started            PaymentStatus = 20
	InProgress         PaymentStatus = 21
	Waiting            PaymentStatus = 22
	Reserved           PaymentStatus = 25
	Authorized         PaymentStatus = 26
	Canceled           PaymentStatus = 30
	Succeeded          PaymentStatus = 40
	Failed             PaymentStatus = 50
	PartiallySucceeded PaymentStatus = 60
	Expired            PaymentStatus = 70
)
