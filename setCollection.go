package main

import (
	"fmt"
)

type SetCollection struct {
	sets map[int]*BoolMap
	kind string
}

func makeSets(kind string) SetCollection {
	var wrapper SetCollection
	outer := make(map[int]*BoolMap)
	for i := range [gridSize]int{} {
		outer[i] = makeBoolMap()
	}
	wrapper.sets = outer
	wrapper.kind = kind
	return wrapper
}

func (collection *SetCollection) constraintExists(outer, inner int) bool {
	return collection.sets[outer].isFalse(inner)
}

func (collection *SetCollection) addConstraint(outer, inner int) {
	collection.sets[outer].set(inner, false)
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

		if collection.constraintExists(value, j+1) { // because 0 is empty value
			fmt.Print("C", " ")
		} else {
			fmt.Print("P", " ")
		}
	}
}

func calculateBlockNumber(row int, col int) int {
	return (row / 3) + (col/3)*3
}
