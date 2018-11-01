package moneybear

import (
	"github.com/pkg/errors"
)

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
	return createMoney(createAmount(amount), curr), nil
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

	return createMoney(m.amount.add(m2.amount), m.currency), nil
}

// Subtract will create a new Money instance with the original amount decreased by the second.
func (m *Money) Subtract(m2 *Money) (*Money, error) {
	if err := m.checkCurrencyEqual(m2); err != nil {
		return nil, err
	}

	return createMoney(m.amount.subtract(m2.amount), m.currency), nil
}

// Multiply will create a new Money instance with the original amount multiplied by the multiplier. Rounding is applied, so precision may be lost in the process.
func (m *Money) Multiply(multiplier int64) *Money {
	return createMoney(m.amount.multiply(multiplier), m.currency)
}

// Divide will create a new Money instance with the original amount divided by the divisor. Rounding is applied, so precision may be lost in the process.
func (m *Money) Divide(divisor int64) *Money {
	return createMoney(m.amount.divide(divisor), m.currency)
}

// Allocate will allocate the amount of a Money object according to a list of ratios. The remainder will be split as evenly as possible.
func (m *Money) Allocate(ratios ...int) ([]*Money, error) {
	if len(ratios) == 0 {
		return nil, errors.New("No ratios were supplied")
	}

	var monies []*Money

	total := 0
	for _, r := range ratios {
		total += r
	}

	remainder := m.Amount()

	for _, r := range ratios {
		allocated := m.amount.allocate(r, total)
		remainder -= allocated.value
		monies = append(monies, createMoney(allocated, m.currency))
	}

	for i := 0; remainder > 0; i++ {
		monies[i], _ = monies[i].Add(createMoney(createAmount(1), m.currency))
		remainder--
	}

	return monies, nil
}

// Percentage returns a new Money instance that is a percentage of this. Rounding is applied, so precision may be lost in the process.
func (m *Money) Percentage(percent int64) (*Money, error) {
	if percent < 0 || percent > 100 {
		return nil, errors.New("The percentage must be between 0 and 100")
	}

	return createMoney(m.amount.percentage(percent), m.currency), nil
}

// Equals checks if the currency & amount are equal between the two Money instances.
func (m *Money) Equals(m2 *Money) (bool, error) {
	if err := m.checkCurrencyEqual(m2); err != nil {
		return false, err
	}

	return m.amount.value == m2.amount.value, nil
}

// GreaterThan checks if this is greater than the other Money object.
func (m *Money) GreaterThan(m2 *Money) (bool, error) {
	if err := m.checkCurrencyEqual(m2); err != nil {
		return false, err
	}

	return m.amount.value > m2.amount.value, nil

}

// GreaterThanOrEqual checks if this is greater than or equal to the other Money object.
func (m *Money) GreaterThanOrEqual(m2 *Money) (bool, error) {
	if err := m.checkCurrencyEqual(m2); err != nil {
		return false, err
	}

	return m.amount.value >= m2.amount.value, nil

}

// LessThan checks if this is less than the other Money object.
func (m *Money) LessThan(m2 *Money) (bool, error) {
	if err := m.checkCurrencyEqual(m2); err != nil {
		return false, err
	}

	return m.amount.value < m2.amount.value, nil

}

// LessThanOrEqual checks if this is less than or equal to the other Money object.
func (m *Money) LessThanOrEqual(m2 *Money) (bool, error) {
	if err := m.checkCurrencyEqual(m2); err != nil {
		return false, err
	}

	return m.amount.value <= m2.amount.value, nil

}

func (m *Money) checkCurrencyEqual(m2 *Money) error {
	if !m.currency.isEqual(m2.currency) {
		return errors.New("The currency does not match")
	}
	return nil
}

func createMoney(amount *Amount, curr *Currency) *Money {
	return &Money{
		amount,
		curr,
	}
}
