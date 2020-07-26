package neuron

import "math"

// Activation functions
//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----

func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func SigmoidGradient(y float64) float64 {
	return (1.0 - y) * y
}

func ReLU(x float64) float64 {
	if x > 0.0 {
		return x
	} else {
		return 0.0
	}
}

func ReLUGradient(y float64) float64 {
	if x > 0.0 {
		return 1.0
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

func LReLUGradient(y float64) float64 {
	if y > 0.0 {
		return 1.0
	} else {
		return 0.01
	}
}
