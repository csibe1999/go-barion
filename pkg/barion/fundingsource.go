package barion

//go:generate enumer -type=FundingSource -json
type FundingSources int

const (
	All FundingSources = iota
	Balance
	BankCard
	GooglePay
	ApplePay
	BankTransfer
)
