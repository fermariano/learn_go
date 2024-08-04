package main

import "fmt"

type k float64

// sensor function type
type s func() k

func realS() k {
	return 0
}

func calibrate(sensor s, offset k) s { // recebe uma funcao do tipo s e um offset em k
	return func() k { // cria a func anonima
		return sensor() + offset // retorna a funcao + k
	}
}

func calibrateMain() {
	var offset k = 5
	sensor := calibrate(realS, offset)
	fmt.Println(sensor())

}
