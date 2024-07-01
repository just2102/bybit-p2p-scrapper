package data

type Currency string

const (
	AED Currency = "AED"
	AMD Currency = "AMD"
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	AZN Currency = "AZN"
	BDT Currency = "BDT"
	BGN Currency = "BGN"
	UAH Currency = "UAH"
	THB Currency = "THB"
	KWD Currency = "KWD"
	CAD Currency = "CAD"
	JPY Currency = "JPY"
	BYN Currency = "BYN"
	RUB Currency = "RUB"
	USD Currency = "USD"
	VND Currency = "VND"
)

var validCurrencies = map[Currency]bool{
	AED: true,
	AMD: true,
	ARS: true,
	AUD: true,
	AZN: true,
	BDT: true,
	BGN: true,
	UAH: true,
	THB: true,
	KWD: true,
	CAD: true,
	JPY: true,
	BYN: true,
	RUB: true,
	USD: true,
	VND: true,
}

func IsValidCurrency(currency Currency) bool {
	return validCurrencies[currency]
}
