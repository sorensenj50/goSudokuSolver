package main

//
//type Puzzle struct {
//	cellValues Grid
//
//	rowSets   SetCollection
//	colSets   SetCollection
//	blockSets SetCollection
//	future    *GridSetCollection
//}
//
//func makePuzzle() Puzzle {
//	var puzzle = Puzzle{}
//	puzzle.cellValues = makeGrid()
//	puzzle.rowSets = makeSets("row")
//	puzzle.colSets = makeSets("col")
//	puzzle.blockSets = makeSets("block")
//	puzzle.future = makeGridSet()
//
//	return puzzle
//}
//
//func makeSetPuzzle(data [gridSize][gridSize]int) Puzzle {
//	puzzle := makePuzzle()
//	puzzle.cellValues.setGrid(data)
//	return puzzle
//}
//
//func (puzzle *Puzzle) addConstraints() {
//	puzzle.cellValues.iterate(puzzle.addConstraintsHelper, true)
//}
//
//func (puzzle *Puzzle) addConstraintsHelper(row, col int) {
//	value := puzzle.cellValues.get(row, col)
//	puzzle.rowSets.addConstraint(row, indexAdjust(value))
//	puzzle.colSets.addConstraint(col, indexAdjust(value))
//	puzzle.blockSets.addConstraint(calculateBlockNumber(row, col), indexAdjust(value))
//}
//
//func (puzzle *Puzzle) display() {
//	puzzle.cellValues.display()
//}
//
//func (puzzle *Puzzle) removeConstraint(row, col int) {
//	value := puzzle.cellValues.get(row, col)
//	puzzle.rowSets.removeConstraint(row, indexAdjust(value))
//	puzzle.colSets.removeConstraint(col, indexAdjust(value))
//	puzzle.blockSets.removeConstraint(calculateBlockNumber(row, col), indexAdjust(value))
//}
//
//func (puzzle *Puzzle) addFutureConstraint(row, col int) {
//	num := puzzle.cellValues.get(row, col)
//	puzzle.future.addConstraint(row, col, indexAdjust(num))
//}
//
//// sudoku is 1 indexed but arrays are zero indexed
//func indexAdjust(num int) int {
//	if num == 0 {
//		return 0
//	} else {
//		return num - 1
//	}
//}
