package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

// kelvinToCelsius converts k to c
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func temperature_types() {
	var k kelvin = 294.0
	var c celsius
	var f fahrenheit = 84.0

	c = kelvinToCelsius(k)
	c = k.celsius()

	fmt.Println(c)

	c = f.celsius()

	fmt.Println(c)

}
