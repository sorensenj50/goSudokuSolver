package main

func (puzzle *Puzzle) generateRandomData() {
	puzzle.cellValues.iterate(puzzle.insertWithinConstraints)
}

func (puzzle *Puzzle) insertWithinConstraints(row, col int) {
	jointPossibilities := puzzle.getJointPossibilities(row, col, calculateBlockNumber(row, col))
	topValue := jointPossibilities.pop()
	puzzle.cellValues.insert(row, col, topValue)
	puzzle.addConstraintsHelper(row, col)
}
