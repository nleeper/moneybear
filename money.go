package moneybear

// Money is the main class.
type Money struct {
	amount   int
	currency string
}

// New creates a new instance of the Money class.
func New(amount int, currency string) (*Money, error) {
	// TODO - validate currency code, and amount?
	return &Money{
		amount,
		currency,
	}, nil
}

// USD creates a new instances of the Money class with currency USD.
func USD(amount int) (*Money, error) {
	return &Money{
		amount,
		"USD",
	}, nil
}
