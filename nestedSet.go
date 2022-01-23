package main

type NestedSet struct {
	sets map[int]map[int]struct{}
}

func makeSet() NestedSet {
	var wrapper NestedSet
	outer := make(map[int]map[int]struct{})
	for i := range [gridSize]int{} {
		outer[i] = make(map[int]struct{})
	}
	wrapper.sets = outer
	return wrapper
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
	wrapper.sets = givenNestedSet.sets
}
