package main

import "fmt"

type ContinueDecider struct {
	shouldContinue bool
	counter        int
}

func makeDecider() *ContinueDecider {
	var decider ContinueDecider
	shouldContinue := true
	decider.shouldContinue = shouldContinue
	return &decider
}
func (decider *ContinueDecider) ensureFalse() {
	decider.shouldContinue = false
}

func (decider *ContinueDecider) ensureTrue() {
	decider.shouldContinue = true
}

func (decider *ContinueDecider) increment() {
	decider.counter += 1
}

func (puzzle *Puzzle) deduce() {
	decider := makeDecider()
	for decider.shouldContinue {
		decider.ensureFalse()
		for rowNum := range [gridSize]int{} {
			for colNum := range [gridSize]int{} {
				puzzle.deduceHelper(rowNum, colNum, decider)
			}
		}
	}
}

func (puzzle *Puzzle) deduceHelper(row, col int, decider *ContinueDecider) {
	if puzzle.cellValues.get(row, col) != 0 {
		return
	}

	jointPossibilities := puzzle.getJointPossibilities(row, col, calculateBlockNumber(row, col))

	fmt.Println("Possibilities for ", row, col)
	jointPossibilities.display()
	if jointPossibilities.getNumKeys(true) == 1 {
		puzzle.cellValues.insert(row, col, jointPossibilities.pop())
		puzzle.addConstraintsHelper(row, col)
		decider.ensureTrue()
	}
}

func (puzzle *Puzzle) getJointPossibilities(row, col, block int) *ArraySet {
	rowSet := puzzle.rowSets.sets[row]
	colSet := puzzle.colSets.sets[col]
	blockSet := puzzle.blockSets.sets[block]

	return rowSet.intersection(true, colSet, blockSet)
}
