package customer

import "fmt"

type Customer struct {
	id               int
	haircutCompleted chan bool
}

func (c *Customer) WaitForHaircutToBeCompleted() {
	<-c.haircutCompleted
}

func (c *Customer) HaircutCompleted() {
	fmt.Printf("Customer %d is leaving with a haircut\n", c.id)
	c.haircutCompleted <- true
}

func (c *Customer) GetCustomerId() int {
	return c.id
}

func NewCustomer(id int) Customer {
	return Customer{id: id, haircutCompleted: make(chan bool)}
}
