package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	COP = "COP"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, COP:
		return true
	}
	return false
}
