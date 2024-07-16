package main

import (
	"fmt"
)

func main() {
	const distance = 236000000000000000

	const (
		lightSpeed    = 299792
		secondsPerDay = 86400
	)

	fmt.Println(distance/lightSpeed/secondsPerDay/365, "light years")

}
