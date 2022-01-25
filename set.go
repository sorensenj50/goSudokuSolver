package main

import (
	"fmt"
)

type BoolMap struct {
	bMap map[int]bool
}

func makeBoolMap() *BoolMap {
	var set BoolMap
	set.bMap = make(map[int]bool)

	for i := range [gridSize]int{} {
		set.bMap[i+1] = true
	}

	return &set
}

func (wrapper *BoolMap) toggle(key int) {
	value := wrapper.bMap[key]
	wrapper.bMap[key] = !value
}

func (wrapper *BoolMap) setTrue(key int) {
	wrapper.bMap[key] = true
}

func (wrapper *BoolMap) setFalse(key int) {
	wrapper.bMap[key] = false
}

func (wrapper *BoolMap) pop() int {
	for key, _ := range wrapper.bMap {
		return key
	}
	return 0
}

func (wrapper *BoolMap) isTrue(key int) bool {
	return wrapper.bMap[key]
}

func (wrapper *BoolMap) isFalse(key int) bool {
	return !wrapper.isTrue(key)
}

func (wrapper *BoolMap) getKeys(wantTrueValues bool) []int {
	slice := []int{}
	for key, trueValue := range wrapper.bMap {
		if wantTrueValues && trueValue {
			slice = append(slice, key)
		}

		if !wantTrueValues && !trueValue {
			slice = append(slice, key)
		}
	}
	return slice
}

func (wrapper *BoolMap) getNumKeys(wantTrueValues bool) int {
	counter := 0
	for _, trueValue := range wrapper.bMap {
		if wantTrueValues && trueValue {
			counter += 1
		}

		if !wantTrueValues && !trueValue {
			counter += 1
		}
	}
	return counter
}

func (wrapper *BoolMap) intersection(wantTrueValues bool, other, another *BoolMap) *BoolMap {
	unionSet := makeBoolMap()

	for i := range [gridSize]int{} {
		key := i + 1
		if wantTrueValues && wrapper.isTrue(key) && other.isTrue(key) && another.isTrue(key) {
			unionSet.setTrue(key)
		}

		if !wantTrueValues && !wrapper.isTrue(key) && !other.isTrue(key) && !another.isTrue(key) {
			unionSet.setFalse(key)
		}
	}
	return unionSet
}

func (wrapper *BoolMap) union(wantTrueValues bool, other, another *BoolMap) *BoolMap {
	intersectionSet := makeBoolMap()

	for i := range [gridSize]int{} {
		key := i + 1
		if wantTrueValues && wrapper.isTrue(key) || other.isTrue(key) || another.isTrue(key) {
			intersectionSet.setTrue(key)
		}

		if !wantTrueValues && !wrapper.isTrue(key) || !other.isTrue(key) || !another.isTrue(key) {
			intersectionSet.setFalse(key)
		}
	}

	return intersectionSet
}

func (wrapper *BoolMap) display() {
	for key, value := range wrapper.bMap {
		fmt.Println(key, ": ", value)
	}
}
