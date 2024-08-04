package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Println(k, "k")
		time.Sleep(time.Second)
	}
}

func measureTemperatureType(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Println(k, "k")
		time.Sleep(time.Second)
	}
}

func fakeSensor2() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func function_parameter() {
	measureTemperature(3, fakeSensor2)
	measureTemperatureType(3, fakeSensor2)
}
