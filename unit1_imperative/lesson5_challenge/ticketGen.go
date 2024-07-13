/*
Start by building a prototype that generates 10 random tickets and displays them in a tabular format
with a nice header

The table should have four columns:
 - The spaceline company providing the service
 - The duration in days for the trip to Mars (one-way)
 - Whether the price covers a return trip
 - The price in millions of dollars 
For each ticket, randomly select one of the following spacelines: Space Adventures,
SpaceX, or Virgin Galactic.
Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km
away from Earth at the time.
Randomly choose the speed the ship will travel, from 16 to 30 km/s. This will determine
the duration for the trip to Mars and also the ticket price. Make faster ships more expensive, ranging in price from $36 million to $50 million. Double the price for round trips
*/

package main

import (
	"fmt"
	"math/rand"
)

func generateTicket() {
	fmt.Printf("%-20s %-5s %-10s %5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("=========================================")

	// generate spaceline name
	for i := 0; i < 10; i++ {
		spaceline := rand.Intn(3)
		var spacelineText = ""

		switch spaceline {
		case 0:
			spacelineText = "Virgin Galactic"
		case 1:
			spacelineText = "SpaceX"
		case 2:
			spacelineText = "Space Adventures"
		}

		const distance = 62100000
		velo := (rand.Intn(14) + 16) * 3600 // km/s -> km/h

		days := (distance / velo) / 24 // hours -> days

		price := days + 20

		// generate trip type
		tripT := rand.Intn(2)
		var tripText = ""

		switch tripT {
		case 0:
			tripText = "one-way"
		case 1:
			tripText = "round-trip"
			price *= 2
		}

		fmt.Printf("%-20s %-5d %-10s %5v\n", spacelineText, days, tripText, price)
	}

}

func main() {
	generateTicket()
}
