/*
Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis. Write a program
to determine how fast a ship would need to travel (in km/h) in order to reach Malacandra in 28 days. Assume a distance of 56,000,000 km.
*/

package main

import ("fmt")

func malacandra() {
	const (
		distance = 56000000.0
	 	days = 28
	 	hours = 28 * 24
	)

	var vlcty = distance / hours;

	fmt.Println(vlcty)
	 
}