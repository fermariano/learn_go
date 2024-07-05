/*
The distance between Earth and Mars varies from nearby to opposite sides
of the sun. Write a program that generates a random distance from 56,000,000 to
401,000,000 km
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = (rand.Intn(345) + 56) * 1000000
	// 345 because its 401 - 56 para nao ultrapassar o valor -> a soma é para colocar a partir daquele valor
	// multipliquei por 100k no fim pra nao ficar colocando numero grande :p

	fmt.Println(number)
}
