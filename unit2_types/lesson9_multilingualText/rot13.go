package main

import (
	"fmt"
)

func rot13() {
	message := "uv vagreangvbany fcrpr fgngvba"

	for i := 0; i < len(message); i++ {
		c := message[i]

		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
}
