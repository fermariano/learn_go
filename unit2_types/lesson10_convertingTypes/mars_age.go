package main

import (
	"fmt"
	"strconv"
)

func mars_age() {

	number := "28"
	numberConv, _ := strconv.Atoi(number)

	fmt.Println(numberConv)

	age := 41
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays

	fmt.Println("i am", marsAge, "years old on mars")
}
