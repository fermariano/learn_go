package main

import (
	"fmt"
	"strings"
)

func branch_if() {
	var command = "go east"

	if command == "go east" {
		fmt.Println("you head further up the mountain")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("didn't quite get that")
	}

	// outro jeito usando o strings

	if strings.Compare(command, "go east") == 0 {
		fmt.Println("you head further up the mountain")
	} else if strings.Compare(command, "go inside") == 0 {
		fmt.Println("you enter the cave where you live out the rest of your life")
	} else {
		fmt.Println("didn't quite get that")
	}
}
