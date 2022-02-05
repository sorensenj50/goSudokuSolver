package main

const gridSize = 9

func main() {
	newPuzzle := makeNewPuzzle()
	newPuzzle.fillGrid()

	newPuzzle.makeGaps(0.5)

	newPuzzle.display()

	newPuzzle.markGiven()

	newPuzzle.fillGrid()

}
