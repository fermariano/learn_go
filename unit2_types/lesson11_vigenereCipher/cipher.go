package main

import (
	"fmt"
	"strings"
)

func cipher() {
	text := "your message goes here"
	keyword := "GOLANG"
	i := 0
	message := ""

	text = strings.ToUpper(text)
	text = strings.Replace(text, " ", "", 100)

	for _, c := range text {
		c = c - 'A'

		if i == len(keyword) {
			i = 0
		}
		k := keyword[i] - 'A'
		//fmt.Println(c, keyword[i])
		conv := (int32(c)+int32(k)+26)%26 + 'A'
		message += string(conv)

		i++
	}

	fmt.Println(message)
}
