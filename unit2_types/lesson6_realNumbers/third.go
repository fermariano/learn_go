package main

import "fmt"

func third() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)

	// formats the value of third with a width and with precision

	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)

}
