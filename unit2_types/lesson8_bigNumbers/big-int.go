package main

import (
	"fmt"
	"math/big"
)

func big_int() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	// newInt isnt going to help for a number like 24 quintillion - it wont fit in an int64 so instead you can create a big.int from a string

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10) // 24 quintillion is in base 10

	fmt.Println(lightSpeed, secondsPerDay)
	fmt.Println(distance)

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("that is", days, "days of travel at light speed")
}
