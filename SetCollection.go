package main

import "fmt"

type SetCollection struct {
	sets map[int]*DualSet
	kind string
}

func makeSets(kind string) SetCollection {
	var wrapper SetCollection
	outer := make(map[int]*DualSet)
	for i := range [gridSize]int{} {
		outer[i] = makeDualSet()
	}
	wrapper.sets = outer
	wrapper.kind = kind
	return wrapper
}

func (collection *SetCollection) constraintExists(outer, inner int) bool {
	return collection.sets[outer].constraints.doesContain(inner)
}

func (collection *SetCollection) addConstraint(outer, inner int) {
	collection.sets[outer].addConstraint(inner)
}

func (collection *SetCollection) getPossibilities(outer int) *Set {
	return collection.sets[outer].possibilities
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

type DualSet struct {
	constraints   *Set
	possibilities *Set
}

func (dual *DualSet) addConstraint(value int) {
	dual.constraints.insert(value)
	dual.possibilities.remove(value)
}

func (dual *DualSet) getNumPossibilities() int {
	return dual.possibilities.getLength()
}

func makeDualSet() *DualSet {
	var dualSet DualSet
	dualSet.constraints = makeSet()
	dualSet.possibilities = makeFullSet()
	return &dualSet
}
