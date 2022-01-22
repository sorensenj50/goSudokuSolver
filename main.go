package main

import (
	"fmt"
)

const gridSize = 9
const gridDimensions = 2
const minNum = 1
const maxNum = 9

var exists = struct{}{}

func main() {
	fmt.Println("Hello World")

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

	puzzle := makePuzzle(samplePuzzleValues)
	//puzzle = puzzle.generateRandomSolution()

	newPuzzle := testPuzzle(puzzle)
	displayGrid(puzzle.cellValues)
	displayGrid(newPuzzle.cellValues)

}

func testPuzzle(puzzle Puzzle) Puzzle {
	puzzle.cellValues[0][0] = 100
	return puzzle
}
