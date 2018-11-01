package moneybear

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
	return &Money{
		&Amount{amount},
		curr,
	}, nil
}

// Amount returns the current amount value for the Money object.
func (m *Money) Amount() int64 {
	return m.amount.value
}

// Currency returns the current currency value for the Money object.
func (m *Money) Currency() *Currency {
	return m.currency
}

// Add will increase the amount stored in the current Money object.
func (m *Money) Add(amount int64) *Money {
	m.amount.value += amount
	return m
}
