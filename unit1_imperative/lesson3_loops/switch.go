package main

import (
	"fmt"
)

func main() {
	var command = "a"

	switch command {
	case "go inside":
		fmt.Println("bla bla bla")
	case "enter cave", "go east":
		fmt.Println("yesssssss")
	case "read sign":
		fmt.Println("reading...")
	default:
		fmt.Println("this is a default message :p")
	}

	var room = "underwater"

	switch room {
	case "cave":
		fmt.Println("cavern!")
	case "lake":
		fmt.Println("icyyy")
		fallthrough
	case "underwater":
		fmt.Println("so cold here")

	}
}
