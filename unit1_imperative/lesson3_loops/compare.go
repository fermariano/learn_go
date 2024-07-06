package main

import (
	"fmt"
	"strings"
)

func compare() {
	var age = 41
	var minor = age < 18

	var test = strings.EqualFold("apple", "aple")

	fmt.Println("at age ", age, "am I a minor? ", minor)
	fmt.Println(test)
}
