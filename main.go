package main

import (
	"sleeping-barber-dilemma/barber_shop"
	"sleeping-barber-dilemma/constants"
	"sleeping-barber-dilemma/customer"
	"time"
)

func main() {
	barberShop := barber_shop.NewBarberShop(constants.NumberOfSeats, constants.NumberOfBarbers)

	barberShop.Open()
	ingestCustomerIntoShop(barberShop)
	barberShop.WaitTillAllBarberReturnsHome()
}

func ingestCustomerIntoShop(barberShop barber_shop.BarberShop) {
	for i := 1; ; i++ {
		if barberShop.IsShopClose() {
			return
		}
		go barberShop.CustomerVisit(customer.NewCustomer(i))
		time.Sleep(constants.NewCustomerIngestionInterval)
	}
}
