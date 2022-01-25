package main

import "fmt"

type Puzzle struct {
	cellValues Grid

	rowSets   SetCollection
	colSets   SetCollection
	blockSets SetCollection
}

func makePuzzle() Puzzle {
	var puzzle = Puzzle{}
	puzzle.cellValues = makeGrid()
	puzzle.rowSets = makeSets("row")
	puzzle.colSets = makeSets("col")
	puzzle.blockSets = makeSets("block")

	return puzzle
}

func makeRandomPuzzle() Puzzle {
	puzzle := makePuzzle()
	puzzle.generateRandomData()
	return puzzle
}

func makeSetPuzzle(data [gridSize][gridSize]int) Puzzle {
	puzzle := makePuzzle()
	puzzle.cellValues.setGrid(data)
	return puzzle
}

func (puzzle *Puzzle) addConstraints() {
	puzzle.cellValues.iterate(puzzle.addConstraintsHelper, true)
}

func (puzzle *Puzzle) addConstraintsHelper(row, col int) {
	value := puzzle.cellValues.get(row, col)
	fmt.Println("Value Adding to Constraints ", value)
	puzzle.rowSets.addConstraint(row, value)
	puzzle.colSets.addConstraint(col, value)
	puzzle.blockSets.addConstraint(calculateBlockNumber(row, col), value)
}