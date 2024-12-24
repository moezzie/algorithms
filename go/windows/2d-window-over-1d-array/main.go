package main

import "fmt"

func main() {
	gridWidth := 3
	//gridHeight := 3

	windowX1 := 1
	windowY1 := 0
	windowX2 := 2
	windowY2 := 1

	window := getWindow(data(), gridWidth, windowX1, windowY1, windowX2, windowY2)
	fmt.Println(window)
}

// Get the 2d window
func getWindow(data []int, dataWidth, x1, y1, x2, y2 int) [][]int {
	width := x2 - x1 + 1
	height := y2 - y1 + 1

	window := new2d(width, height)

	for x := 0; x < width; x++ {
		window[x] = data[x1+(x*dataWidth) : (x1+(x*dataWidth))+height]
	}

	return window
}

// Create a new 2d array of size width x height
func new2d(width, height int) [][]int {
	window := make([][]int, width)
	for i, _ := range window {
		window[i] = make([]int, height)
	}

	return window
}

// Returns the data as a 1d array
func data() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
}
