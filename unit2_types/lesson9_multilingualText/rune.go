package main

import (
	"fmt"
)

func rune_test() {
	var (
		pi    rune = 960
		alpha rune = 940
		omega rune = 969
	)

	var bang byte = 33

	var teste = "shalom"

	// var smiley rune = 128515

	// fmt.Printf("%c\n", smiley)

	fmt.Println(pi, alpha, omega, bang) // prints its numeric number

	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang) // prints πάω!

	fmt.Printf("%c\n", teste[5]) // prints m

	fmt.Println(len(teste)) // length

	// lowercase to uppercase
	c := 'y'
	c = c - 'a' + 'A'
	fmt.Printf("%c", c)


}
