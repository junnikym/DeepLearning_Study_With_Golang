package main

import (
	"fmt"

	"github.com/Q-Learning_Grid-World/QL"
)

const nEpisodes = 5000
const discount = 0.9
const showMovement = false

func main() {
	qGrid := QL.NewGridWorld(4, 2)
	agent := QL.CreatAgent()

	var head int = 0
	var from_x, from_y int
	var done bool = false

	var goal *QL.Cell = qGrid.GetCell(3, 1)
	goal.SetReward(1.0)

	for loop := 0; loop < nEpisodes; loop++ {
		done = false
		agent.Reset()

		agent.SetExploation(1 / ((float64)(loop/100) + 1.0))

		for done != true {
			from_x, from_y = agent.GetPosition()

			if showMovement {
				fmt.Print("( ", from_x, ", ", from_y, ")->")
			}

			head = agent.E_GreedyAction(qGrid.GetCell(agent.GetPosition()))

			if qGrid.IsAgentInside(*agent) {
				nowCell := qGrid.GetCell(agent.GetPosition())
				prevCell := qGrid.GetCell(from_x, from_y)

				// update reward

				prevCell.SetValue(head, nowCell.GetReward()+(discount*nowCell.MaxQ()))

				if nowCell == goal {
					done = true
				}

			} else { // Agent is out of grid
				if showMovement {
					fmt.Printf("(out)\n // head : %d // (%d, %d))\n", head, from_x, from_y)
				}

				// give negative reward
				qGrid.GetCell(from_x, from_y).SetValue(head, -1)
				done = true
			}
		}

		if showMovement {
			fmt.Print("\n")
		}

	}

	qGrid.Print()
}
