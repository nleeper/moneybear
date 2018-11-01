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

	new, _ := USD(expected)
	assert.Equal(t, expected, new.Amount())

	new.Add(100).Add(30)

	assert.Equal(t, expected+130, new.Amount())
}
