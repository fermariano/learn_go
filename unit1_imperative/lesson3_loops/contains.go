package main

import (
	"fmt"
	"strings"
)

func contains() {
	fmt.Println("You find yourself in a dimly lit cavern")
	var command = "walk outside"
	var test = "outside"
	var exit = strings.Contains(command, test)

	fmt.Println("You leave the cave: ", exit)
}
