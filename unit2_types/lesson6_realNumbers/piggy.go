package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var total float32

	for total < 200 {
		var random = rand.Intn(2)

		if random == 0 {
			total += 5
		} else if random == 1 {
			total += 10
		} else if random == 2 {
			total += 25
		}

		fmt.Println(total / 10)
	}

}
