package barber_shop

import (
	"fmt"
	"sleeping-barber-dilemma/customer"
	"sync/atomic"
)

type BarberShop struct {
	waitingRoom         chan customer.Customer
	numberOfBarbers     int
	barbers             []Barber
	isShopClosed        int32
	barberReturningHome chan bool
}

func NewBarberShop(numberOfSeats int, numberOfBarbers int) BarberShop {
	return BarberShop{
		waitingRoom:         make(chan customer.Customer, numberOfSeats),
		numberOfBarbers:     numberOfBarbers,
		barbers:             make([]Barber, numberOfBarbers),
		barberReturningHome: make(chan bool, numberOfBarbers),
	}
}

func (bs *BarberShop) Open() {
	for i := 1; i <= bs.numberOfBarbers; i++ {
		b := NewBarber(i, bs)
		bs.barbers[i-1] = b
		go b.Work()
	}
}

func (bs *BarberShop) Close() {
	fmt.Println("Barbershop is closing now.")
	close(bs.waitingRoom)
	atomic.AddInt32(&bs.isShopClosed, int32(1))
	fmt.Println("Barbershop is closed")
}

func (bs *BarberShop) IsShopClose() bool {
	return atomic.LoadInt32(&bs.isShopClosed) != 0
}

func (bs *BarberShop) WaitTillAllBarberReturnsHome() {
	for i := 0; i < bs.numberOfBarbers; i++ {
		<-bs.barberReturningHome
	}
	fmt.Println("Barbershop is closed, haircut completed for all customers and all barbers have gone home.")
}

func (bs *BarberShop) CustomerVisit(customer customer.Customer) {
	fmt.Printf("Customer %d entered barber shop\n", customer.GetCustomerId())
	select {
	case bs.waitingRoom <- customer:
		customer.WaitForHaircutToBeCompleted()
	default:
		fmt.Printf("Customer %d left without a haircut as the waiting room is full\n", customer.GetCustomerId())
	}
}

func (bs *BarberShop) BarberReturnsToHome() {
	bs.barberReturningHome <- true
}

func (bs *BarberShop) GetWaitingRoom() chan customer.Customer {
	return bs.waitingRoom
}
