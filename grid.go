package main

//
//import (
//	"fmt"
//	"strconv"
//)
//
//type Grid struct {
//	grid [gridSize][gridSize]int
//}
//
//func makeGrid() Grid {
//	wrapper := Grid{}
//	return wrapper
//}
//
//func (wrapper *Grid) setGrid(data [gridSize][gridSize]int) {
//	wrapper.grid = data
//}
//
//func (wrapper *Grid) insert(rowNum, colNum, value int) {
//	wrapper.grid[rowNum][colNum] = value
//}
//
//func (wrapper *Grid) remove(rowNum, colNum int) {
//	wrapper.grid[rowNum][colNum] = 0
//}
//
//func (wrapper *Grid) get(rowNum, colNum int) int {
//	return wrapper.grid[rowNum][colNum]
//}
//
//func (wrapper *Grid) reset(given *Grid) {
//	wrapper.grid = given.grid
//}
//
//func (wrapper *Grid) iterate(f func(row, col int), ignoreZeros bool) {
//	for rowNum := range [gridSize]int{} {
//		for colNum := range [gridSize]int{} {
//			if wrapper.get(rowNum, colNum) == 0 && ignoreZeros {
//
//			} else {
//				f(rowNum, colNum)
//			}
//		}
//	}
//}
//
//func (wrapper *Grid) display() {
//	string := "\n  "
//	for rowNum := range [gridSize]int{} {
//		for colNum := range [gridSize]int{} {
//			string += " " + strconv.Itoa(wrapper.get(rowNum, colNum)) + " "
//
//			if colNum == 2 || colNum == 5 {
//				string += " | "
//			}
//
//			if colNum == 8 {
//				string += "\n  "
//			}
//		}
//		if rowNum == 2 || rowNum == 5 {
//			string += " ------------------------------- \n  "
//		}
//	}
//	fmt.Print(string)
//}
//
//func moveIndicesForward(row, col int) (int, int) {
//	if col == 8 && row == 8 {
//		return 8, 8
//	} else if col == 8 {
//		return row + 1, 0
//	} else {
//		return row, col + 1
//	}
//}
//
//func moveIndicesBackward(row, col int) (int, int) {
//	if col == 0 && row == 0 {
//		return 0, 0
//	} else if col == 0 {
//		return row - 1, 8
//	} else {
//		return row, col - 1
//	}
//}
