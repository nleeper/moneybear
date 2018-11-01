package moneybear

import (
	"strings"

	"github.com/pkg/errors"
)

// Currency represents the values making up a currency.
type Currency struct {
	code     string
	decimals int
}

var validCurrencies = map[string]*Currency{
	"USD": &Currency{"USD", 2},
}

func getCurrencyByCode(currencyCode string) (*Currency, error) {
	if currency, ok := validCurrencies[strings.ToUpper(currencyCode)]; ok {
		return currency, nil
	}
	return nil, errors.New("Invalid currency code")
}
