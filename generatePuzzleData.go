package main

import "fmt"

func (puzzle *Puzzle) backTrackInsertion() {
	row := 0
	col := 0
	popAdjust := 0

	for {
		jointPossibilities := puzzle.getJointPossibilities(row, col, calculateBlockNumber(row, col))
		topValue := jointPossibilities.pop(popAdjust)

		fmt.Println(row, col, topValue)

		shouldBacktrack := topValue == 0 || (jointPossibilities.getNumKeys(true) == 1 && popAdjust >= 0)
		if shouldBacktrack {
			row, col = moveIndicesBackward(row, col)
			puzzle.removeConstraint(row, col)
			puzzle.cellValues.remove(row, col)
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
