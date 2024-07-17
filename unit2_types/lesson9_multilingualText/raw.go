package main

import (
	"fmt"
)

func raw() {
	fmt.Println("peace be upon\nupon you be peace") // literal string
	fmt.Println(`strings can span multiple lines with the \n escape sequence`) // raw string

	fmt.Println(`
		peace be upon you
		upon you be peace`)

	// both types are string

	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)
}
