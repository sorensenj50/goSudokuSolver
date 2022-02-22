package main

import (
	"fmt"
	"strconv"
)

type CellTracker struct {
	tracker map[string]bool
}

func makeTracker() CellTracker {
	var tracker CellTracker
	tracker.tracker = make(map[string]bool)
	return tracker
}

func (wrapper *CellTracker) set(row, col int, to bool) {
	id := strconv.Itoa(row) + strconv.Itoa(col)
	wrapper.tracker[id] = to
}

func (wrapper *CellTracker) cantPlace(row, col int) bool {
	value, exists := wrapper.tracker[getKey(row, col)]
	if exists {
		return true
	} else {
		return value
	}
}

func (wrapper *CellTracker) display() {
	for row := range [gridSize]int{} {
		for col := range [gridSize]int{} {
			fmt.Print(wrapper.cantPlace(row, col), " ")

		}
		fmt.Println("")
	}
}

func getKey(numOne, numTwo int) string {
	return strconv.Itoa(numOne) + strconv.Itoa(numTwo)
}
