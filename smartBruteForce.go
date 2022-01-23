package main

import (
	"math/rand"
)

type PossibilityMatrix struct {
	matrix [9][9][]int
}

func makePossibilityMatrix() PossibilityMatrix {
	array := [9][9][]int{}
	var matrix PossibilityMatrix
	matrix.matrix = array
	return matrix

}

func (puzzle *Puzzle) fillPossibilityMatrix() PossibilityMatrix {
	matrix := makePossibilityMatrix()
	rowPossibilities := puzzle.givenRowSets.getDifferenceSet()
	colPossibilities := puzzle.givenColSets.getDifferenceSet()
	blockPossibilities := puzzle.givenBlockSets.getDifferenceSet()

	for i := range puzzle.neededSpaces.array {
		row := puzzle.neededSpaces.getRowNumber(i)
		col := puzzle.neededSpaces.getColumnNumber(i)

		rowSet := rowPossibilities.getValues(row)
		colSet := colPossibilities.getValues(col)
		blockSet := blockPossibilities.getValues(calculateBlockNumber(row, col))
		matrix.matrix[row][col] = mergeDifferenceSets(rowSet, colSet, blockSet)
	}
	return matrix
}

func (matrix *PossibilityMatrix) getRandomConstrainedNumber(row, col int) int {
	arrayLength := len(matrix.matrix[row][col])
	randomIndex := getRandomIndex(arrayLength)
	return matrix.matrix[row][col][randomIndex]
}

func getRandomIndex(length int) int {
	setRandomSeed()
	return rand.Intn(length)
}

func (puzzle *Puzzle) generateConstrainedRandomSolution() {

	for i := range puzzle.neededSpaces.array {
		row := puzzle.neededSpaces.getRowNumber(i)
		col := puzzle.neededSpaces.getColumnNumber(i)

		number := puzzle.matrix.getRandomConstrainedNumber(row, col)
		puzzle.cellValues.insert(row, col, number)
		puzzle.rowSets.insert(row, number)
		puzzle.colSets.insert(col, number)
		puzzle.blockSets.insert(calculateBlockNumber(row, col), number)
	}

}

func (puzzle *Puzzle) runConstrainedRandom() {
	continueLooping := true
	for continueLooping {
		puzzle.generateConstrainedRandomSolution()
		if puzzle.checkSolved() {
			continueLooping = false
		} else {
			puzzle.reset()
		}
	}
}
