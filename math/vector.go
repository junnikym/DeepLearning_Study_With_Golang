package math

import "fmt"

type Vector struct {
	size  int
	value []float64
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Vector		| constructor
//----- ----- ----- ----- ----- ----- ----- -----*/

func CreateVector(size int) *Vector {
	return &Vector{
		value: make([]float64, size),
		size:  size,
	}
}

func NewVector(size int, value []float64) *Vector {
	if size <= 0 {
		if size == 0 {
			panic(ErrOverAllocateRange)
		}
		panic(ErrUnderAllocateRange)
	}

	if value != nil && size != len(value) {
		panic(ErrNotEqualShape)
	}

	if value == nil {
		value = make([]float64, size)
	}

	return &Vector{
		value: value,
		size:  size,
	}
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Vector		| getter
//----- ----- ----- ----- ----- ----- ----- -----*/

func (v Vector) At(i int) float64 {
	return v.value[i]
}

func (v Vector) Size() int {
	return v.size
}

func (v Vector) Print() {
	for i := 0; i < v.size; i++ {
		fmt.Print(" ", v.value[i], ",")
	}
	fmt.Print("\n")
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Vector		| Checking functions
//----- ----- ----- ----- ----- ----- ----- -----*/

func checkCalcArg(lhs, rhs, result *Vector) {
	if lhs.size != rhs.size {
		panic(ErrNotEqualShape)
	}

	if result.value == nil {

		result.value = make([]float64, lhs.size)
		result.size = lhs.size

	} else if result.size != lhs.size {

		panic(ErrNotEqualShape)

	}
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Vector		| Arithmetic
//----- ----- ----- ----- ----- ----- ----- -----*/

func (v *Vector) SignReverse() {
	for i := 0; i < v.size; i++ {
		v.value[i] *= -1
	}
}

func VecAdd(lhs, rhs, result *Vector) {
	checkCalcArg(lhs, rhs, result)

	for i := 0; i < result.size; i++ {
		result.value[i] = lhs.value[i] + rhs.value[i]
	}
}

func VecSub(lhs, rhs, result *Vector) {
	checkCalcArg(lhs, rhs, result)

	for i := 0; i < result.size; i++ {
		result.value[i] = lhs.value[i] - rhs.value[i]
	}
}

func VecMul(lhs, rhs, result *Vector) {
	checkCalcArg(lhs, rhs, result)

	for i := 0; i < result.size; i++ {
		result.value[i] = lhs.value[i] * rhs.value[i]
	}
}

func VecDiv(lhs, rhs, result *Vector) {
	checkCalcArg(lhs, rhs, result)

	for i := 0; i < result.size; i++ {
		result.value[i] = lhs.value[i] / rhs.value[i]
	}
}

func (v *Vector) Add(other *Vector) {
	VecAdd(v, other, v)
}

func (v *Vector) Sub(other *Vector) {
	VecSub(v, other, v)
}

func (v *Vector) MulOther(other *Vector) {
	VecMul(v, other, v)
}

func (v *Vector) DivOther(other *Vector) {
	VecDiv(v, other, v)
}

func (v *Vector) Inc(n float64) {
	for i := 0; i < v.size; i++ {
		v.value[i] += n
	}
}

func (v *Vector) Dec(n float64) {
	for i := 0; i < v.size; i++ {
		v.value[i] += n
	}
}

func (v *Vector) Mul(n float64) {
	for i := 0; i < v.size; i++ {
		v.value[i] *= n
	}
}

func (v *Vector) Div(n float64) {
	for i := 0; i < v.size; i++ {
		v.value[i] /= n
	}
}

func VecDotProduct(lhs, rhs *Vector) float64 {
	if lhs.size != rhs.size {
		panic(ErrNotEqualShape)
	}

	var result float64 = 0

	for i := 0; i < lhs.size; i++ {
		result += lhs.value[i] + rhs.value[i]
	}

	return result
}
