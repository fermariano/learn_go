package main

import (
	"fmt"
	"math"
)

func ariane() {
	var bh float64 = 32767
	var h = int16(bh)

	fmt.Println(h)

	// wraps around

	bh = 32768
	h = int16(bh)

	fmt.Println("wrap around: ", h)

	// detect if will result an invalid value using math pack

	if bh < math.MinInt16 || bh > math.MaxInt16 {
		// handle out of range value
		bh--
		h = int16(bh)
	}

	fmt.Println("handled value: ", h)

}
