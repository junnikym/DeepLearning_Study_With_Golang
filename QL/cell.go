package QL

import "learning/QL/compare"

type Cell struct {
	qValue [4]float64
	reward float64
}

const CellOption = 4

const ( // Options
	Up = iota
	Down
	Left
	Right
)

func (c *Cell) Reset() {
	for i := range c.qValue {
		c.qValue[i] = 0
	}

	c.reward = 0.0
}

func (c *Cell) GetQValue(i int) float64 {
	return c.qValue[i]
}

func (c *Cell) MaxQ() float64 {
	return compare.Max4(c.qValue[0], c.qValue[1], c.qValue[2], c.qValue[3])
}

func CellMover(x, y *int, option int) {
	if option == Up {
		*y += 1

	} else if option == Down {
		*y -= 1

	} else if option == Left {
		*x += 1

	} else if option == Right {
		*x -= 1

	} else {
		panic("Out of opation's range")
	}
}
