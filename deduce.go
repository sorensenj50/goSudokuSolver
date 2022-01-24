package main

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

	numPossibilities, value := puzzle.calculateJointPossibilities(row, col, calculateBlockNumber(row, col))
	if numPossibilities == 1 {
		puzzle.cellValues.insert(row, col, value)
		puzzle.addConstraintsHelper(row, col)
		decider.ensureTrue()

	}
}

func (puzzle *Puzzle) calculateJointPossibilities(row, col, block int) (numOfPossibilities, value int) {
	rowPossibilities := puzzle.rowSets.getPossibilities(row)
	colPossibilities := puzzle.colSets.getPossibilities(col)
	blockPossibilities := puzzle.blockSets.getPossibilities(block)

	jointPossibilities := rowPossibilities.intersection(colPossibilities, blockPossibilities)

	return jointPossibilities.getLength(), jointPossibilities.pop()
}

func newTrue() *bool {
	trueValue := true
	return &trueValue
}

//func (puzzle *Puzzle) adjustNestedSets(row, col, value int) {
//	puzzle.rowSets.insert(row, value)
//	puzzle.colSets.insert(col, value)
//	puzzle.blockSets.insert(calculateBlockNumber(row, col), value)
//}
//
//func (puzzle *Puzzle) updatePossibilityMatrix(row, col, value int) {
//
//}
