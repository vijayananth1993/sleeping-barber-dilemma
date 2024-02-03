package customer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCustomerId_ShouldReturnCustomerId(t *testing.T) {
	assert.Equal(t, 1, NewCustomer(1).GetCustomerId())
}

func TestHaircutCompleted_ShouldSendEventInHaircutCompletedCompletedChannel(t *testing.T) {
	c := customer{id: 1, haircutCompleted: make(chan bool, 1)}
	c.HaircutCompleted()
	assert.Equal(t, true, <-c.haircutCompleted)
}

func TestHaircutCompleted_ShouldReadEventInHaircutCompletedCompletedChannel(t *testing.T) {
	c := customer{id: 1, haircutCompleted: make(chan bool, 1)}
	c.haircutCompleted <- true
	c.WaitForHaircutToBeCompleted()
	assert.Equal(t, 0, len(c.haircutCompleted))
}
