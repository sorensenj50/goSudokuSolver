package main

import (
	"math/rand"
	"time"
)

func (puzzle *Puzzle) runNaiveRandom() {
	continueLooping := true
	for continueLooping {
		puzzle.generateRandomSolution()
		if puzzle.checkSolved() {
			continueLooping = false
		} else {
			puzzle.reset()
		}
	}
}

func (puzzle *Puzzle) generateRandomSolution() {
	for i := range puzzle.neededSpaces.array {
		cellLocation := puzzle.neededSpaces.array[i]
		row := cellLocation[0]
		col := cellLocation[1]

		randomNumber := generateRandomNumber()
		puzzle.cellValues.insert(row, col, randomNumber)
		puzzle.rowSets.insert(row, randomNumber)
		puzzle.colSets.insert(col, randomNumber)
		puzzle.blockSets.insert(calculateBlockNumber(row, col), randomNumber)
	}
}

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNum-minNum+1) + minNum
}
