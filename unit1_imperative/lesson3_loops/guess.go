package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = 50
	var nRand = rand.Intn(100) + 1

	for number != nRand {
		if nRand > number {
			fmt.Println(nRand, ": too big")
		} else {
			fmt.Println(nRand, ": too small")
		}
		nRand = rand.Intn(99) + 1
	}
	fmt.Println("correct number is ", nRand)
}
