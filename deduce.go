package main

//func (puzzle *Puzzle) deduce() {
//	didDeduce := false
//	for i := range puzzle.neededSpaces.array {
//		row := puzzle.neededSpaces.getRowNumber(i)
//		col := puzzle.neededSpaces.getColumnNumber(i)
//
//		possibilities := puzzle.matrix.matrix[row][col]
//		if len(possibilities) == 1 {
//			puzzle.cellValues.insert(row, col, possibilities[0])
//			didDeduce = true
//		} else {
//			didDeduce = false
//		}
//	}
//}
//
//func (puzzle *Puzzle) adjustNestedSets(row, col, value int) {
//	puzzle.rowSets.insert(row, value)
//	puzzle.colSets.insert(col, value)
//	puzzle.blockSets.insert(calculateBlockNumber(row, col), value)
//}
//
//func (puzzle *Puzzle) updatePossibilityMatrix(row, col, value int) {
//
//}
