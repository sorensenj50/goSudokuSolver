package main

import (
	"sync"
)

const gridSize = 9

var wg = sync.WaitGroup{}

func main() {
	sharedBasePuzzle := makePuzzle()
	concurrentSolve(sharedBasePuzzle, 2)

}

func concurrentSolve(puzzle Puzzle, numSolvers int) {
	channel := make(chan string)
	for i := 0; i < numSolvers; i++ {
		puzzleCopy := puzzle
		referenceToCopy := &puzzleCopy
		wg.Add(1)
		go solveThenDisplay(referenceToCopy, channel, i)
	}
	wg.Wait()
}

func solveThenDisplay(puzzle *Puzzle, channel chan string, id int) {
	puzzle.fillGrid(channel, id)
	wg.Done()
}

// hello
