package main

import (
	"fmt"
)

func launch() {
	launch := false

	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("ready for launch: " + launchText)

	var yesNo string

	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}

	fmt.Println("ready for launch: ", yesNo)

	// inverse conversion

	yesNo = "no"

	launch = (yesNo == "yes")
	fmt.Println("ready for launch", launch)

}
