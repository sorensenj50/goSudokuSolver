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
}

func makePuzzle(data [gridSize][gridSize]int) Puzzle {
	var puzzle Puzzle

	neededSpaces := makeCoordinates()
	nestedRowSets := makeSet()
	nestedColSets := makeSet()
	nestedBlockSets := makeSet()

	for row := range [gridSize]int{} {
		for col := range [gridSize]int{} {
			value := data[row][col]
			if value == 0 {

			} else { // only want non-zero values in sets
				nestedRowSets.insert(row, value)
				nestedColSets.insert(col, value)
				nestedBlockSets.insert(calculateBlockNumber(row, col), value)
			}
		}
	}

	puzzle.cellValues = makeGrid(data)
	puzzle.givenValues = makeGrid(data)

	puzzle.neededSpaces = neededSpaces

	puzzle.rowSets = nestedRowSets
	puzzle.colSets = nestedColSets
	puzzle.blockSets = nestedBlockSets

	puzzle.givenRowSets = nestedRowSets
	puzzle.givenColSets = nestedColSets
	puzzle.blockSets = nestedBlockSets

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
