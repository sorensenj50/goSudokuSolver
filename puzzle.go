package main

import (
	"fmt"
	"strconv"
)

var exists = struct{}{}

type Puzzle struct {
	cellValues   [9][9]int
	neededSpaces [][2]int

	rowSets   map[int]map[int]struct{}
	colSets   map[int]map[int]struct{}
	blockSets map[int]map[int]struct{}
}

func makePuzzle(data [9][9]int) Puzzle {
	var puzzle Puzzle

	rowSets := make(map[int]map[int]struct{})
	colSets := make(map[int]map[int]struct{})
	blockSets := make(map[int]map[int]struct{})
	neededSpaces := [][2]int{}

	for i := range [9]int{} {
		rowSets[i] = make(map[int]struct{})
		colSets[i] = make(map[int]struct{})
		blockSets[i] = make(map[int]struct{})
	}

	for j := range [9]int{} {
		for k := range [9]int{} {
			value := data[j][k]
			rowSets[j][value] = exists
			colSets[k][value] = exists
			blockSets[calculateBlockNumber(j, k)][value] = exists
			if value == 0 {
				neededSpaces = append(neededSpaces, [2]int{j, k})
			}
		}
	}
	puzzle.cellValues = data
	puzzle.rowSets = rowSets
	puzzle.colSets = colSets
	puzzle.blockSets = blockSets
	puzzle.neededSpaces = neededSpaces

	return puzzle
}

func calculateBlockNumber(i int, j int) int {
	return (i / 3) + (j/3)*3
}

func (puzzle Puzzle) display() {
	fmt.Println()
	fmt.Println("  Puzzle")
	string := "  "
	for i := range [9]int{} {
		for j := range [9]int{} {
			string += strconv.Itoa(puzzle.cellValues[i][j])
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
