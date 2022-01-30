package main

import "fmt"

func (puzzle *Puzzle) generateRandomData() {
	puzzle.cellValues.iterate(puzzle.insertWithinConstraints, false)
}

func (puzzle *Puzzle) insertWithinConstraints(row, col int) {
	jointPossibilities := puzzle.getJointPossibilities(row, col, calculateBlockNumber(row, col))
	topValue := jointPossibilities.pop()

	puzzle.cellValues.insert(row, col, topValue)
	puzzle.addConstraintsHelper(row, col)
}

func (puzzle *Puzzle) backTrackInsertion() {
	for rowNum := range [gridSize]int{} {
		for colNum := range [gridSize]int{} {
			puzzle.backTrackRecursive(rowNum, colNum)
		}
	}
}

func (puzzle *Puzzle) backTrackRecursive(row, col int) {
	jointPossibilities := puzzle.getJointPossibilities(row, col, calculateBlockNumber(row, col))
	topValue := jointPossibilities.pop()

	fmt.Print("At: ", row, col)
	puzzle.cellValues.display("")

	if topValue != 0 {
		puzzle.cellValues.insert(row, col, topValue)
		puzzle.addConstraintsHelper(row, col)
	} else {
		fmt.Print("Recursive Call Made")
		puzzle.backTrackRecursive(getPreviousRowAndCol(row, col))
	}
}

func getPreviousRowAndCol(row, col int) (newRow, newCol int) {
	if col == 0 {
		return row - 1, 8
	}
	return row, col - 1
}
