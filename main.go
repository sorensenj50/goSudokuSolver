package main

const gridSize = 9

var numsArray = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

//var exists = struct{}{}

func main() {
	test()

}

func samplePuzzle() {
	samplePuzzleValues := getPuzzleData("puzzleTwo")

	puzzle := makeSetPuzzle(samplePuzzleValues)
	puzzle.cellValues.display()

	puzzle.addConstraints()
	puzzle.deduce()

	puzzle.rowSets.displayAll()

	puzzle.cellValues.display()
}

func test() {
	newPuzzle := makeNewPuzzle()

	newPuzzle.isValidBlock()
}

func randomTest() {
	puzzle := makePuzzle()
	puzzle.backTrackInsertion()
}
