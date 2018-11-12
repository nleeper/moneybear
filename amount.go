package moneybear

import "math"

// Amount stores the amount as an int64 value.
type Amount struct {
	value int64
}

func (amt *Amount) add(amt2 *Amount) *Amount {
	return createAmount(amt.value + amt2.value)
}

func (amt *Amount) subtract(amt2 *Amount) *Amount {
	return createAmount(amt.value - amt2.value)
}

func (amt *Amount) multiply(multiplier int64) *Amount {
	return createAmount(round(float64(amt.value) * float64(multiplier)))
}

func (amt *Amount) divide(divisor int64) *Amount {
	return createAmount(round(float64(amt.value) / float64(divisor)))
}

func (amt *Amount) percentage(percent int64) *Amount {
	return amt.multiply(percent).divide(100)
}

func (amt *Amount) allocate(ratio, total int) *Amount {
	return createAmount(int64(math.Floor((float64(amt.value) * float64(ratio)) / float64(total))))
}

func round(val float64) int64 {
	return int64(math.RoundToEven(val))
}

func absolute(amount int64) int64 {
	if amount >= 0 {
		return amount
	}
	return -amount
}

func createAmount(val int64) *Amount {
	return &Amount{val}
}
