package main

import (
	"math/rand"
	"time"
)

type NewPuzzle struct {
	values    [gridSize][gridSize]int
	generator RandomGenerator
}

func makeNewPuzzle() *NewPuzzle {
	newPuzzle := NewPuzzle{}
	return &newPuzzle
}

func (wrapper *NewPuzzle) addValues() {
	row := 0
	col := 0
	shouldContinue := true

	for shouldContinue {
		num := wrapper.generator.get()
		if wrapper.isValid(num, row, col) {
			wrapper.values[row][col] = num
			wrapper.generator.reset()
		} else {
			wrapper.generator.setInvalid(num)
		}
	}
}

func (wrapper *NewPuzzle) isValidRow(value, row int) bool {
	for index := range [gridSize]int{} {
		num := wrapper.values[row][index]
		if num == value {
			return false
		}
	}
	return true
}

func (wrapper *NewPuzzle) isValidCol(value, col int) bool {
	for index := range [gridSize]int{} {
		num := wrapper.values[index][col]
		if num == value {
			return false
		}
	}
	return true
}

func (wrapper *NewPuzzle) isValidBlock(value, row, col int) bool {
	incrementArray := [...]int{-1, 0, 1}
	centerRow, centerCol := getBlockCenterCoordinates(row, col)

	for _, vertical := range incrementArray {
		for horizontal := range incrementArray {
			if wrapper.values[centerRow+vertical][centerCol+horizontal] == value {
				return false
			}
		}
	}
	return true
}

func (wrapper *NewPuzzle) isValid(value, row, col int) bool {
	isValidRow := wrapper.isValidRow(value, row)
	isValidCol := wrapper.isValidCol(value, col)
	isValidBlock := wrapper.isValidBlock(value, row, col)
	return isValidRow && isValidCol && isValidBlock
}

func getBlockCenterCoordinates(row, col int) (int, int) {
	if row < 3 && col < 3 {
		return 1, 1
	} else if row < 3 && col < 6 {
		return 1, 4
	} else if row < 3 {
		return 1, 7
	} else if row < 6 && col < 3 {
		return 4, 1
	} else if row < 6 && col < 6 {
		return 4, 4
	} else if row < 6 {
		return 4, 7
	} else if col < 3 {
		return 7, 1
	} else if col < 6 {
		return 7, 4
	} else {
		return 7, 7
	}
}

type RandomGenerator struct {
	possibilities [gridSize]bool
}

func (wrapper *RandomGenerator) get() int {
	num := randomNumber(0, len(wrapper.possibilities))
	for range [gridSize]int{} {
		if wrapper.possibilities[num] {
			return num
		} else {
			num = (num + 1) % 9
		}
	}
	return 0
}

func (wrapper *RandomGenerator) setInvalid(num int) {
	wrapper.possibilities[num] = false
}

func (wrapper *RandomGenerator) reset() {
	trueValues := [gridSize]bool{true, true, true, true, true, true, true, true, true}
	wrapper.possibilities = trueValues
}

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
