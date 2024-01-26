package constants

import "time"

const (
	NumberOfBarbers              = 5
	NumberOfSeats                = 3
	ShopOperatingHours           = 10 * time.Second
	NewCustomerIngestionInterval = 250 * time.Millisecond
	HaircutDuration              = time.Duration(500) * time.Millisecond
)
