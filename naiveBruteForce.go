package main

import (
	"math/rand"
	"time"
)

func runNaiveRandom(puzzle Puzzle) Puzzle {
	continueLooping := true
	for continueLooping {
		newSolution := puzzle.generateRandomSolution()
		if puzzle.checkSolved() {
			continueLooping = false
			return newSolution
		}
	}
	return Puzzle{}
}

func (puzzle Puzzle) generateRandomSolution() Puzzle {
	for i := range puzzle.neededSpaces {
		cellLocation := puzzle.neededSpaces[i]
		row := cellLocation[0]
		col := cellLocation[1]

		randomNumber := generateRandomNumber()
		puzzle.cellValues[row][col] = randomNumber
		puzzle.rowSets[row][randomNumber] = exists
		puzzle.colSets[col][randomNumber] = exists
		puzzle.blockSets[calculateBlockNumber(row, col)][randomNumber] = exists

	}
	return puzzle
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNum-minNum+1) + minNum
}
