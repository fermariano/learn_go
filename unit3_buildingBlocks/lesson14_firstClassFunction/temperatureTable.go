package main

import "fmt"

type ce float64
type fa float64

type data func(ce, fa) (fa, ce)

func getData(celsius ce, fahrenheit fa) (fa, ce) {
	return celsius.fahrenheit(), fahrenheit.celsius()
}

func (celsius ce) fahrenheit() fa {
	return fa((celsius * 9 / 5) + 32)
}

func (fah fa) celsius() ce {
	return ce((fah - 32) * 5 / 9)
}

func drawTableC(getData data, c ce, f fa) {
	fah, _ := getData(c, f)
	fmt.Printf("| %6.1f | %6.1f |\n", c, fah)
}

func drawTableF(getData data, c ce, f fa) {
	_, cel := getData(c, f)
	fmt.Printf("| %6.1f | %6.1f |\n", f, cel)
}

func main() {

	// celsius table
	fmt.Printf("====================\n")
	fmt.Printf("| %6s | %6s |\n", "ºC", "ºF")
	var i ce

	for i = -40; i <= 100; i += 5 {
		drawTableC(getData, i, 0)
	}

	// fahrenheit table
	fmt.Printf("====================\n")
	fmt.Printf("| %6s | %6s |\n", "ºF", "ºC")
	var j fa

	for j = -40; j <= 100; j += 5 {
		drawTableF(getData, 0, j)
	}

}
