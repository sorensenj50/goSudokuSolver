package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const gridSize = 9

func main() {
	router := gin.Default()

	router.GET("/",
		func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, infoResponse())
		},
	)

	router.GET("/info",
		func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, infoResponse())
		},
	)

	router.GET("/puzzle",
		func(c *gin.Context) {
			if queryParam, hasParam := c.GetQuery("difficulty"); hasParam {
				difficulty := difficultyMapper(queryParam)
				response := generatePuzzle(difficulty)
				response.Message = "Using specified " + queryParam + " difficulty."
				c.IndentedJSON(http.StatusOK, response)
				return
			}
			response := generatePuzzle(0.5)
			response.Message = "No difficulty specified. Using Medium."
			c.IndentedJSON(http.StatusOK, response)
			return
		},
	)

	router.Run("localhost:8080")
}

func generatePuzzle(difficulty float64) PuzzleResponse {
	var toReturn = PuzzleResponse{}
	puzzle := makePuzzle()
	ref := &puzzle
	toReturn.Solution = ref.fillGrid().values
	ref.makeGaps(difficulty)
	toReturn.Prompt = ref.values
	return toReturn
}
