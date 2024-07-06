package main

import (
	"fmt"
)

func quick_check3() {
	var room = "cave"

	if (room == "cave") {
		fmt.Println("you are at a cave")
	} else if (room == "entrance") {
		fmt.Println("you are at entrance")
	} else if (room == "mountain") {
		fmt.Println("you are in a cave")
	} else {
		fmt.Println("idk where you are :p sorry")
	}
}