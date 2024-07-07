package main

import (
	"fmt"
	"math/rand"
)

func short_if() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("space adventures")
	} else if num == 1 {
		fmt.Println("spacex")
	} else {
		fmt.Println("galactic")
	} // num is no longer in scope

	switch num := rand.Intn(3); num {
	case 1:
		fmt.Println("yesss")
	case 2:
		fmt.Println("bla bla bla")
	default:
		fmt.Println("default message :p")
	} // num is no longer in scope too
}
