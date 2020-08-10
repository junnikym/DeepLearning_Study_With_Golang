package QL

import "math/rand"

type Agent struct {
	x, y   int     // position
	reward float64 // reward point
}

func CreateAgent() *Agent {
	return &Agent{0, 0, 0}
}

func (a *Agent) MoveRandomly() {
	CellMover(&a.x, &a.y, rand.Intn(CellOption))
}
