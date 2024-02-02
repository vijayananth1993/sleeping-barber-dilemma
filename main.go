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
	barberShop.Close()
	barberShop.WaitTillAllBarberReturnsHome()
}

func ingestCustomerIntoShop(barberShop barber_shop.BarberShop) {
	startTime := time.Now()
	for i := 1; ; i++ {
		if isTimeForShopClosure(startTime) {
			return
		}
		go barberShop.CustomerVisit(customer.NewCustomer(i))
		time.Sleep(constants.NewCustomerIngestionInterval)
	}
}

func isTimeForShopClosure(startTime time.Time) bool {
	return time.Now().Sub(startTime) > constants.ShopOperatingHours
}
