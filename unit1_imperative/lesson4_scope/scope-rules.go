package main

import (
	"fmt"
	"math/rand"
)

var eraTest = "AD" // era is available throughout the package -> global

func scope_rules() {
	year := 2018 // era and year are in scope
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month { // era, year and month are in scope
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	fmt.Println(eraTest, year, month, rand.Intn(daysInMonth)+1)

}
