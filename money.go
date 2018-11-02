package moneybear

import "github.com/pkg/errors"

// Amount stores the amount as an int64 value.
type Amount struct {
	value int64
}

// Money stores the amount and currency.
type Money struct {
	amount   *Amount
	currency *Currency
}

// New creates a new instance of the Money class.
func New(amount int64, currency string) (*Money, error) {
	curr, err := getCurrencyByCode(currency)
	if err != nil {
		return &Money{}, err
	}

	return &Money{
		&Amount{amount},
		curr,
	}, nil
}

// USD creates a new instances of the Money class with currency USD.
func USD(amount int64) (*Money, error) {
	curr, _ := getCurrencyByCode("USD")
	return createMoney(amount, curr), nil
}

// Amount returns the current amount value for the Money object.
func (m *Money) Amount() int64 {
	return m.amount.value
}

// Currency returns the current currency value for the Money object.
func (m *Money) Currency() *Currency {
	return m.currency
}

// Add will create a new Money instance with the original amount increased by the second.
func (m *Money) Add(m2 *Money) (*Money, error) {
	if err := m.checkCurrencyEqual(m2); err != nil {
		return nil, err
	}

	newAmount := m.amount.value + m2.amount.value

	return createMoney(newAmount, m.currency), nil
}

// Subtract will create a new Money instance with the original amount decreased by the second.
func (m *Money) Subtract(m2 *Money) (*Money, error) {
	if err := m.checkCurrencyEqual(m2); err != nil {
		return nil, err
	}

	newAmount := m.amount.value - m2.amount.value

	return createMoney(newAmount, m.currency), nil
}

// Multiply will create a new Money instance with the original amount multiplied by the multiplier.
func (m *Money) Multiply(multiplier int64) *Money {
	newAmount := m.amount.value * multiplier

	return createMoney(newAmount, m.currency)
}

// Divide will create a new Money instance with the original amount divided by the divisor.
func (m *Money) Divide(divisor int64) *Money {
	newAmount := m.amount.value / divisor

	return createMoney(newAmount, m.currency)
}

// Equals checks if the currency & amount are equal between the two Money instances.
func (m *Money) Equals(m2 *Money) (bool, error) {
	if err := m.checkCurrencyEqual(m2); err != nil {
		return false, err
	}

	return m.amount.value == m2.amount.value, nil
}

// func (m *Money) GreaterThan(m2 *Money) bool {

// }

func (m *Money) checkCurrencyEqual(m2 *Money) error {
	if !m.currency.isEqual(m2.currency) {
		return errors.New("The currency does not match")
	}
	return nil
}

func createMoney(amount int64, curr *Currency) *Money {
	return &Money{
		&Amount{amount},
		curr,
	}
}
