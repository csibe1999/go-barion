package barion

type PaymentRequestResponse struct {
	Errors                []barionError          `json:"Errors,omitempty"`
	PaymentID             string                 `json:"PaymentId,omitempty"`
	PaymentRequestID      string                 `json:"PaymentRequestId,omitempty"`
	Status                PaymentStatus          `json:"Status,omitempty"`
	QRUrl                 string                 `json:"QRUrl,omitempty"`
	RecurrenceResult      RecurrenceResult       `json:"RecurrenceResult,omitempty"`
	Transactions          []ProcessedTransaction `json:"Transactions,omitempty"`
	GatewayURL            string                 `json:"GatewayUrl,omitempty"`
	CallbackURL           string                 `json:"CallbackUrl,omitempty"`
	RedirectURL           string                 `json:"RedirectUrl,omitempty"`
	ThreeDSAuthClientData string                 `json:"ThreeDSAuthClientData,omitempty"`
	TraceID               string                 `json:"TraceId,omitempty"`
}
