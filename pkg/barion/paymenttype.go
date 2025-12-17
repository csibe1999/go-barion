package barion

// PaymentType enumerates the possible payment types as defined by the Barion API.
type PaymentType string

const (
	Immediate      PaymentType = "Immediate"
	Reservation    PaymentType = "Reservation"
	DelayedCapture PaymentType = "DelayedCapture"
)
