package main

import "strings"

func errorResponse() Info {
	var response = Info{}
	response.Message = "Error. If you pass a difficulty param, make sure it is a float"
	return response
}

func infoResponse() Info {
	var response = Info{}
	response.Message = "You've reached the Go Sudoku Solver API! " +
		"The /puzzle GET endpoint will return a randomly generated puzzle and its solution. " +
		"You can specify a difficulty with the \"difficulty\" query parameter--either easy, medium, or hard. "
	return response
}

func difficultyMapper(difficulty string) float64 {
	standardized := strings.ToLower(difficulty)
	if standardized == "easy" {
		return 0.3
	} else if standardized == "medium" {
		return 0.4
	} else {
		return 0.5
	}
}

type PuzzleResponse struct {
	Solution [gridSize][gridSize]int
	Prompt   [gridSize][gridSize]int
	Message  string
}

type Info struct {
	Message string
}

//var wg = sync.WaitGroup{}
//func concurrentSolve(puzzle Puzzle, numSolvers int) {
//	crossChannel := make(chan string)
//	for i := 0; i < numSolvers; i++ {
//		puzzleCopy := puzzle
//		referenceToCopy := &puzzleCopy
//		wg.Add(1)
//
//		go func() {
//			referenceToCopy.fillGrid(crossChannel)
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	fmt.Println("Here")
//}
