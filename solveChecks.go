package main

import (
	"fmt"
	"strconv"
)

const (
	duplicateNums PuzzleResult = iota
	tooFewNums
)

func displayGrid(grid [gridSize][gridSize]int) {
	fmt.Println()
	fmt.Println("  Puzzle")
	string := "  "
	for i := range [gridSize]int{} {
		for j := range [gridSize]int{} {
			string += strconv.Itoa(grid[i][j])
			if j == 2 || j == 5 {
				string += "  "
				string += "|"
				string += "  "
			} else {
				string += " "
			}
		}
		string += "\n  "
		if i == 2 || i == 5 {
			string += "--------------------------"
			string += "\n  "
		}
	}
	fmt.Print(string)
}

func (puzzle Puzzle) checkSolved() bool {
	for i := range [gridSize]int{} {
		incorrectRow := len(puzzle.rowSets[i]) != 9
		incorrectCol := len(puzzle.rowSets[i]) != 9
		incorrectBlock := len(puzzle.colSets[i]) != 9

		if incorrectRow || incorrectCol || incorrectBlock {
			return false
		}
	}
	return true
}
