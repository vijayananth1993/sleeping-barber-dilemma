package main

import (
	"sleeping-barber-dilemma/barber_shop"
	"sleeping-barber-dilemma/constants"
	"sleeping-barber-dilemma/customer"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	barberShop := barber_shop.NewBarberShop(constants.NumberOfSeats, constants.NumberOfBarbers)

	barberShop.Open()
	wg.Add(1)
	go func() {
		for i := 1; ; i++ {
			if isTimeForShopClosure(startTime) {
				barberShop.Close()
				wg.Done()
				return
			}
			go barberShop.CustomerVisit(customer.NewCustomer(i))
			time.Sleep(constants.NewCustomerIngestionInterval)
		}
	}()

	wg.Wait()
	barberShop.WaitTillAllBarberReturnsHome()
}

func isTimeForShopClosure(startTime time.Time) bool {
	return time.Now().Sub(startTime) > constants.ShopOperatingHours
}
