package main

import (
	"fmt"
	"strconv"
)

func calculateBlockNumber(i int, j int) int {
	return (i / 3) + (j/3)*3
}

func displayGrid(grid [9][9]int) {
	fmt.Println()
	fmt.Println("  Puzzle")
	string := "  "
	for i := range [9]int{} {
		for j := range [9]int{} {
			string += strconv.Itoa(grid[i][j])
			if j == 2 || j == 5 {
				string += "  "
				string += "|"
				string += "  "
			} else {
				string += " "
			}
		}
		string += "\n  "
		if i == 2 || i == 5 {
			string += "--------------------------"
			string += "\n  "
		}
	}
	fmt.Print(string)
}
