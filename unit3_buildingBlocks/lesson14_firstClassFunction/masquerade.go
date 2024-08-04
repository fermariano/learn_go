package main

import "fmt"

var f = func() {
	fmt.Println("dress up for the masquerade")
}

func masquerade() {
	f()

	g := func(message string) {
		fmt.Println(message)
	}

	g("hellooo")

	func() {
		fmt.Println("function anonynous")
	}()
}
