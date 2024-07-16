package main

import (
	"fmt"
)


func alpha() {

	const (
		lightSpeed = 299792 // km/h
		secondsPerDay = 86400
	)
	
	var distance int64 = 41.3e12

	fmt.Println("Alpha Centauri is", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("that is", days, "days of travel at light speed")
}