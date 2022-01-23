package main

type Puzzle struct {
	cellValues  Grid
	givenValues Grid

	neededSpaces Coordinates

	rowSets   NestedSet
	colSets   NestedSet
	blockSets NestedSet

	givenRowSets   NestedSet
	givenColSets   NestedSet
	givenBlockSets NestedSet

	matrix PossibilityMatrix
}

func makePuzzle(data [gridSize][gridSize]int) Puzzle {
	var puzzle Puzzle

	neededSpaces := makeCoordinates()
	nestedRowSets, nestedColSets, nestedBlockSets := puzzle.makeSets()

	for row := range [gridSize]int{} {
		for col := range [gridSize]int{} {
			value := data[row][col]

			if value == 0 { // only want non-zero values in sets
				neededSpaces.appendWrapper([2]int{row, col})
			} else {
				nestedRowSets.insert(row, value)
				nestedColSets.insert(col, value)
				nestedBlockSets.insert(calculateBlockNumber(row, col), value)
			}
		}
	}

	puzzle.establishGrid(makeGrid(data))
	puzzle.neededSpaces = neededSpaces
	puzzle.establishSets(nestedRowSets, nestedColSets, nestedBlockSets)

	puzzle.matrix = puzzle.fillPossibilityMatrix()

	return puzzle
}

func calculateBlockNumber(row int, col int) int {
	return (row / 3) + (col/3)*3
}

func (puzzle *Puzzle) checkSolved() bool {
	for i := range [gridSize]int{} {
		incorrectRow := !puzzle.rowSets.checkSolved(i)
		incorrectCol := !puzzle.colSets.checkSolved(i)
		incorrectBlock := !puzzle.givenBlockSets.checkSolved(i)

		if incorrectRow || incorrectCol || incorrectBlock {
			return false
		}
	}
	return true
}

func (puzzle *Puzzle) reset() {

	puzzle.rowSets.reset(&puzzle.givenRowSets)
	puzzle.colSets.reset(&puzzle.givenColSets)
	puzzle.blockSets.reset(&puzzle.givenBlockSets)

	puzzle.cellValues.reset(&puzzle.givenValues)
}

func (puzzle *Puzzle) establishSets(row, col, block NestedSet) {
	puzzle.rowSets = row
	puzzle.givenRowSets = row.copySet() // to avoid identical pointers

	puzzle.colSets = col
	puzzle.givenColSets = col.copySet()

	puzzle.blockSets = block
	puzzle.givenBlockSets = block.copySet()

}

func (puzzle *Puzzle) establishGrid(grid Grid) {
	puzzle.cellValues = grid
	puzzle.givenValues = grid
}

func (puzzle *Puzzle) makeSets() (row, col, block NestedSet) {
	return makeSet(), makeSet(), makeSet()
}
