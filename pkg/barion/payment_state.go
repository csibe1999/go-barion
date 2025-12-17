package barion

import (
	"time"

	"github.com/shopspring/decimal"
)

type paymentStateRequest struct {
	POSKey    string `json:"POSKey,omitempty"`
	PaymentID string `json:"PaymentId"`
}

// CardType represents the bank card type as defined by the Barion API.
type CardType string

const (
	CardTypeUnknown         CardType = "Unknown"
	CardTypeMasterCard      CardType = "MasterCard"
	CardTypeVisa            CardType = "Visa"
	CardTypeAmericanExpress CardType = "AmericanExpress"
	CardTypeElectron        CardType = "Electron"
	CardTypeMaestro         CardType = "Maestro"
	CardTypeDinersClub      CardType = "DinersClub"
	CardTypeDiscover        CardType = "Discover"
)

type FundingBankCard struct {
	MaskedPan      string   `json:"MaskedPan,omitempty"`
	BankCardType   CardType `json:"BankCardType,omitempty"`
	ValidThruYear  string   `json:"ValidThruYear,omitempty"`
	ValidThruMonth string   `json:"ValidThruMonth,omitempty"`
}

type FundingInformation struct {
	BankCard          *FundingBankCard  `json:"BankCard,omitempty"`
	AuthorizationCode string            `json:"AuthorizationCode,omitempty"`
	ProcessResult     CardProcessResult `json:"ProcessResult,omitempty"`
}

// CardProcessResult represents the documented outcomes of a card funding attempt.
type CardProcessResult string

const (
	CardProcessResultSuccessful                 CardProcessResult = "Successful"
	CardProcessResultProblemWithCard            CardProcessResult = "ProblemWithCard"
	CardProcessResultLowFunds                   CardProcessResult = "LowFunds"
	CardProcessResultLostOrStolenCard           CardProcessResult = "LostOrStolenCard"
	CardProcessResultDeclined                   CardProcessResult = "Declined"
	CardProcessResultFraudulentTransaction      CardProcessResult = "FraudulentTransaction"
	CardProcessResultCardSystemError            CardProcessResult = "CardSystemError"
	CardProcessResultScaSoftDeclined            CardProcessResult = "ScaSoftDeclined"
	CardProcessResultFailedDuringAuthentication CardProcessResult = "FailedDuringAuthentication"
	CardProcessResultCardNotSupported           CardProcessResult = "CardNotSupported"
	CardProcessResultUnknown                    CardProcessResult = "Unknown"
	CardProcessResultThreeDsNotEnabled          CardProcessResult = "ThreeDsNotEnabled"
)

type PaymentState struct {
	PaymentID             string                       `json:"PaymentId"`
	PaymentRequestID      string                       `json:"PaymentRequestId"`
	OrderNumber           *string                      `json:"OrderNumber,omitempty"`
	POSId                 string                       `json:"POSId,omitempty"`
	POSName               string                       `json:"POSName,omitempty"`
	POSOwnerEmail         string                       `json:"POSOwnerEmail,omitempty"`
	POSOwnerCountry       string                       `json:"POSOwnerCountry,omitempty"`
	Status                PaymentStatus                `json:"Status,omitempty"`
	PaymentType           PaymentType                  `json:"PaymentType,omitempty"`
	AllowedFundingSources []FundingSources             `json:"AllowedFundingSources,omitempty"`
	FundingSource         *FundingSources              `json:"FundingSource,omitempty"`
	FundingInformation    *FundingInformation          `json:"FundingInformation,omitempty"`
	RecurrenceType        RecurrenceType               `json:"RecurrenceType,omitempty"`
	TraceID               string                       `json:"TraceId,omitempty"`
	GuestCheckout         bool                         `json:"GuestCheckout,omitempty"`
	CreatedAt             *time.Time                   `json:"CreatedAt,omitempty"`
	CompletedAt           *time.Time                   `json:"CompletedAt,omitempty"`
	ValidUntil            *time.Time                   `json:"ValidUntil,omitempty"`
	ReservedUntil         *time.Time                   `json:"ReservedUntil,omitempty"`
	DelayedCaptureUntil   *time.Time                   `json:"DelayedCaptureUntil,omitempty"`
	Transactions          []DetailedPaymentTransaction `json:"Transactions,omitempty"`
	Total                 decimal.Decimal              `json:"Total,omitempty"`
	Currency              Currency                     `json:"Currency,omitempty"`
	SuggestedLocale       *Locale                      `json:"SuggestedLocale,omitempty"`
	FraudRiskScore        *int                         `json:"FraudRiskScore,omitempty"`
	CallbackURL           *string                      `json:"CallbackUrl,omitempty"`
	RedirectURL           *string                      `json:"RedirectUrl,omitempty"`
	Errors                []barionError                `json:"Errors,omitempty"`
	PaymentMethod         PaymentMethod                `json:"PaymentMethod,omitempty"`
}
