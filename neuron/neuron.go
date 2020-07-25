package neuron

type Neuron struct {
	weight float64
	bias   float64

	input  float64
	output float64

	n_layer int
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Neuron		| constructor
//----- ----- ----- ----- ----- ----- ----- -----*/

func NewNeuron(w float64, b float64) *Neuron {
	out := Neuron{weight: w, bias: b}
	return &out
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Neuron		| get funcs
//----- ----- ----- ----- ----- ----- ----- -----*/

func (n Neuron) GetWeight() float64 {
	return n.weight
}

func (n Neuron) GetBias() float64 {
	return n.bias
}

func (n Neuron) getActivationGradient(input float64) float64 {
	// for Linear or Identity activation functions
	return 1.0

	// for ReLU activation functions
	/*
		if x > 0.0
			return 1.0;

		return 0.0
	*/
}

func (n Neuron) getActivation(input float64) float64 {
	// for Linear or Identity activation functions
	return input

	// for ReLU activation functions
	// ..
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Neuron		| run funcs
//----- ----- ----- ----- ----- ----- ----- -----*/

func (n *Neuron) FeedForward(input float64) float64 {
	// output y = f( \sigma )
	//	\sigma = weight * input * x + bias
	// for multiple inputs,
	//	\sigma = ( weight_0 * x_0 ) + ( weight_1 * x_0 )

	sigma := n.weight*n.input + n.bias
	n.input = input

	n.output = n.getActivation(sigma)

	return n.output
}

func (n *Neuron) BackPropagation(target float64) {
	alpha := 0.1 // Learning rate
	gradient := (n.output - target) * n.getActivationGradient(n.output)

	n.weight -= alpha * gradient * n.input
	n.bias -= alpha * gradient * 1.0
}
