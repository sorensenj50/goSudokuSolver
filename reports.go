package main

import "fmt"

func reportSolve(puzzle Puzzle, f func(Puzzle) Puzzle) {
	fmt.Print("Starting Solve")
	solution := f(puzzle)
	displayGrid(solution.cellValues)
}
