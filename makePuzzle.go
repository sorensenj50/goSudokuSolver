package main

type PuzzleResult int

type Puzzle struct {
	cellValues   [gridSize][gridSize]int
	neededSpaces [][gridDimensions]int

	rowSets   map[int]map[int]struct{}
	colSets   map[int]map[int]struct{}
	blockSets map[int]map[int]struct{}
}

func makePuzzle(data [gridSize][gridSize]int) Puzzle {
	var puzzle Puzzle

	rowSets := make(map[int]map[int]struct{})
	colSets := make(map[int]map[int]struct{})
	blockSets := make(map[int]map[int]struct{})
	neededSpaces := [][gridDimensions]int{}

	for i := range [9]int{} {
		rowSets[i] = make(map[int]struct{})
		colSets[i] = make(map[int]struct{})
		blockSets[i] = make(map[int]struct{})
	}

	for row := range [gridSize]int{} {
		for col := range [gridSize]int{} {
			value := data[row][col]
			if value == 0 {
				neededSpaces = append(neededSpaces, [gridDimensions]int{row, col})
			} else { // only want non-zero values in sets
				rowSets[row][value] = exists
				colSets[col][value] = exists
				blockSets[calculateBlockNumber(row, col)][value] = exists
			}
		}
	}
	puzzle.cellValues = data

	puzzle.neededSpaces = neededSpaces

	puzzle.rowSets = rowSets
	puzzle.colSets = colSets
	puzzle.blockSets = blockSets

	return puzzle
}

func calculateBlockNumber(row int, col int) int {
	return (row / 3) + (col/3)*3
}
