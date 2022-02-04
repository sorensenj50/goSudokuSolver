package main

//
//import "fmt"
//
//func (puzzle *Puzzle) backTrackInsertion() {
//	row := 0
//	col := 0
//	shouldBacktrack := false
//
//	for {
//		jointPossibilities := puzzle.getJointPossibilities(row, col, calculateBlockNumber(row, col))
//		topValue := jointPossibilities.pop()
//
//		fmt.Println(row, col, topValue)
//		puzzle.display()
//
//		shouldBacktrack = topValue == 0
//
//		fmt.Println("shouldBackTrack", shouldBacktrack)
//		fmt.Println("num keys", jointPossibilities.getNumKeys(true))
//
//		if shouldBacktrack {
//			row, col = moveIndicesBackward(row, col)
//
//			puzzle.removeConstraint(row, col)
//			puzzle.addFutureConstraint(row, col)
//			puzzle.cellValues.remove(row, col)
//
//		} else {
//			puzzle.cellValues.insert(row, col, topValue)
//			puzzle.addConstraintsHelper(row, col)
//			puzzle.future.reset()
//
//			row, col = moveIndicesForward(row, col)
//			if row == 8 && col == 8 {
//				break
//			}
//		}
//	}
//}
