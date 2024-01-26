package barber_shop

import (
	"fmt"
	"sleeping-barber-dilemma/customer"
	"sync"
)

type BarberShop struct {
	wg              sync.WaitGroup
	waitingRoom     chan customer.Customer
	numberOfBarbers int
	barbers         []Barber
	isShopClosed    bool
}

func NewBarberShop(numberOfSeats int, numberOfBarbers int) BarberShop {
	return BarberShop{
		wg:              sync.WaitGroup{},
		waitingRoom:     make(chan customer.Customer, numberOfSeats),
		numberOfBarbers: numberOfBarbers,
		barbers:         make([]Barber, numberOfBarbers),
	}
}

func (bs *BarberShop) Open() {
	for i := 1; i <= bs.numberOfBarbers; i++ {
		b := NewBarber(i, bs)
		bs.barbers[i-1] = b
		go b.Work()
		bs.wg.Add(1)
	}
}

func (bs *BarberShop) Close() {
	fmt.Println("Barbershop is closing now.")
	close(bs.waitingRoom)
	bs.isShopClosed = true
	fmt.Println("Barbershop is closed")
}

func (bs *BarberShop) IsShopClose() bool {
	return bs.isShopClosed
}

func (bs *BarberShop) WaitTillAllBarberReturnsHome() {
	bs.wg.Wait()
	fmt.Println("Barbershop is closed, haircut completed for all customers and all barbers have gone home.")
}

func (bs *BarberShop) CustomerVisit(customer customer.Customer) {
	fmt.Printf("Customer %d entered barber shop\n", customer.GetCustomerId())
	select {
	case bs.waitingRoom <- customer:
		bs.wg.Add(1)
		customer.WaitForHaircutToBeCompleted()
		bs.wg.Done()
	default:
		fmt.Printf("Customer %d left without a haircut as the waiting room is full\n", customer.GetCustomerId())
	}
}

func (bs *BarberShop) BarberReturnsToHome() {
	bs.wg.Done()
}

func (bs *BarberShop) GetWaitingRoom() chan customer.Customer {
	return bs.waitingRoom
}
