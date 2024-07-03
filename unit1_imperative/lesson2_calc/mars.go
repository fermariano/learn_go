// weight loss program
package main

import (
	"fmt"
)

// main is the func where it all begins

func main() {
	fmt.Print("My weight on the surface of Mars is: ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print("lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print("years old")

	fmt.Println("\nOR ")

	fmt.Print("My weight on the surface of Mars is: ", 149.0*0.3783, " lbs, and I would be ", 41*365/687, "years old")

	fmt.Println("\nOR")

	fmt.Printf("My weight on the surface of Mars is %v lbs, and I would be %v years old", 149.0*0.3783, 41*365/687)
}
