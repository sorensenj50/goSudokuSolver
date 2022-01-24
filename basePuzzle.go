package main

type Puzzle struct {
	cellValues Grid

	rowSets   SetCollection
	colSets   SetCollection
	blockSets SetCollection
}

func makePuzzle(data [gridSize][gridSize]int) Puzzle {
	var puzzle = Puzzle{}
	puzzle.cellValues = makeGrid(data)
	puzzle.rowSets = makeSets("row")
	puzzle.colSets = makeSets("col")
	puzzle.blockSets = makeSets("block")

	return puzzle
}

func (puzzle *Puzzle) addConstraints() {
	puzzle.cellValues.iterate(puzzle.addConstraintsHelper)
}

func (puzzle *Puzzle) addConstraintsHelper(row, col int) {
	value := puzzle.cellValues.get(row, col)
	if value != 0 {
		puzzle.rowSets.addConstraint(row, value)
		puzzle.colSets.addConstraint(col, value)
		puzzle.blockSets.addConstraint(calculateBlockNumber(row, col), value)
	}
}
