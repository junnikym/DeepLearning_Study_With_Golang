package math

import (
	"fmt"
)

type Matrix struct {
	row, col int
	value    []float64
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Matrix		| constructor
//----- ----- ----- ----- ----- ----- ----- -----*/

// CreateMatrix : Create Only Matrix; this function do not substitution variable
//				 If you want substitution variable use 'NewMatrix' fuction
func CreateMatrix(row, col int) *Matrix {
	return &Matrix{
		value: make([]float64, row*col),
		row:   row,
		col:   col,
	}
}

// NewMatrix : Create New Matrix
func NewMatrix(row, col int, value []float64) *Matrix {

	// index out of range detective
	if row <= 0 || col <= 0 {
		if row == 0 || col == 0 {
			panic(ErrOverAllocateRange)
		}
		panic(ErrUnderAllocateRange)
	}

	// when matrix allocated array size is not equal with
	// argument row and col generate error
	if value != nil && row*col != len(value) {
		panic(ErrNotEqualShape)
	}

	if value == nil {
		value = make([]float64, row*col)
	}

	return &Matrix{
		value: value,
		row:   row,
		col:   col,
	}
}

func (m *Matrix) Init(row, col int) {
	m.value = make([]float64, row*col)
	m.row = row
	m.col = col
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Matrix		| getter
//----- ----- ----- ----- ----- ----- ----- -----*/

func (m Matrix) At(r, c int) float64 {
	return m.value[(r*m.col + c)]
}

func (m Matrix) Row() int {
	return m.row
}

func (m Matrix) Col() int {
	return m.col
}

func (m Matrix) Print() {
	var calced_row int = 0

	for r := 0; r < m.row; r++ {
		calced_row = r * m.col

		for c := 0; c < m.col; c++ {
			fmt.Print(" ", m.value[calced_row+c], ",")
		}
		fmt.Print("\n")
	}
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Matrix		| setter
//----- ----- ----- ----- ----- ----- ----- -----*/

func (m *Matrix) ReplaceAt(r, c int, value float64) {
	m.value[(r*m.col + c)] = value
}

/*----- ----- ----- ----- ----- ----- ----- -----//
	Matrix		| Calcuate function
//----- ----- ----- ----- ----- ----- ----- -----*/

func (m *Matrix) Mul(rhs, result *Vector) {
	if m.row > result.size {
		panic(ErrNotEqualShape)
	}
	if m.col > rhs.size {
		panic(ErrNotEqualShape)
	}

	var i int = 0

	for r := 0; r < m.row; r++ {
		result.value[r] = 0.0

		i = r * m.row

		for c := 0; c < m.col; c++ {
			result.value[r] += m.value[i] * rhs.value[c]
			i++
		}
	}
}

func (m *Matrix) MulTrans(rhs, result *Vector) {
	if m.row > rhs.size {
		panic(ErrNotEqualShape)
	}
	if m.col > result.size {
		panic(ErrNotEqualShape)
	}

	var i int = 0

	for c := 0; c < m.col; c++ {
		result.value[c] = 0.0

		i = c

		for r := 0; r < m.row; r++ {
			result.value[c] += m.value[i] * rhs.value[r]
			i += m.col
		}
	}
}
