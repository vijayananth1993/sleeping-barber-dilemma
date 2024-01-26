package barber_shop

import (
	"fmt"
	"sleeping-barber-dilemma/constants"
	"sleeping-barber-dilemma/customer"
	"time"
)

type Barber struct {
	id         int
	barberShop *BarberShop
}

func (b *Barber) Work() {
	defer b.returnHome()
	for {
		if b.barberShop.IsShopClose() {
			b.haircutAllWaitingCustomers()
			return
		}
		b.haircutOrGoToSleep()
	}
}

func (b *Barber) haircutOrGoToSleep() {
	select {
	case c, ok := <-b.barberShop.GetWaitingRoom():
		if ok {
			b.haircut(c)
		}
	default:
		b.sleep()
	}
}

func (b *Barber) returnHome() {
	fmt.Printf("Barber %d returning home\n", b.id)
	b.barberShop.BarberReturnsToHome()
}

func (b *Barber) haircutAllWaitingCustomers() {
	fmt.Printf("Barber %d, processing waiting customers\n", b.id)
	for c := range b.barberShop.GetWaitingRoom() {
		b.haircut(c)
	}
	fmt.Printf("Barber %d, processed all waiting customers\n", b.id)
}

func (b *Barber) sleep() {
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

func (b *Barber) haircut(customer customer.Customer) {
	fmt.Printf("Barber %d is cutting hair for customer %d\n", b.id, customer.GetCustomerId())
	time.Sleep(constants.HaircutDuration)
	fmt.Printf("Barber %d finished with customer %d\n", b.id, customer.GetCustomerId())
	customer.HaircutCompleted()
}

func NewBarber(id int, barberShop *BarberShop) Barber {
	return Barber{id: id, barberShop: barberShop}
}
