package main

import (
	"fmt"

	"github.com/FCNN/math"
	"github.com/FCNN/neuron"
)

func main() {
	input := math.NewVector(2, []float64{
		0.0, 0.0,
	})

	target := math.NewVector(1, []float64{
		0.3,
	})

	output := math.CreateVector(2)

	nn := neuron.NewNN(2, 1, 1)
	nn.SetAlpha(0.1)

	for i := 0; i < 100; i++ {
		fmt.Println("loop [", i, "] :")

		nn.SetInput(input)

		nn.PropForward()

		nn.CopyOutput(output, false)
		output.Print()

		nn.PropBackward(target)

		fmt.Print("\n")
	}
}
