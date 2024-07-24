package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	i := 0
	message := ""

	for _, c := range cipherText {
		c = c - 'A'

		if i == len(keyword) {
			i = 0
		}
		k := keyword[i] - 'A'
		//fmt.Println(c, keyword[i])
		conv := (int32(c)-int32(k)+26)%26 + 'A'
		message += string(conv)

		i++
	}

	fmt.Println(message)
}
