package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var total float32

	for total < 2000 {
		var random = rand.Intn(2)

		if random == 0 {
			total += 50
		} else if random == 1 {
			total += 100
		} else if random == 2 {
			total += 250
		}

		fmt.Println(total / 100)
	}

}
