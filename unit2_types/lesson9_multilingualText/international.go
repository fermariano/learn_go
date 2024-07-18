package main

import (
	"fmt"
	//"unicode/utf8"
)

func main() {
	message := "Hola Estación Espacial Internacional"

	// c, size := utf8.DecodeRuneInString(message)

	for _, c := range message {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			c += 13

			if c > 'z' || c > 'Z' && c < 'a' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)

	}
}
