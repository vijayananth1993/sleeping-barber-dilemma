package customer

//go:generate mockgen -source=customer.go -destination=./mocks/customer.go

import "fmt"

type Customer interface {
	WaitForHaircutToBeCompleted()
	HaircutCompleted()
	GetCustomerId() int
}

type customer struct {
	id               int
	haircutCompleted chan bool
}

func (c *customer) WaitForHaircutToBeCompleted() {
	<-c.haircutCompleted
}

func (c *customer) HaircutCompleted() {
	fmt.Printf("Customer %d is leaving with a haircut\n", c.id)
	c.haircutCompleted <- true
}

func (c *customer) GetCustomerId() int {
	return c.id
}

func NewCustomer(id int) Customer {
	return &customer{id: id, haircutCompleted: make(chan bool, 1)}
}
