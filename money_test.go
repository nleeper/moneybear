package moneybear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	new, _ := New(100, "USD")
	assert.Equal(t, 100, new.amount)
	assert.Equal(t, "USD", new.currency)
}

func TestUSD(t *testing.T) {
	new, _ := USD(101)
	assert.Equal(t, 101, new.amount)
	assert.Equal(t, "USD", new.currency)
}
