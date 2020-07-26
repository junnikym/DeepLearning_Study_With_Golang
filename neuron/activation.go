package neuron

import (
	"math"

	mat "github.com/FNCC/math"
)

// Activation functions
//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----

func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func ReLU(x float64) float64 {
	if x > 0.0 {
		return x
	} else {
		return 0.0
	}
}

func LReLU(x float64) float64 {
	if x > 0.0 {
		return 1.0
	} else {
		return 0.01
	}
}

// Gradient function

func SigmoidGradient(y float64) float64 {
	return (1.0 - y) * y
}

func ReLUGradient(y float64) float64 {
	if y > 0.0 {
		return 1.0
	} else {
		return 0.0
	}
}

func LReLUGradient(y float64) float64 {
	if y > 0.0 {
		return 1.0
	} else {
		return 0.01
	}
}

// apply activation function to Vector

func SigmoidToVec(v *mat.Vector) {
	for i := 0; i < v.Size()-1; i++ { // -1 : for not include bias
		v.ReplaceAt(i, Sigmoid(v.At(i)))
	}
}

func ReLUToVec(v *mat.Vector) {
	for i := 0; i < v.Size()-1; i++ { // -1 : for not include bias
		v.ReplaceAt(i, ReLU(v.At(i)))
	}
}

func LReLUToVec(v *mat.Vector) {
	for i := 0; i < v.Size()-1; i++ { // -1 : for not include bias
		v.ReplaceAt(i, LReLU(v.At(i)))
	}
}
