package main

import (
	"fmt"
	"strconv"
)

func numberString() {
	countdown := 10

	// using itoa

	str := "launch in t minus " + strconv.Itoa(countdown) + " seconds"
	fmt.Println(str)

	// using sprintf

	countdown = 9
	str = fmt.Sprintf("launch in T minus %v seconds", countdown)
	fmt.Println(str)

	// using atoi + error

	count, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("something went wrong")
	}

	fmt.Println(count)
}
