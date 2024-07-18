package main

import (
	"fmt"
	"unicode/utf8"
)

func spanish() {
	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes")                    // 15 bytes
	fmt.Println(utf8.RuneCountInString(question), "runes") // 12 runes (characters)

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes\n", c, size)

	// range

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}

	for _, c := range question {
		fmt.Printf("%c\n", c)
	}
}
