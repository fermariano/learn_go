package main

import (
	"fmt"
)

func constant() {

	const (
		distance = 24000000000000000000
		lightSpeed = 299792
		secondsPerDay = 86400
	)

	const days = distance / lightSpeed / secondsPerDay

	fmt.Println("andromeda galaxy is", days, "light days away")

}