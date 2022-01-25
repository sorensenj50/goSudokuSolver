package main

import "fmt"

func (puzzle *Puzzle) generateRandomData() {
	puzzle.cellValues.iterate(puzzle.insertWithinConstraints, false)
}

func (puzzle *Puzzle) insertWithinConstraints(row, col int) {
	fmt.Print(row, col, " ")
	jointPossibilities := puzzle.getJointPossibilities(row, col, calculateBlockNumber(row, col))
	topValue := jointPossibilities.pop()

	fmt.Println("Joint Possibilities")
	fmt.Println(jointPossibilities)
	fmt.Print(topValue, " ")
	fmt.Println()

	puzzle.cellValues.insert(row, col, topValue)
	puzzle.addConstraintsHelper(row, col)
}
