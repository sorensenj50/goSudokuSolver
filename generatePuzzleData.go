package main

func (puzzle *Puzzle) generateRandomData() {
	puzzle.cellValues.iterate(puzzle.insertWithinConstraints, false)
}

func (puzzle *Puzzle) insertWithinConstraints(row, col int) {
	jointPossibilities := puzzle.getJointPossibilities(row, col, calculateBlockNumber(row, col))
	topValue := jointPossibilities.pop(0)

	puzzle.cellValues.insert(row, col, topValue)
	puzzle.addConstraintsHelper(row, col)
}

func (puzzle *Puzzle) backTrackInsertion() {
	row := 0
	col := 0
	popAdjust := 0

	for {

		jointPossibilities := puzzle.getJointPossibilities(row, col, calculateBlockNumber(row, col))
		topValue := jointPossibilities.pop(popAdjust)

		if topValue == 0 {
			row, col = moveIndicesBackward(row, col)
			puzzle.removeConstraint(row, col)
			popAdjust++
		} else {
			puzzle.cellValues.insert(row, col, topValue)
			puzzle.addConstraintsHelper(row, col)

			row, col = moveIndicesForward(row, col)
			popAdjust = 0
			if row == 8 && col == 8 {
				break
			}
		}
	}
}
