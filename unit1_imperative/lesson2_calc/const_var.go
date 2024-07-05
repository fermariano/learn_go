// how long does it take to get to mars?

package main

import (
	"fmt"
)

func var_go() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000

	fmt.Println(distance/lightSpeed, "seconds")

	// declare multiple variables at once
	var (
		distance2  = 56000000
		speed2     = 100800
		hourPerDay = 24
	)

	var a, b, c = 1, 2, 3

	fmt.Println(distance2 / speed2 / hourPerDay)
	fmt.Println(a, b, c)

	// increment and assignments

	var weight = 149.0
	weight = weight * 0.3783
	weight *= 0.3783

	var age = 41
	age = age + 1
	age += 1
	age++
}
