package barber_shop

//go:generate mockgen -source=barber.go -destination=./mocks/barber.go

import (
	"fmt"
	"sleeping-barber-dilemma/constants"
	"sleeping-barber-dilemma/customer"
	"time"
)

type Barber interface {
	Work()
}

type barber struct {
	id         int
	barberShop BarberShop
}

func (b *barber) Work() {
	defer b.returnHome()
	for {
		if b.barberShop.IsShopClose() {
			b.haircutAllWaitingCustomers()
			return
		}
		b.haircutOrGoToSleep()
	}
}

func (b *barber) haircutOrGoToSleep() {
	select {
	case c, ok := <-b.barberShop.GetWaitingRoom():
		if ok {
			b.haircut(c)
		}
	default:
		b.sleep()
	}
}

func (b *barber) returnHome() {
	fmt.Printf("Barber %d returning home\n", b.id)
	b.barberShop.BarberReturnsToHome()
}

func (b *barber) haircutAllWaitingCustomers() {
	fmt.Printf("Barber %d, processing waiting customers\n", b.id)
	for c := range b.barberShop.GetWaitingRoom() {
		b.haircut(c)
	}
	fmt.Printf("Barber %d, processed all waiting customers\n", b.id)
}

func (b *barber) sleep() {
	fmt.Printf("Barber %d is seelping untill customer awakes\n", b.id)
	for {
		if b.barberShop.IsShopClose() {
			return
		}
		select {
		case customer, ok := <-b.barberShop.GetWaitingRoom():
			if ok {
				fmt.Printf("Customer %d wakes up Barber %d\n", customer.GetCustomerId(), b.id)
				b.haircut(customer)
			}
			return
		}
	}
}

func (b *barber) haircut(customer customer.Customer) {
	fmt.Printf("Barber %d is cutting hair for customer %d\n", b.id, customer.GetCustomerId())
	time.Sleep(constants.HaircutDuration)
	fmt.Printf("Barber %d finished with customer %d\n", b.id, customer.GetCustomerId())
	customer.HaircutCompleted()
}

func NewBarber(id int, barberShop BarberShop) Barber {
	return &barber{id: id, barberShop: barberShop}
}
