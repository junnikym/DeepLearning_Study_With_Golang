package main

import (
	"learning/QL/QL"
	"math/rand"
)

func main() {
	qGrid := QL.CreateGridWorld(5, 5)
	var moveDirection int = 0

	for loop := 0; loop < 10000; loop++ {
		moveDirection = rand.Intn(4)
	}

}
