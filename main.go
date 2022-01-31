package main

const gridSize = 9

var numsArray = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

//var exists = struct{}{}

func main() {
	randomPuzzle()
	//samplePuzzle()
	//test()

}

func samplePuzzle() {
	samplePuzzleValues := [gridSize][gridSize]int{
		{0, 0, 0, 2, 6, 0, 7, 0, 1},
		{6, 8, 0, 0, 7, 0, 0, 9, 0},
		{1, 9, 0, 0, 0, 4, 5, 0, 0},
		{8, 2, 0, 1, 0, 0, 0, 4, 0},
		{0, 0, 4, 6, 0, 2, 9, 0, 0},
		{0, 5, 0, 0, 0, 3, 0, 2, 8},
		{0, 0, 9, 3, 0, 0, 0, 7, 4},
		{0, 4, 0, 0, 5, 0, 0, 3, 6},
		{7, 0, 3, 0, 1, 8, 0, 0, 0},
	}

	puzzle := makeSetPuzzle(samplePuzzleValues)
	puzzle.cellValues.display()

	puzzle.addConstraints()
	puzzle.deduce()

	puzzle.rowSets.displayAll()

	puzzle.cellValues.display()
}

func randomPuzzle() {
	puzzle := makePuzzle()
	puzzle.backTrackInsertion()

	puzzle.cellValues.display()
}

func test() {
	samplePuzzleValues := [gridSize][gridSize]int{
		{0, 0, 0, 2, 6, 0, 7, 0, 1},
		{6, 8, 0, 0, 7, 0, 0, 9, 0},
		{1, 9, 0, 0, 0, 4, 5, 0, 0},
		{8, 2, 0, 1, 0, 0, 0, 4, 0},
		{0, 0, 4, 6, 0, 2, 9, 0, 0},
		{0, 5, 0, 0, 0, 3, 0, 2, 8},
		{0, 0, 9, 3, 0, 0, 0, 7, 4},
		{0, 4, 0, 0, 5, 0, 0, 3, 6},
		{7, 0, 3, 0, 1, 8, 0, 0, 0},
	}

	puzzle := makeSetPuzzle(samplePuzzleValues)
	puzzle.cellValues.insert(0, 4, 9)
	puzzle.display()

}
