package main

type Coordinates struct {
	array [][gridDimensions]int
}

func makeCoordinates() Coordinates {
	var coordinates Coordinates
	array := [][gridDimensions]int{}
	coordinates.array = array
	return coordinates
}

func (coordinates *Coordinates) appendWrapper(pair [gridDimensions]int) {
	coordinates.array = append(coordinates.array, pair)
}

func (coordinates *Coordinates) getRowNumber(i int) int {
	return coordinates.array[i][0]
}

func (coordinates *Coordinates) getColumnNumber(i int) int {
	return coordinates.array[i][1]
}
