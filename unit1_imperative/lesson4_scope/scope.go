package main

import (
	"fmt"
	"math/rand"
)

func scope() {
	var count = 0

	for count < 10 { // a new scope begins
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	} // this scope ends
}
