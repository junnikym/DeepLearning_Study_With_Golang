package neuron

import (
	"fmt"
	"math/rand"

	"github.com/FCNN/math"
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

func NewNN(n_input, n_output, n_hiddenLayer int) *NeuralNetwork {
	var result NeuralNetwork

	result.n_act = make([]int, n_hiddenLayer+2) // + 2 is input and output layers

	result.n_act[0] = n_input + 1 // [0] is  input layer; + 1 is include bias

	for i := 1; i < n_hiddenLayer+1; i++ {
		result.n_act[i] = n_input + 1 // default value
	}

	result.n_act[n_hiddenLayer+1] = n_output + 1

	result.Init(result.n_act, n_hiddenLayer)

	return &result
}

func (n *NeuralNetwork) Init(n_act []int, n_hiddenLayer int) {
	n.n_input = n_act[0] - 1                // act[0]		: input
	n.n_output = n_act[n_hiddenLayer+1] - 1 // -1 : except bias

	n.n_layer = n_hiddenLayer + 2 // + input, output (2 layer)

	n.bias = 1
	n.alpha = 0.15

	// itialize NN->act;	activation reesult
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	n.act = make([]math.Vector, n.n_layer)

	for i := 0; i < n.n_layer; i++ {
		n.act[i].Init(n.n_act[i])                // allocate act memory
		n.act[i].ReplaceAt(n.n_act[i]-1, n.bias) // set : last one is bias
	}

	// initialize NN->gardients
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	n.gradient = make([]math.Vector, n.n_layer)

	for i := 0; i < n.n_layer; i++ {
		n.gradient[i].Init(n.n_act[i])
	}

	// initialize NN->weight
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	n.weight = make([]math.Matrix, n.n_layer-1) // -1 : weight is between layers

	for i := 0; i < len(n.weight); i++ {
		n.weight[i].Init(
			n.act[i+1].Size()-1, // Next layer ( -1 : bias )
			n.act[i].Size(),     // Prev layer
		) // Row x Column = Next x Prev(include bias)

		// assign random value to weight
		for r := 0; r < n.weight[i].Row(); r++ {
			for c := 0; c < n.weight[i].Col(); c++ {
				n.weight[i].ReplaceAt(r, c, rand.Float64()*0.1)
			}
		}
	}
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Neuron		| getter
//----- ----- ----- ----- ----- ----- ----- -----*/

func (n *NeuralNetwork) CopyOutput(copy *math.Vector, is_copyBias bool) {

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

func (n *NeuralNetwork) DebugPrint() {
	fmt.Println("Neural Network : ")
	fmt.Println("\t[ bias : ", n.bias, "], [ alpha : ", n.alpha, "]")

	// number of each layer
	fmt.Println("number of each layer : ")
	fmt.Printf(
		"\t[ input : %d ], [ output : %d ], [ n_layer : %d ] \n",
		n.n_input, n.n_output, n.n_act,
	)

	// activation
	fmt.Println("act : ")
	fmt.Println("\t", n.act)

	// gradient
	fmt.Println("gradient : ")
	fmt.Println("\t", n.gradient)
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Neuron		| setter
//----- ----- ----- ----- ----- ----- ----- -----*/

func (n *NeuralNetwork) SetInput(input *math.Vector) {
	if input.Size() < n.n_input {
		panic("input size is different with NN->input size")
	}

	for i := 0; i < n.n_input; i++ {
		n.act[0].ReplaceAt(i, input.At(i)) // n.act[0] is Input Layer
	}
}

func (n *NeuralNetwork) SetAlpha(x float64) {
	n.alpha = x
}

func updateWeight(weight *math.Matrix, nextGradient, prevAct *math.Vector, alpha float64) {

	var delta_weight float64 = 0

	for r := 0; r < weight.Row(); r++ {
		for c := 0; c < weight.Col(); c++ {
			delta_weight = alpha * nextGradient.At(r) * prevAct.At(c)

			weight.ReplaceAt(r, c, weight.At(r, c)+delta_weight)
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

	/* =========================
	 * Calculate Gradients
	 ========================= */

	// Calculate output layer Gradients
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	layer = len(n.gradient) - 1

	for i := 0; i < n.gradient[layer].Size()-1; i++ {
		output = n.act[layer].At(i)

		n.gradient[layer].ReplaceAt(
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
				(n.gradient[layer].At(i) * ReLUGradient(n.act[layer].At(i))),
			)
		}
	}

	// Update weights
	//----- ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
	for layer = len(n.weight) - 1; layer >= 0; layer-- {
		updateWeight(&n.weight[layer], &n.gradient[layer+1], &n.act[layer], n.alpha)
	}
}
