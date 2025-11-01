package barion

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

// NameInformation describes a person's name block returned by the API.
type NameInformation struct {
	LoginName        string `json:"LoginName,omitempty"`
	FirstName        string `json:"FirstName,omitempty"`
	LastName         string `json:"LastName,omitempty"`
	OrganizationName string `json:"OrganizationName,omitempty"`
}

// UserInformation represents payer/payee information.
type UserInformation struct {
	Name  *NameInformation `json:"Name,omitempty"`
	Email string           `json:"Email,omitempty"`
}

// Item describes a product or service included in a transaction.
type Item struct {
	Name                  string          `json:"Name,omitempty"`
	Description           string          `json:"Description,omitempty"`
	ImageURL              string          `json:"ImageUrl,omitempty"`
	Quantity              decimal.Decimal `json:"Quantity,omitempty"`
	Unit                  string          `json:"Unit,omitempty"`
	UnitPrice             decimal.Decimal `json:"UnitPrice,omitempty"`
	ItemTotal             decimal.Decimal `json:"ItemTotal,omitempty"`
	SKU                   string          `json:"SKU,omitempty"`
	ThreeDSAuthClientData string          `json:"ThreeDSAuthClientData,omitempty"`
	TraceID               string          `json:"TraceId,omitempty"`
}

type PayeeTransaction struct {
	POSTransactionID string          `json:"POSTransactionId,omitempty"`
	Payee            string          `json:"Payee,omitempty"`
	Total            decimal.Decimal `json:"Total,omitempty"`
	Comment          string          `json:"Comment,omitempty"`
}

// PaymentTransaction captures the fields required when creating a payment via the API.
type PaymentTransaction struct {
	POSTransactionID  string             `json:"POSTransactionId"`
	Payee             string             `json:"Payee"`
	Total             decimal.Decimal    `json:"Total"`
	Comment           string             `json:"Comment,omitempty"`
	PayeeTransactions []PayeeTransaction `json:"PayeeTransactions,omitempty"`
	Items             []Item             `json:"Items"`
}

// ProcessedTransaction represents the API response payload for processed transactions.
type ProcessedTransaction struct {
	POSTransactionID string            `json:"POSTransactionId,omitempty"`
	TransactionID    string            `json:"TransactionId,omitempty"`
	Status           TransactionStatus `json:"Status,omitempty"`
	Currency         Currency          `json:"Currency,omitempty"`
	TransactionTime  string            `json:"TransactionTime,omitempty"`
	RelatedID        *string           `json:"RelatedId,omitempty"`
}

func (pt *ProcessedTransaction) UnmarshalJSON(data []byte) error {
	type processedTransactionAlias struct {
		POSTransactionID string            `json:"POSTransactionId,omitempty"`
		TransactionID    string            `json:"TransactionId,omitempty"`
		Status           TransactionStatus `json:"Status,omitempty"`
		Currency         Currency          `json:"Currency,omitempty"`
		TransactionTime  string            `json:"TransactionTime,omitempty"`
	}

	var aux struct {
		processedTransactionAlias
		RelatedID *json.RawMessage `json:"RelatedId"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	pt.POSTransactionID = aux.POSTransactionID
	pt.TransactionID = aux.TransactionID
	pt.Status = aux.Status
	pt.Currency = aux.Currency
	pt.TransactionTime = aux.TransactionTime

	if aux.RelatedID == nil {
		pt.RelatedID = nil
		return nil
	}

	raw := *aux.RelatedID
	if string(raw) == "null" {
		pt.RelatedID = nil
		return nil
	}

	if len(raw) > 0 && raw[0] == '"' {
		var related string
		if err := json.Unmarshal(raw, &related); err != nil {
			return err
		}
		pt.RelatedID = &related
		return nil
	}

	var number json.Number
	if err := json.Unmarshal(raw, &number); err != nil {
		return err
	}

	related := number.String()
	pt.RelatedID = &related

	return nil
}

// DetailedPaymentTransaction mirrors the documentation for PaymentState responses.
type DetailedPaymentTransaction struct {
	TransactionID    string            `json:"TransactionId,omitempty"`
	POSTransactionID string            `json:"POSTransactionId,omitempty"`
	TransactionTime  *time.Time        `json:"TransactionTime,omitempty"`
	Total            decimal.Decimal   `json:"Total,omitempty"`
	Currency         Currency          `json:"Currency,omitempty"`
	Payer            *UserInformation  `json:"Payer,omitempty"`
	Payee            *UserInformation  `json:"Payee,omitempty"`
	Comment          string            `json:"Comment,omitempty"`
	Status           TransactionStatus `json:"Status,omitempty"`
	TransactionType  TransactionType   `json:"TransactionType,omitempty"`
	Items            []Item            `json:"Items,omitempty"`
	RelatedID        *string           `json:"RelatedId,omitempty"`
}
