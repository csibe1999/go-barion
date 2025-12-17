package barion

// FundingSources represents the possible funding sources for a Barion payment.
type FundingSources string

const (
	All          FundingSources = "All"
	Balance      FundingSources = "Balance"
	BankCard     FundingSources = "BankCard"
	GooglePay    FundingSources = "GooglePay"
	ApplePay     FundingSources = "ApplePay"
	BankTransfer FundingSources = "BankTransfer"
)
