package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Puzzle struct {
	values    [gridSize][gridSize]int
	generator RandomGenerator
	tracker   CellTracker
}

func makePuzzle() *Puzzle {
	newPuzzle := Puzzle{}
	newPuzzle.generator = makeGenerator()
	newPuzzle.tracker = makeTracker()
	return &newPuzzle
}

func (wrapper *Puzzle) reset() {
	for row := range [gridSize]int{} {
		for col := range [gridSize]int{} {
			if wrapper.tracker.canPlace(row, col) {
				wrapper.values[row][col] = 0
			}
		}
	}

	wrapper.generator.reset()
}

func (wrapper *Puzzle) markGiven() {
	for row := range [gridSize]int{} {
		for col := range [gridSize]int{} {
			num := wrapper.values[row][col]
			wrapper.tracker.set(row, col, num == 0)
		}
	}
}

func (wrapper *Puzzle) fillGrid() {
	row := 0
	col := 0
	numIterations := 0

	for {
		numIterations += 1
		num := wrapper.generator.get()

		// can't do anything because cell is given
		if !wrapper.tracker.canPlace(row, col) {
			row, col = wrapper.findNextGap(row, col)

			// resetting to get out stuck situation
		} else if numIterations >= 2500 {
			row, col, numIterations = 0, 0, 0
			wrapper.reset()

			// can't place, needs to go backwards
		} else if num == 0 {
			row, col = wrapper.findPreviousGap(row, col)
			wrapper.generator.reset()
			wrapper.generator.setInvalid(wrapper.values[row][col] - 1)
			wrapper.values[row][col] = 0

			// placing num
		} else if wrapper.isValid(num, row, col) {
			wrapper.values[row][col] = num
			if row == 8 && col == 8 {
				wrapper.generator.reset()
				wrapper.display()
				break
			} else {
				wrapper.generator.reset()
				row, col = wrapper.findNextGap(row, col)
			}

			// num can't work, tuning generator
		} else {
			wrapper.generator.setInvalid(num - 1)
		}
	}
}

func (wrapper *Puzzle) isValidRow(value, row int) bool {
	for index := range [gridSize]int{} {
		num := wrapper.values[row][index]
		if num == value {
			return false
		}
	}
	return true
}

func (wrapper *Puzzle) isValidCol(value, col int) bool {
	for index := range [gridSize]int{} {
		num := wrapper.values[index][col]
		if num == value {
			return false
		}
	}
	return true
}

func (wrapper *Puzzle) isValidBlock(value, row, col int) bool {
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

func (wrapper *Puzzle) isValid(value, row, col int) bool {
	isValidRow := wrapper.isValidRow(value, row)
	isValidCol := wrapper.isValidCol(value, col)
	isValidBlock := wrapper.isValidBlock(value, row, col)
	return isValidRow && isValidCol && isValidBlock
}

func (wrapper *Puzzle) display() {
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

func (wrapper *Puzzle) makeGaps(probability float32) {
	for row := range [gridSize]int{} {
		for col := range [gridSize]int{} {
			randomFloat := rand.Float32()
			if randomFloat > probability {
				wrapper.values[row][col] = 0
			}
		}
	}
}

func (wrapper *Puzzle) findNextGap(row, col int) (int, int) {
	row, col = moveIndicesForward(row, col)
	for {
		if wrapper.tracker.canPlace(row, col) {
			return row, col
		} else if row >= 8 && col >= 8 {
			return 8, 8
		} else {
			row, col = moveIndicesForward(row, col)
		}
	}
}

func (wrapper *Puzzle) findPreviousGap(row, col int) (int, int) {
	row, col = moveIndicesBackward(row, col)
	for {
		if wrapper.tracker.canPlace(row, col) {
			return row, col
		} else if row <= 0 && col <= 0 {
			return 0, 0
		} else {
			row, col = moveIndicesBackward(row, col)
		}
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
