package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" // era is available throughout the package -> global

func main() {

	for count := 0; count < 10; count++ {
		year := rand.Intn(2100) + 1 // era and year are in scope
		month := rand.Intn(12) + 1
		daysInMonth := 31

		switch month { // era, year and month are in scope
		case 2:
			if year%4 == 0 && year%100 != 0 || year%400 == 0 {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		fmt.Println(era, year, month, rand.Intn(daysInMonth)+1)
	}

}
