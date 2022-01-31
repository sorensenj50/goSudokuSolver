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

	return puzzle
}

func makeSetPuzzle(data [gridSize][gridSize]int) Puzzle {
	puzzle := makePuzzle()
	puzzle.cellValues.setGrid(data)
	return puzzle
}

func (puzzle *Puzzle) addConstraints() {
	puzzle.cellValues.iterate(puzzle.addConstraintsHelper, true)
}

func (puzzle *Puzzle) addConstraintsHelper(row, col int) {
	value := puzzle.cellValues.get(row, col)
	puzzle.rowSets.addConstraint(row, value-1) // sodoku is 1 indexed but arrays are zero indexed
	puzzle.colSets.addConstraint(col, value-1)
	puzzle.blockSets.addConstraint(calculateBlockNumber(row, col), value-1)
}

func (puzzle *Puzzle) display() {
	puzzle.cellValues.display()
}

func (puzzle *Puzzle) removeConstraint(row, col int) {
	value := puzzle.cellValues.get(row, col)
	puzzle.rowSets.removeConstraint(row, value-1)
	puzzle.colSets.removeConstraint(col, value-1)
	puzzle.blockSets.removeConstraint(calculateBlockNumber(row, col), value-1)

}
