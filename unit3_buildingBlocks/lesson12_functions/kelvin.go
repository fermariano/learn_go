package main

import (
	"fmt"
)

// converts k to c
func kelvinToCelsius(k float64) float64 { // receive float64 return float64
	k -= 273.15
	return k
}

// converts c to f
func celsiusToFahrenheit(c float64) float64 {
	c = (c * 9.0 / 5.0) + 32.0
	return c
}

// converts k to f
func kelvintoFahrenheit(k float64) float64 {
	k = kelvinToCelsius(k)
	return celsiusToFahrenheit(k)
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "is ", celsius, "celsius")

	celsius = kelvinToCelsius(233.0)
	fmt.Println(celsius)

	fmt.Println(kelvintoFahrenheit(0.0))

}
