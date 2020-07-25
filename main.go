package main

import (
	"fmt"
	"learning/deep_learning_go/math"
)

func main() {
	mat_a := math.NewMatrix(4, 4, []float64{
		11, 12, 13, 14,
		21, 22, 23, 24,
		31, 32, 33, 34,
		41, 42, 43, 44,
	})

	vec_a := math.NewVector(4, []float64{
		1, 2, 3, 4,
	})

	fmt.Println(mat_a)
	fmt.Println(vec_a)

	result := math.CreateVector(4)
	mat_a.MulTrans(vec_a, result)

	result.Print()
}
