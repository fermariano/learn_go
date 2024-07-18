package main

import (
	"fmt"
)

func ceasar() {
	message := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(message); i++ {
		c := message[i]

		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			c -= 3

			if c < 'a' || c > 'z' && c > 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
}
