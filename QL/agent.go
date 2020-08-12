package QL

import (
	"fmt"
	"math/rand"
)

type Agent struct {
	x, y   int     // position
	reward float64 // reward point

	// for E-Greedy
	exploration float64
}

func CreatAgent() *Agent {
	return &Agent{0, 0, 0, 0}
}

func (a *Agent) Reset() {
	a.x = 0
	a.y = 0
	a.reward = 0
}

func (a *Agent) MoveRandomly() int {
	direction := rand.Intn(CellOption)

	CellMover(&a.x, &a.y, direction)

	return direction
}

func (a *Agent) GetPosition() (int, int) {
	return a.x, a.y
}

func (a *Agent) PrintPosition() {
	fmt.Println(" Agent position : ( ", a.x, ", ", a.y, " )")
}

/*--------------------------------------------------
 * E-Greedy Action
 --------------------------------------------------*/

func (a *Agent) E_GreedyAction(currentCell *Cell) int {
	if rand.Float64() < a.exploration {
		return a.MoveRandomly()
	} else {
		direction := currentCell.MaxQ_Direction()
		CellMover(&a.x, &a.y, direction)

		return direction
	}
}

func (a *Agent) SetExploation(e float64) {
	a.exploration = e
}

func (a *Agent) GetExploation() float64 {
	return a.exploration
}
