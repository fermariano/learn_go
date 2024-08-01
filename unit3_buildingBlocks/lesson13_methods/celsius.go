package main

import (
	"fmt"
)

func celsius_study() {
	type celsius float64

	const degrees = 20.4

	var temperature celsius = degrees
	temperature += 10
	fmt.Println(temperature)

	// error - float64 with celsius

	var warmUp float64 = 10
	temperature += celsius(warmUp) // invalid operation: mismatched types

	// types cant mix

	type fahrenheit float64

	var c celsius = 20
	var f fahrenheit = 20

	if c == f { // invalid operation: mismatched types celsius and fahrenheit
	}

	c += f
}
