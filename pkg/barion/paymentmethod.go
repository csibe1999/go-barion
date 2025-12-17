package barion

// PaymentMethod represents the documented payment completion methods.
type PaymentMethod string

const (
	PaymentMethodUnknown      PaymentMethod = "Unknown"
	PaymentMethodBarionWallet PaymentMethod = "BarionWallet"
	PaymentMethodBankCard     PaymentMethod = "BankCard"
	PaymentMethodGooglePay    PaymentMethod = "GooglePay"
	PaymentMethodApplePay     PaymentMethod = "ApplePay"
	PaymentMethodOpenBanking  PaymentMethod = "OpenBanking"
	PaymentMethodOther        PaymentMethod = "Other"
)
