package main

import (
	"fmt"
	"math/rand"
)

func infinity() {
	var degrees = 0

	for {
		fmt.Println(degrees)

		degrees++

		if degrees >= 360 {
			degrees = 0
			var test = rand.Intn(2)
			if test == 0 {
				fmt.Println("rand", test)
				break
			}
		}
	}
}
