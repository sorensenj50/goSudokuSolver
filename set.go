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

func (wrapper *BoolMap) set(key int, value bool) {
	wrapper.bMap[key] = value
}

func (wrapper *BoolMap) pop() int {
	for key, value := range wrapper.bMap {
		if value {
			return key
		}
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
	intersectionSet := makeBoolMap()

	for i := range [gridSize]int{} {
		key := i + 1
		if (wrapper.isTrue(key) == wantTrueValues) && (other.isTrue(key) == wantTrueValues) && (another.isTrue(key) == wantTrueValues) {
			intersectionSet.set(key, wantTrueValues)
			fmt.Println(key)
			fmt.Println("is True")
			fmt.Println("wrapper", wrapper.isTrue(key))
			fmt.Println("other", other.isTrue(key))
			fmt.Println("another", another.isTrue(key))
		} else {
			intersectionSet.set(key, !wantTrueValues)
		}
	}

	return intersectionSet
}

func (wrapper *BoolMap) union(wantTrueValues bool, other, another *BoolMap) *BoolMap {
	unionSet := makeBoolMap()

	for i := range [gridSize]int{} {
		key := i + 1

		if wrapper.isTrue(key) || other.isTrue(key) || another.isTrue(key) {
			unionSet.set(key, wantTrueValues)
		} else {
			unionSet.set(key, !wantTrueValues)
		}
	}

	return unionSet
}

func (wrapper *BoolMap) display() {
	for key, value := range wrapper.bMap {
		fmt.Println(key, ": ", value)
	}
}
