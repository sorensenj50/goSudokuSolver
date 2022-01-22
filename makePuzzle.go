package main

var exists = struct{}{}

type Puzzle struct {
	initialValues   [9][9]int
	valuesWithInput [9][9]int
	neededSpaces    [][2]int

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
	puzzle.initialValues = data
	puzzle.valuesWithInput = data
	puzzle.rowSets = rowSets
	puzzle.colSets = colSets
	puzzle.blockSets = blockSets
	puzzle.neededSpaces = neededSpaces

	return puzzle
}
