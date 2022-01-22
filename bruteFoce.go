package main

func (puzzle Puzzle) generateRandomSolution() {
	for i := range puzzle.neededSpaces {
		cellLocation := puzzle.neededSpaces[i]
		puzzle.valuesWithInput[cellLocation[0]][cellLocation[1]] = generateRandomNumber()
	}
}

func generateRandomNumber() int {
	return 2
}
