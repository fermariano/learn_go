package main

import (
	"fmt"
	"math/rand"
)

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func sensorStudy() {

	var sensorDeclared func() kelvin

	sensor := fakeSensor // assigns a function to a variable
	fmt.Println(sensor())

	sensorDeclared = realSensor
	fmt.Println(sensorDeclared())
}
