package QL

import "fmt"

type GridWorld struct {
	width, height int
	cell          []Cell
}

func CreateGridWorld(width, height int) *GridWorld {
	return &GridWorld{
		cell:   make([]Cell, width*height),
		width:  width,
		height: height,
	}
}

func (g *GridWorld) GetCell(x, y int) *Cell {
	if g.isInside(x, y) == false {
		panic("Out of Grid Range")
	}

	return &g.cell[x+(g.width*y)]
}

func (g *GridWorld) isInside(x, y int) bool {
	if x >= g.width && x < 0 {
		return false
	}
	if y >= g.height && y < 0 {
		return false
	}

	return true
}

func printSigned(a float64) {
	if a > 0 {
		fmt.Printf(" +%1.1f  ", a)

	} else if a < 0 {
		fmt.Printf("  %1.1f  ", a)

	} else { // a == 0
		fmt.Printf("  0.0  ")

	}
}

func (g *GridWorld) Print() {

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			fmt.Print("----------------------")
		}
		fmt.Printf("-\n")
		/*----------------------------------------
		 * Drawing Cell -> Up
		 ----------------------------------------*/
		for x := 0; x < g.width; x++ {
			fmt.Print("|       ") // left blank

			printSigned(
				g.GetCell(x, y).GetQValue(Up),
			)

			fmt.Print("       ") // right blank
		}

		fmt.Printf("|\n")
		/*----------------------------------------
		 * Drawing Cell -> Left, Right
		 ----------------------------------------*/
		fmt.Print("|")
		for x := 0; x < g.width; x++ {
			printSigned( // print left
				g.GetCell(x, y).GetQValue(Left),
			)

			fmt.Print("       ") // middle blank

			printSigned( // print right
				g.GetCell(x, y).GetQValue(Right),
			)

			fmt.Print("|")
		}

		fmt.Printf("\n")
		/*----------------------------------------
		 * Drawing Cell -> Down
		 ----------------------------------------*/
		for x := 0; x < g.width; x++ {
			fmt.Print("|       ") // left blank

			printSigned(
				g.GetCell(x, y).GetQValue(Down),
			)

			fmt.Print("       ") // right blank
		}

		fmt.Printf("|\n")
	}

	for x := 0; x < g.width; x++ {
		fmt.Print("----------------------")
	}
	fmt.Printf("-\n")
}
