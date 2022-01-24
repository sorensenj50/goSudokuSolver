package main

type Set struct {
	set map[int]struct{}
}

func makeSet() *Set {
	var set Set
	set.set = make(map[int]struct{})
	return &set
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
