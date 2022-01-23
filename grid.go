package main

import (
	"fmt"
	"strconv"
)

type Grid struct {
	grid [gridSize][gridSize]int
}

func makeGrid(data [gridSize][gridSize]int) Grid {
	wrapper := Grid{}
	wrapper.grid = data
	return wrapper
}

func (wrapper *Grid) insert(rowNum, colNum, value int) {
	wrapper.grid[rowNum][colNum] = value
}

func (wrapper *Grid) get(rowNum, colNum int) int {
	return wrapper.grid[rowNum][colNum]
}

func (wrapper *Grid) reset(given *Grid) {
	wrapper.grid = given.grid
}

func (wrapper *Grid) display(message string) {
	string := "\n   " + message + " \n  "
	for rowNum := range [gridSize]int{} {
		for colNum := range [gridSize]int{} {
			string += " " + strconv.Itoa(wrapper.get(rowNum, colNum)) + " "

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
