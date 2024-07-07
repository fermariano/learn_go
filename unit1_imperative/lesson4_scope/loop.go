package main

import (
	"fmt"
)

func loop() {
	var count = 0

	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}

	fmt.Println(count) // count remains in scope

	// another way

	for count2 := 10; count2 > 0; count2-- {
		fmt.Println(count2)
	}

	// fmt.Println(count2) <-- error: not declared because there's not on scope anymore
}
