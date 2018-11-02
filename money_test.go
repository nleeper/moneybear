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
	var expected int64 = 300

	new := createUSD(expected)
	assert.Equal(t, expected, new.Amount())

	new2 := new.Divide(100)

	assert.Equal(t, expected/100, new2.Amount())
	assert.Equal(t, new.currency, new2.currency)
}

func TestEquals(t *testing.T) {
	new := createUSD(100)
	new2 := createUSD(100)

	eq, _ := new.Equals(new2)
	assert.True(t, eq)

	_, err := new.Equals(create(100, "GBP"))
	assert.EqualError(t, err, "The currency does not match")
}

func create(amount int64, currency string) *Money {
	m, _ := New(amount, currency)
	return m
}

func createUSD(amount int64) *Money {
	m, _ := USD(amount)
	return m
}
