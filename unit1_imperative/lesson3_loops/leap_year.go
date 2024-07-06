package main

import (
	"fmt"
)

func leap_year() {
	var year = 2100

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("leap year")
	} else {
		fmt.Println("not a leap year")
	}
}
