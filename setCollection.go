package main

import (
	"fmt"
	"strconv"
)

type SetCollection struct {
	sets map[int]*ArraySet
	kind string
}

func makeSets(kind string) SetCollection {
	var wrapper SetCollection
	outer := make(map[int]*ArraySet)
	for i := range [gridSize]int{} {
		outer[i] = makeArraySet()
	}
	wrapper.sets = outer
	wrapper.kind = kind
	return wrapper
}

func (collection *SetCollection) constraintExists(outer, inner int) bool {
	return !collection.sets[outer].get(inner)
}

func (collection *SetCollection) addConstraint(outer, inner int) {
	collection.sets[outer].set(inner, false)
}

func (collection *SetCollection) removeConstraint(outer, inner int) {
	collection.sets[outer].set(inner, true)
}

func (collection *SetCollection) getPossibilities(outer int) []int {
	return collection.sets[outer].getKeys(true)
}

func (collection *SetCollection) getNumConstraints(outer int) int {
	return collection.sets[outer].getNumKeys(false)
}

func (collection *SetCollection) displayStaticTop() {
	fmt.Println()
	fmt.Println("Kind Along Y-Axis, Number Along X-Axis")
	fmt.Println("Kind =", collection.kind)
	fmt.Println()
	fmt.Print("      1 2 3 4 5 6 7 8 9\n")
	fmt.Print("      -----------------\n  ")
}

func (collection *SetCollection) displayAll() {
	collection.displayStaticTop()
	for i := range [gridSize]int{} {
		collection.displayOneHelper(i)
		fmt.Print("\n  ")
	}
}

func (collection *SetCollection) displayOne(value int) {
	collection.displayStaticTop()
	collection.displayOneHelper(value)

}
func (collection *SetCollection) displayOneHelper(value int) {
	for j := range [gridSize]int{} {
		if j == 0 {
			fmt.Print(value, " | ")
		}

		if collection.constraintExists(value, j) {
			fmt.Print("C", " ")
		} else {
			fmt.Print("P", " ")
		}
	}
}

func calculateBlockNumber(row int, col int) int {
	return (row / 3) + (col/3)*3
}

type GridSetCollection struct {
	grid map[string]*ArraySet
}

func makeGridSet() *GridSetCollection {
	var grid GridSetCollection
	grid.grid = make(map[string]*ArraySet)
	for rowNum := range [gridSize]int{} {
		for colNum := range [gridSize]int{} {
			set := makeArraySet()
			grid.grid[getStringRowCol(rowNum, colNum)] = set
		}
	}
	return &grid
}

func (wrapper *GridSetCollection) get(row, col int) *ArraySet {
	return wrapper.grid[getStringRowCol(row, col)]
}

func (wrapper *GridSetCollection) addConstraint(row, col, num int) {
	wrapper.grid[getStringRowCol(row, col)].set(num, false)
}

func (wrapper *GridSetCollection) reset(row, col int) {
	wrapper.displaySelect(row, col)
	wrapper = makeGridSet()
	wrapper.displaySelect(row, col)
}

func (wrapper *GridSetCollection) display() {
	for rowNum := range [gridSize]int{} {
		for colNum := range [gridSize]int{} {
			fmt.Print(rowNum, colNum)
			set := wrapper.get(rowNum, colNum)
			set.display()
			fmt.Println("")
		}
	}
}

func (wrapper *GridSetCollection) displaySelect(row, col int) {
	wrapper.get(row, col).displayStatic()
}

func getStringRowCol(row, col int) string {
	return strconv.Itoa(row) + strconv.Itoa(col)
}
