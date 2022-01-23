package main

import "fmt"

func makeInnerSet() map[int]struct{} {
	return make(map[int]struct{})
}

type NestedSet struct {
	sets map[int]map[int]struct{}
}

func makeSet() NestedSet {
	var wrapper NestedSet
	outer := make(map[int]map[int]struct{})
	for i := range [gridSize]int{} {
		outer[i] = makeInnerSet()
	}
	wrapper.sets = outer
	return wrapper
}

func makeSolvedSet() NestedSet {
	set := makeSet()
	for outer := range [gridSize]int{} {
		for inner := range [gridSize]int{} {
			set.insert(outer, inner)
		}
	}
	return set
}

func (wrapper *NestedSet) getDifferenceSet() NestedSet {
	differenceSet := makeSet()
	for outer := range [gridSize]int{} {
		for inner := range [gridSize]int{} {
			if !wrapper.checkExists(outer, inner+1) {
				differenceSet.insert(outer, inner+1)
			}
		}
	}
	return differenceSet
}

func (wrapper *NestedSet) getInnerDifferenceSet(outer int) map[int]struct{} {
	set := makeInnerSet()
	for inner := range [gridSize]int{} {
		if !wrapper.checkExists(outer, inner+1) {
			set[inner] = exists
		}
	}
	return set
}

func (wrapper *NestedSet) checkExists(outer, inner int) bool {
	_, ok := wrapper.sets[outer][inner]
	return ok
}

func (wrapper *NestedSet) insert(outer, inner int) {
	wrapper.sets[outer][inner] = exists
}

func (wrapper *NestedSet) getLength(outer int) int {
	return len(wrapper.sets[outer])
}

func (wrapper *NestedSet) getValues(outer int) map[int]struct{} {
	return wrapper.sets[outer]
}

func (wrapper *NestedSet) checkSolved(outer int) bool {
	return wrapper.getLength(outer) == gridSize
}

func (wrapper *NestedSet) reset(givenNestedSet *NestedSet) {
	wrapper.sets = givenNestedSet.copySet().sets
}

func (wrapper *NestedSet) displaySingle(outer int) {
	for key, _ := range wrapper.sets[outer] {
		fmt.Print(key, " ")
	}
}

func (wrapper *NestedSet) displayAll() {
	for key, _ := range wrapper.sets {
		fmt.Print(key, " | ")
		wrapper.displaySingle(key)
		fmt.Print("\n")
	}
}

func (wrapper *NestedSet) copySet() NestedSet {
	set := makeSet()
	for outerKey, outerValue := range wrapper.sets {
		for innerKey, _ := range outerValue {
			set.insert(outerKey, innerKey)
		}
	}
	return set
}

func checkExistsInner(set map[int]struct{}, num int) bool {
	_, exists := set[num]
	return exists
}

func mergeDifferenceSets(row, col, block map[int]struct{}) []int {
	nums := []int{}
	for i := range [gridSize]int{} {
		i = i + 1
		if checkExistsInner(row, i) && checkExistsInner(col, i) && checkExistsInner(block, i) {
			nums = append(nums, i)
		}
	}
	return nums
}
