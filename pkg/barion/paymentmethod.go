package barion

import (
	"encoding/json"
	"fmt"
)

// PaymentMethod represents the documented payment completion methods.
type PaymentMethod int

const (
	PaymentMethodUnknown      PaymentMethod = 0
	PaymentMethodBarionWallet PaymentMethod = 10
	PaymentMethodBankCard     PaymentMethod = 20
	PaymentMethodGooglePay    PaymentMethod = 30
	PaymentMethodApplePay     PaymentMethod = 40
	PaymentMethodOpenBanking  PaymentMethod = 50
	PaymentMethodOther        PaymentMethod = 200
)

var paymentMethodToString = map[PaymentMethod]string{
	PaymentMethodUnknown:      "Unknown",
	PaymentMethodBarionWallet: "BarionWallet",
	PaymentMethodBankCard:     "BankCard",
	PaymentMethodGooglePay:    "GooglePay",
	PaymentMethodApplePay:     "ApplePay",
	PaymentMethodOpenBanking:  "OpenBanking",
	PaymentMethodOther:        "Other",
}

var stringToPaymentMethod = map[string]PaymentMethod{
	"Unknown":      PaymentMethodUnknown,
	"BarionWallet": PaymentMethodBarionWallet,
	"BankCard":     PaymentMethodBankCard,
	"GooglePay":    PaymentMethodGooglePay,
	"ApplePay":     PaymentMethodApplePay,
	"OpenBanking":  PaymentMethodOpenBanking,
	"Other":        PaymentMethodOther,
}

func (m PaymentMethod) String() string {
	if val, ok := paymentMethodToString[m]; ok {
		return val
	}
	return fmt.Sprintf("PaymentMethod(%d)", m)
}

// MarshalJSON serialises the payment method to its documented string representation.
func (m PaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

// UnmarshalJSON deserialises the payment method from a documented string representation.
func (m *PaymentMethod) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return fmt.Errorf("PaymentMethod should be a string, got %s", data)
	}
	if val, ok := stringToPaymentMethod[value]; ok {
		*m = val
		return nil
	}
	return fmt.Errorf("%s does not belong to PaymentMethod values", value)
}

// PaymentMethodValues returns every supported payment method.
func PaymentMethodValues() []PaymentMethod {
	return []PaymentMethod{
		PaymentMethodUnknown,
		PaymentMethodBarionWallet,
		PaymentMethodBankCard,
		PaymentMethodGooglePay,
		PaymentMethodApplePay,
		PaymentMethodOpenBanking,
		PaymentMethodOther,
	}
}
