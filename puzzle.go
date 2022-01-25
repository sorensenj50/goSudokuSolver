package main

type Puzzle struct {
	cellValues Grid

	rowSets   SetCollection
	colSets   SetCollection
	blockSets SetCollection
}

func makePuzzle() Puzzle {
	var puzzle = Puzzle{}
	puzzle.cellValues = makeGrid()
	puzzle.rowSets = makeSets("row")
	puzzle.colSets = makeSets("col")
	puzzle.blockSets = makeSets("block")

	puzzle.generateRandomData()
	return puzzle
}

func (puzzle *Puzzle) addConstraintsHelper(row, col int) {
	value := puzzle.cellValues.get(row, col)
	puzzle.rowSets.addConstraint(row, value)
	puzzle.colSets.addConstraint(col, value)
	puzzle.blockSets.addConstraint(calculateBlockNumber(row, col), value)
}
