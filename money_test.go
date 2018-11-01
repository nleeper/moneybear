package moneybear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var expected int64 = 100

	new, _ := New(expected, "USD")
	assert.Equal(t, expected, new.Amount())
	assert.Equal(t, "USD", new.Currency().code)
}

func TestUSD(t *testing.T) {
	var expected int64 = 101

	new, _ := USD(expected)
	assert.Equal(t, int64(expected), new.Amount())
	assert.Equal(t, "USD", new.Currency().code)
}

func TestInvalidCurrency(t *testing.T) {
	_, err := New(100, "USX")
	assert.EqualError(t, err, "Invalid currency code")
}

func TestAdd(t *testing.T) {
	var expected int64 = 300

	new := createUSD(expected)
	assert.Equal(t, expected, new.Amount())

	new2, _ := new.Add(createUSD(130))

	assert.Equal(t, expected+130, new2.Amount())
}

func TestAddDifferentCurrency(t *testing.T) {
	new := createUSD(100)

	_, err := new.Add(create(50, "GBP"))
	assert.EqualError(t, err, "The currency does not match")
}

func TestSubtract(t *testing.T) {
	var expected int64 = 300

	new := createUSD(expected)
	assert.Equal(t, expected, new.Amount())

	new2, _ := new.Subtract(createUSD(130))

	assert.Equal(t, expected-130, new2.Amount())
}

func TestSubtractDifferentCurrency(t *testing.T) {
	new := createUSD(100)

	_, err := new.Subtract(create(50, "GBP"))
	assert.EqualError(t, err, "The currency does not match")
}

func TestMultiply(t *testing.T) {
	var expected int64 = 300

	new := createUSD(expected)
	assert.Equal(t, expected, new.Amount())

	new2 := new.Multiply(3)

	assert.Equal(t, expected*3, new2.Amount())
	assert.Equal(t, new.currency, new2.currency)
}

func TestDivide(t *testing.T) {
	var expected int64 = 600

	new := createUSD(expected)
	assert.Equal(t, expected, new.Amount())

	assert.Equal(t, int64(29), new.Divide(21).Amount())
	assert.Equal(t, int64(27), new.Divide(22).Amount())
}

func TestPercentage(t *testing.T) {
	new := createUSD(9100)

	per, _ := new.Percentage(34)
	assert.Equal(t, int64(3094), per.Amount())
}

func TestPercentageOutOfRange(t *testing.T) {
	new := createUSD(1000)

	_, err := new.Percentage(-1)
	assert.EqualError(t, err, "The percentage must be between 0 and 100")

	_, err = new.Percentage(101)
	assert.EqualError(t, err, "The percentage must be between 0 and 100")
}

func TestEquals(t *testing.T) {
	new := createUSD(100)
	new2 := createUSD(100)

	eq, _ := new.Equals(new2)
	assert.True(t, eq)

	_, err := new.Equals(create(100, "GBP"))
	assert.EqualError(t, err, "The currency does not match")
}

func TestGreaterThan(t *testing.T) {
	new := createUSD(100)
	new2 := createUSD(40)
	new3 := createUSD(200)

	gt, _ := new.GreaterThan(new2)
	assert.True(t, gt)

	gt, _ = new.GreaterThan(new3)
	assert.False(t, gt)

	_, err := new.GreaterThan(create(100, "GBP"))
	assert.EqualError(t, err, "The currency does not match")
}

func TestGreaterThanOrEqual(t *testing.T) {
	new := createUSD(100)
	new2 := createUSD(40)
	new3 := createUSD(200)

	gt, _ := new.GreaterThanOrEqual(new2)
	assert.True(t, gt)

	gt, _ = new.GreaterThanOrEqual(new3)
	assert.False(t, gt)

	gt, _ = new.GreaterThanOrEqual(createUSD(100))
	assert.True(t, gt)

	_, err := new.GreaterThanOrEqual(create(100, "GBP"))
	assert.EqualError(t, err, "The currency does not match")
}

func TestLessThan(t *testing.T) {
	new := createUSD(100)
	new2 := createUSD(40)
	new3 := createUSD(200)

	gt, _ := new.LessThan(new2)
	assert.False(t, gt)

	gt, _ = new.LessThan(new3)
	assert.True(t, gt)

	_, err := new.LessThan(create(100, "GBP"))
	assert.EqualError(t, err, "The currency does not match")
}

func TestLessThanOrEqual(t *testing.T) {
	new := createUSD(100)
	new2 := createUSD(40)
	new3 := createUSD(200)

	gt, _ := new.LessThanOrEqual(new2)
	assert.False(t, gt)

	gt, _ = new.LessThanOrEqual(new3)
	assert.True(t, gt)

	gt, _ = new.LessThanOrEqual(createUSD(100))
	assert.True(t, gt)

	_, err := new.LessThanOrEqual(create(100, "GBP"))
	assert.EqualError(t, err, "The currency does not match")
}

func TestAllocate(t *testing.T) {
	new := createUSD(1003)

	split, _ := new.Allocate(50, 50)
	assert.Equal(t, int64(502), split[0].Amount())
	assert.Equal(t, int64(501), split[1].Amount())

	new2 := createUSD(100)

	split2, _ := new2.Allocate(1, 3)
	assert.Equal(t, int64(25), split2[0].Amount())
	assert.Equal(t, int64(75), split2[1].Amount())
}

func create(amount int64, currency string) *Money {
	m, _ := New(amount, currency)
	return m
}

func createUSD(amount int64) *Money {
	m, _ := USD(amount)
	return m
}
