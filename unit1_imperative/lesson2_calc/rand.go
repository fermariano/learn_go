package main

import (
	"fmt"
	"math/rand"
)

func random() {
	var num = rand.Float32() + 1
	fmt.Println(num)

	var num2 = rand.Intn(10) + 1 // intn means int not signed
	fmt.Println(num2)
}
