package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type NewPuzzle struct {
	values    [gridSize][gridSize]int
	generator RandomGenerator
}

func makeNewPuzzle() *NewPuzzle {
	newPuzzle := NewPuzzle{}
	newPuzzle.generator = makeGenerator()
	return &newPuzzle
}

func (wrapper *NewPuzzle) reset() {
	wrapper.values = [gridSize][gridSize]int{}
	wrapper.generator.reset()
}

func (wrapper *NewPuzzle) addValues() {
	row := 0
	col := 0
	numIterations := 0

	for {
		numIterations += 1
		num := wrapper.generator.get()

		if numIterations >= 2500 {
			row, col, numIterations = 0, 0, 0
			wrapper.reset()
		} else if num == 0 {
			row, col = moveIndicesBackward(row, col)
			wrapper.generator.reset()
			wrapper.generator.setInvalid(wrapper.values[row][col] - 1)
			wrapper.values[row][col] = 0

		} else if wrapper.isValid(num, row, col) {
			wrapper.values[row][col] = num
			if row == 8 && col == 8 {
				wrapper.display()
				break
			} else {
				wrapper.generator.reset()
				row, col = moveIndicesForward(row, col)
			}
		} else {
			wrapper.generator.setInvalid(num - 1)
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
		for _, horizontal := range incrementArray {
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

func (wrapper *NewPuzzle) display() {
	string := "\n  "
	for rowNum := range [gridSize]int{} {
		for colNum := range [gridSize]int{} {
			string += " " + strconv.Itoa(wrapper.values[rowNum][colNum]) + " "

			if colNum == 2 || colNum == 5 {
				string += " | "
			}

			if colNum == 8 {
				string += "\n  "
			}
		}
		if rowNum == 2 || rowNum == 5 {
			string += " ------------------------------- \n  "
		}
	}
	fmt.Print(string)
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

func moveIndicesForward(row, col int) (int, int) {
	if col == 8 && row == 8 {
		return 8, 8
	} else if col == 8 {
		return row + 1, 0
	} else {
		return row, col + 1
	}
}

func moveIndicesBackward(row, col int) (int, int) {
	if col == 0 && row == 0 {
		return 0, 0
	} else if col == 0 {
		return row - 1, 8
	} else {
		return row, col - 1
	}
}

type RandomGenerator struct {
	possibilities [gridSize]bool
}

func makeGenerator() RandomGenerator {
	var generator RandomGenerator
	possibilities := [9]bool{}
	for index := range [gridSize]bool{} {
		possibilities[index] = true
	}
	generator.possibilities = possibilities
	return generator
}

func (wrapper *RandomGenerator) get() int {
	num := randomNumber(0, len(wrapper.possibilities))
	for range [gridSize]int{} {
		if wrapper.possibilities[num] {
			return num + 1
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
