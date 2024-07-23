package main

/*
Experiment: input.go
Write a program that converts strings to Booleans:
 The strings “true”, “yes”, or “1” are true.
 The strings “false”, “no”, or “0” are false.
 Display an error message for any other values
*/

import (
	"fmt"
)

func main() {
	stringTest := "5"

	var result bool

	if stringTest == "yes" || stringTest == "true" || stringTest == "1" {
		result = true
		fmt.Println(result)
	} else if stringTest == "no" || stringTest == "false" || stringTest == "0" {
		fmt.Println(result)
	} else {
		fmt.Println("error: invalid value")
	}

}
