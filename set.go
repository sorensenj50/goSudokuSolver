package main

import (
	"fmt"
	"math/rand"
	"time"
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
	if wrapper.isAllFalse() {
		return 0
	} else {
		for true {
			randomNumber := numsArray[getRandomIndex()]
			if wrapper.isTrue(randomNumber) {
				return randomNumber
			}
		}
	}

	return 1
}

func (wrapper *BoolMap) isAllFalse() bool {
	for num := range wrapper.bMap {
		if wrapper.isTrue(num) {
			return false
		}
	}
	return true
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

type ArraySet struct {
	// maps numbers to booleans
	static [9]bool

	// maps ordering (which can be shuffled) to numbers
	shuffled [9]int
}

func makeArraySet() *ArraySet {
	var arraySet ArraySet

	for index := range arraySet.static {
		arraySet.shuffled[index] = index
		arraySet.static[index] = true
	}
	return &arraySet
}

func (wrapper *ArraySet) set(num int, value bool) {
	wrapper.static[num] = value
}

func (wrapper *ArraySet) get(num int) bool {
	return wrapper.static[num]
}

func (wrapper *ArraySet) toggle(num int) {
	value := wrapper.get(num)
	wrapper.set(num, !value)
}

func (wrapper *ArraySet) getNumKeys(need bool) int {
	counter := 0
	for index := range wrapper.static {
		if wrapper.get(index) == need {
			counter += 1
		}
	}
	return counter
}

func (wrapper *ArraySet) getKeys(need bool) []int {
	arraySlice := []int{}
	for index := range wrapper.static {
		if wrapper.get(index) == need {
			arraySlice = append(arraySlice, index)
		}
	}
	return arraySlice
}

func (wrapper *ArraySet) union(need bool, other, another *ArraySet) []int {
	arraySlice := []int{}
	for index := range wrapper.static {
		condition := wrapper.get(index) == need && other.get(index) == need && another.get(index) == need
		if condition {
			arraySlice = append(arraySlice, index)
		}
	}
	return arraySlice
}

func (wrapper *ArraySet) intersection(need bool, other, another *ArraySet) []int {
	arraySlice := []int{}
	for index := range wrapper.static {
		condition := wrapper.get(index) == need || other.get(index) == need || another.get(index) == need
		if condition {
			arraySlice = append(arraySlice, index)
		}
	}
	return arraySlice
}

func (wrapper *ArraySet) display() {
	fmt.Println("Static")
	for index := range wrapper.static {
		fmt.Println(index, "(", index+1, ")", ": ", wrapper.get(index))
	}
	fmt.Println("")
	fmt.Println("Shuffled")
	for index := range wrapper.shuffled {
		fmt.Println(index, ": ", wrapper.shuffled[index])
	}
}

func (wrapper *ArraySet) swap(indexOne, indexTwo int) {
	temp := wrapper.shuffled[indexTwo]
	wrapper.shuffled[indexTwo] = wrapper.shuffled[indexOne]
	wrapper.shuffled[indexOne] = temp
}

func (wrapper *ArraySet) shuffle() {
	for index := range wrapper.static {
		newIndex := getRandomIndex()
		wrapper.swap(index, newIndex)
	}
}

func getRandomIndex() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return 0 + rand.Intn(8-0)
}
