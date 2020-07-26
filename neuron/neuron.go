package neuron

import (
	"math/rand"

	"github.com/FNCC/math"
)

type NeuralNetwork struct {
	bias  float64
	alpha float64

	n_input  int
	n_output int
	n_layer  int

	act      []math.Vector // activation value ( result of single neuron ): act [ layer ] [ Neuron ]
	gradient []math.Vector
	weight   []math.Matrix

	n_act []int
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Neuron		| constructor
//----- ----- ----- ----- ----- ----- ----- -----*/

func NewNeuralNetwork(n_input, n_output, n_hiddenLayer int) *NeuralNetwork {
	var result NeuralNetwork

	result.n_act = make([]int, n_hiddenLayer+2) // + 2 is input and output layers

	result.n_act[0] = n_input + 1 // [0] is  input layer; + 1 is include bias

	for i := 1; i < n_hiddenLayer+1; i++ {
		result.n_act[i] = n_input + 1 // default value
	}

	result.n_act[n_hiddenLayer+1] = n_output + 1

	result.n_input = n_input
	result.n_output = n_output
	result.n_layer = n_hiddenLayer + 2

	result.bias = 1
	result.alpha = 0.15

	// initialize NN->act ;     activation reesult
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	result.act = make([]math.Vector, result.n_layer)

	for i := 0; i < result.n_layer; i++ {
		result.act[i].Init(result.n_act[i])                     // allocate act memory
		result.act[i].ReplaceAt(result.n_act[i]-1, result.bias) // set : last one is bias
	}

	// initialize NN->gardients
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	result.gradient = make([]math.Vector, result.n_layer)
	for i := 0; i < result.n_layer; i++ {
		result.act[i].Init(result.n_act[i])
	}

	// initialize NN->weight
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	result.weight = make([]math.Matrix, result.n_layer-1) // -1 : weight is between layers
	for i := 0; i < len(result.weight); i++ {
		result.weight[i].Init(
			result.act[i+1].Size()-1, // Next layer ( -1 : bias )
			result.act[i].Size(),     // Prev layer
		) // Row x Column = Next x Prev

		// assign random value to weight
		for r := 0; r < result.weight[i].Row(); r++ {
			for c := 0; c < result.weight[i].Col(); c++ {
				result.weight[i].ReplaceAt(r, c, rand.Float64()*0.1)
			}
		}
	}

	return &result
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Neuron		| getter
//----- ----- ----- ----- ----- ----- ----- -----*/

func (n *NeuralNetwork) CopyVector(copy *math.Vector, is_copyBias bool) {

	var outputLayerAct *math.Vector = &n.act[len(n.act)-1]
	var copySize int = n.n_output

	if is_copyBias == true {
		copySize += 1
	}

	if copy.Size() != copySize {
		copy.Init(copySize)
	}

	for i := 0; i < copySize; i++ {
		copy.ReplaceAt(i, outputLayerAct.At(i))
	}
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Neuron		| setter
//----- ----- ----- ----- ----- ----- ----- -----*/

func (n *NeuralNetwork) InputVec(input *math.Vector) {
	if input.Size() < n.n_input {
		panic("input size is different with NN->input size")
	}

	for i := 0; i < n.n_input; i++ {
		n.act[0].ReplaceAt(i, input.At(i)) // n.act[0] is Input Layer
	}
}

func updateWeight(weight *math.Matrix, nextGradient, prevAct *math.Vector, alpha float64) {

	var delta_weight float64 = 0

	for r := 0; r < weight.Row(); r++ {
		for c := 0; c < weight.Col(); c++ {

			delta_weight = alpha * nextGradient.At(r) * prevAct.At(c)

			weight.ReplaceAt(r, c, weight.At(r, c)*delta_weight)
		}
	}
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Neuron		| propagation
//----- ----- ----- ----- ----- ----- ----- -----*/

func (n *NeuralNetwork) PropForward() {
	for layer := 0; layer < len(n.weight); layer++ {
		n.weight[layer].Mul(&n.act[layer], &n.act[layer+1])

		ReLUToVec(&n.act[layer+1])
	}
}

func (n *NeuralNetwork) PropBackward(target *math.Vector) {
	var layer int = 0
	var output float64 = 0.0

	// Calculate output layer Gradients
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	layer = len(n.gradient) - 1

	for i := 0; i < n.gradient[layer].Size(); i++ {
		output = n.act[layer].At(i)

		n.gradient[i].ReplaceAt(
			i,
			(target.At(i)-output)*ReLUGradient(output),
		)
	}

	// Calculate hidden layer Gradients
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	for layer = len(n.weight) - 1; layer >= 0; layer-- {
		n.weight[layer].MulTrans(&n.gradient[layer+1], &n.gradient[layer])

		for i := 0; i < n.act[layer].Size()-1; i++ {
			n.gradient[layer].ReplaceAt(
				i,
				ReLUGradient(n.act[layer].At(i)),
			)
		}
	}

	// Update weights
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	for layer = len(n.weight) - 1; layer >= 0; layer-- {
		updateWeight(&n.weight[layer], &n.gradient[layer+1], &n.act[layer], n.alpha)
	}
}
