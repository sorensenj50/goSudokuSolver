package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ArraySet struct {
	// index acts as a key for boolean mapping
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

	arraySet.shuffle()
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

func (wrapper *ArraySet) iterate(need bool, other, another *ArraySet, f func(need bool, index int, other, another *ArraySet) bool) *ArraySet {
	arraySet := makeArraySet()
	for index := range wrapper.static {
		condition := f(need, index, other, another)
		if condition {
			arraySet.set(index, need)
		} else {
			arraySet.set(index, !need)
		}
	}
	return arraySet
}

func (wrapper *ArraySet) intersection(need bool, other, another *ArraySet) *ArraySet {
	return wrapper.iterate(need, other, another, wrapper.intersectionCondition)
}

func (wrapper *ArraySet) intersectionCondition(need bool, index int, other, another *ArraySet) bool {
	return wrapper.get(index) == need && other.get(index) == need && another.get(index) == need
}

func (wrapper *ArraySet) union(need bool, other, another *ArraySet) *ArraySet {
	return wrapper.iterate(need, other, another, wrapper.unionCondition)
}

func (wrapper *ArraySet) unionCondition(need bool, index int, other, another *ArraySet) bool {
	return wrapper.get(index) == need || other.get(index) == need || another.get(index) == need
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

func (wrapper *ArraySet) shuffle() {
	for index := range wrapper.static {
		newIndex := getRandomIndex()
		wrapper.swap(index, newIndex)
	}
}

func (wrapper *ArraySet) swap(indexOne, indexTwo int) {
	temp := wrapper.shuffled[indexTwo]
	wrapper.shuffled[indexTwo] = wrapper.shuffled[indexOne]
	wrapper.shuffled[indexOne] = temp
}

func (wrapper *ArraySet) pop(adjust int) int {
	for orderIndex := range wrapper.shuffled {
		boolKey := wrapper.shuffled[adjustPopIndex(orderIndex, adjust)]
		if wrapper.get(boolKey) {
			return boolKey + 1
		}
	}
	return 0
}

func getRandomIndex() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return 0 + rand.Intn(8-0)
}

func adjustPopIndex(normal, adjustment int) int {
	return (normal + adjustment) % 9
}
