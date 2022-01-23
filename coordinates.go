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

func (coordinates *Coordinates) append(pair [gridDimensions]int) {
	coordinates.array = append(coordinates.array, pair)
}
