package main

import "fmt"

type Set struct {
	set map[int]struct{}
}

func makeSet() *Set {
	var set Set
	set.set = make(map[int]struct{})
	return &set
}

func makeFullSet() *Set {
	set := makeSet()
	for i := range [gridSize]int{} {
		set.insert(i + 1)
	}
	return set
}

func (set *Set) remove(value int) {
	delete(set.set, value)
}

func (set *Set) insert(value int) {
	set.set[value] = exists
}

func (set *Set) doesContain(value int) bool {
	_, exists := set.set[value]
	return exists
}

func (set *Set) getLength() int {
	return len(set.set)
}

func (set *Set) intersection(other, another *Set) *Set {
	unionSet := makeSet()

	for key, _ := range set.set {
		if other.doesContain(key) && another.doesContain(key) {
			unionSet.insert(key)
		}
	}

	for key, _ := range other.set {
		if set.doesContain(key) && another.doesContain(key) {
			unionSet.insert(key)
		}
	}

	for key, _ := range another.set {
		if set.doesContain(key) && other.doesContain(key) {
			unionSet.insert(key)
		}
	}

	return unionSet
}

func (set *Set) union(other, another *Set) *Set {
	unionSet := makeSet()

	for key, _ := range set.set {
		unionSet.insert(key)
	}

	for key, _ := range other.set {
		unionSet.insert(key)
	}

	for key, _ := range another.set {
		unionSet.insert(key)
	}

	return unionSet
}

func (set *Set) pop() int {
	for key, _ := range set.set {
		return key
	}
	return 0
}

func (set *Set) display() {
	fmt.Println("Set")
	for key, _ := range set.set {
		fmt.Print(key)
	}
	fmt.Println()
}
