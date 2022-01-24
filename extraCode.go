package main

//import "image/color"
//
//type Cell struct {
//	value int
//	rowNum int
//	colNum int
//	possibilities Set
//	constraints Set
//}
//
//func (cell *Cell) addConstraint(value int) {
//	cell.constraints.insert(value)
//	cell.possibilities.remove(value)
//}
//
//func (cell *Cell) getNumPossibilities() int {
//	return cell.possibilities.getLength()
//}
//
//func (cell *Cell) addConstraints(constraints []int) {
//
//}
//
//func makeCell(value, row, col int, constraints []int) {
//	var cell Cell{}
//	cell.value = value
//	cell.colNum = col
//	cell.rowNum = row
//}

//
//import "fmt"
//
//func makeInnerSet() map[int]struct{} {
//	return make(map[int]struct{})
//}
//
//type NestedSet struct {
//	sets map[int]map[int]struct{}
//}
//
//func makeSet() NestedSet {
//	var wrapper NestedSet
//	outer := make(map[int]map[int]struct{})
//	for i := range [gridSize]int{} {
//		outer[i] = makeInnerSet()
//	}
//	wrapper.sets = outer
//	return wrapper
//}
//
//func makeSolvedSet() NestedSet {
//	set := makeSet()
//	for outer := range [gridSize]int{} {
//		for inner := range [gridSize]int{} {
//			set.insert(outer, inner)
//		}
//	}
//	return set
//}
//
//func (wrapper *NestedSet) getDifferenceSet() NestedSet {
//	differenceSet := makeSet()
//	for outer := range [gridSize]int{} {
//		for inner := range [gridSize]int{} {
//			if !wrapper.checkExists(outer, inner+1) {
//				differenceSet.insert(outer, inner+1)
//			}
//		}
//	}
//	return differenceSet
//}
//
//func (wrapper *NestedSet) getInnerDifferenceSet(outer int) map[int]struct{} {
//	set := makeInnerSet()
//	for inner := range [gridSize]int{} {
//		if !wrapper.checkExists(outer, inner+1) {
//			set[inner] = exists
//		}
//	}
//	return set
//}
//
//func (wrapper *NestedSet) checkExists(outer, inner int) bool {
//	_, ok := wrapper.sets[outer][inner]
//	return ok
//}
//
//func (wrapper *NestedSet) insert(outer, inner int) {
//	wrapper.sets[outer][inner] = exists
//}
//
//func (wrapper *NestedSet) getLength(outer int) int {
//	return len(wrapper.sets[outer])
//}
//
//func (wrapper *NestedSet) getValues(outer int) map[int]struct{} {
//	return wrapper.sets[outer]
//}
//
//func (wrapper *NestedSet) checkSolved(outer int) bool {
//	return wrapper.getLength(outer) == gridSize
//}
//
//func (wrapper *NestedSet) reset(givenNestedSet *NestedSet) {
//	wrapper.sets = givenNestedSet.copySet().sets
//}
//
//func (wrapper *NestedSet) displaySingle(outer int) {
//	for key, _ := range wrapper.sets[outer] {
//		fmt.Print(key, " ")
//	}
//}
//
//func (wrapper *NestedSet) displayAll() {
//	for key, _ := range wrapper.sets {
//		fmt.Print(key, " | ")
//		wrapper.displaySingle(key)
//		fmt.Print("\n")
//	}
//}
//
//func (wrapper *NestedSet) copySet() NestedSet {
//	set := makeSet()
//	for outerKey, outerValue := range wrapper.sets {
//		for innerKey, _ := range outerValue {
//			set.insert(outerKey, innerKey)
//		}
//	}
//	return set
//}
//
//func checkExistsInner(set map[int]struct{}, num int) bool {
//	_, exists := set[num]
//	return exists
//}
//
//func mergeDifferenceSets(row, col, block map[int]struct{}) []int {
//	nums := []int{}
//	for i := range [gridSize]int{} {
//		i = i + 1
//		if checkExistsInner(row, i) && checkExistsInner(col, i) && checkExistsInner(block, i) {
//			nums = append(nums, i)
//		}
//	}
//	return nums
//}

