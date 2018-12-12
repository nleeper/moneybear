package moneybear

import (
	"strings"

	"github.com/pkg/errors"
)

// Currency represents the values making up a currency.
type Currency struct {
	code      string
	precision int
	symbol    string
	decimal   string
	thousand  string
}

var validCurrencies = map[string]*Currency{
	"USD": buildCurrency("USD", 2, "$", ".", ","),
	"GBP": buildCurrency("GBP", 2, "£", ".", ","),
	"EUR": buildCurrency("EUR", 2, "€", ".", ","),
}

func buildCurrency(code string, decimals int, symbol string, decimal string, thousand string) *Currency {
	return &Currency{code, decimals, symbol, decimal, thousand}
}

func getCurrencyByCode(currencyCode string) (*Currency, error) {
	if currency, ok := validCurrencies[strings.ToUpper(currencyCode)]; ok {
		return currency, nil
	}
	return nil, errors.New("Invalid currency code")
}

func (c *Currency) isEqual(c2 *Currency) bool {
	return c.code == c2.code
}