//type Puzzle struct {
//	cellValues  Grid
//	givenValues Grid
//
//	neededSpaces Coordinates
//
//	rowSets   NestedSet
//	colSets   NestedSet
//	blockSets NestedSet
//
//	givenRowSets   NestedSet
//	givenColSets   NestedSet
//	givenBlockSets NestedSet
//
//	matrix PossibilityMatrix
//}
//
//func makePuzzle(data [gridSize][gridSize]int) Puzzle {
//	var puzzle Puzzle
//
//	neededSpaces := makeCoordinates()
//	nestedRowSets, nestedColSets, nestedBlockSets := puzzle.makeSets()
//
//	for row := range [gridSize]int{} {
//		for col := range [gridSize]int{} {
//			value := data[row][col]
//
//			if value == 0 { // only want non-zero values in sets
//				neededSpaces.appendWrapper([2]int{row, col})
//			} else {
//				nestedRowSets.insert(row, value)
//				nestedColSets.insert(col, value)
//				nestedBlockSets.insert(calculateBlockNumber(row, col), value)
//			}
//		}
//	}
//
//	puzzle.establishGrid(makeGrid(data))
//	puzzle.neededSpaces = neededSpaces
//	puzzle.establishSets(nestedRowSets, nestedColSets, nestedBlockSets)
//
//	puzzle.matrix = puzzle.fillPossibilityMatrix()
//
//	return puzzle
//}
//
//func calculateBlockNumber(row int, col int) int {
//	return (row / 3) + (col/3)*3
//}
//
//func (puzzle *Puzzle) checkSolved() bool {
//	for i := range [gridSize]int{} {
//		incorrectRow := !puzzle.rowSets.checkSolved(i)
//		incorrectCol := !puzzle.colSets.checkSolved(i)
//		incorrectBlock := !puzzle.givenBlockSets.checkSolved(i)
//
//		if incorrectRow || incorrectCol || incorrectBlock {
//			return false
//		}
//	}
//	return true
//}
//
//func (puzzle *Puzzle) reset() {
//
//	puzzle.rowSets.reset(&puzzle.givenRowSets)
//	puzzle.colSets.reset(&puzzle.givenColSets)
//	puzzle.blockSets.reset(&puzzle.givenBlockSets)
//
//	puzzle.cellValues.reset(&puzzle.givenValues)
//}
//
//func (puzzle *Puzzle) establishSets(row, col, block NestedSet) {
//	puzzle.rowSets = row
//	puzzle.givenRowSets = row.copySet() // to avoid identical pointers
//
//	puzzle.colSets = col
//	puzzle.givenColSets = col.copySet()
//
//	puzzle.blockSets = block
//	puzzle.givenBlockSets = block.copySet()
//
//}
//
//func (puzzle *Puzzle) establishGrid(grid Grid) {
//	puzzle.cellValues = grid
//	puzzle.givenValues = grid
//}
//
//func (puzzle *Puzzle) makeSets() (row, col, block NestedSet) {
//	return makeSet(), makeSet(), makeSet()
//}

//
//import "fmt"
//
//func reportSolve(puzzle Puzzle, f func(Puzzle) Puzzle) {
//	fmt.Print("Starting Solve")
//	f(puzzle)
//	displayGrid(puzzle.cellValues)
//}

//type Coordinates struct {
//	array [][gridDimensions]int
//}
//
//func makeCoordinates() Coordinates {
//	var coordinates Coordinates
//	array := [][gridDimensions]int{}
//	coordinates.array = array
//	return coordinates
//}
//
//func (coordinates *Coordinates) appendWrapper(pair [gridDimensions]int) {
//	coordinates.array = append(coordinates.array, pair)
//}
//
//func (coordinates *Coordinates) getRowNumber(i int) int {
//	return coordinates.array[i][0]
//}
//
//func (coordinates *Coordinates) getColumnNumber(i int) int {
//	return coordinates.array[i][1]
//}

//type Coordinates struct {
//	array [][gridDimensions]int
//}
//
//func makeCoordinates() Coordinates {
//	var coordinates Coordinates
//	array := [][gridDimensions]int{}
//	coordinates.array = array
//	return coordinates
//}
//
//func (coordinates *Coordinates) appendWrapper(pair [gridDimensions]int) {
//	coordinates.array = append(coordinates.array, pair)
//}
//
//func (coordinates *Coordinates) getRowNumber(i int) int {
//	return coordinates.array[i][0]
//}
//
//func (coordinates *Coordinates) getColumnNumber(i int) int {
//	return coordinates.array[i][1]
//}
