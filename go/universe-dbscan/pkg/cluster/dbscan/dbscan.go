package dbscan

import (
	"fmt"
	"image"
	"unsafe"
)

const (
	UNDEFINED = -10
	NOISE     = -1
	X         = 0
	Y         = 1

	INT_SIZE = (int(unsafe.Sizeof(int(1)))*8 - 1)
)

func DBScan(img image.Image, eps int, minSamples int, minLuminecense int) [][]int {
	C := 0

	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	labels := make([][]int, width)
	for idxX, _ := range labels {
		labels[idxX] = make([]int, height)
		for idxY, _ := range labels[idxX] {
			labels[idxX][idxY] = UNDEFINED
		}
	}

	fmt.Println("Convertint")
	gray := convertToBrightnesMap(img, minLuminecense)

	fmt.Println("Clustering")
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if gray[x][y] < uint8(minLuminecense) {
				continue
			}

			// If this pixel
			if labels[x][y] != UNDEFINED {
				continue
			}

			N, numN := rangeQuery(gray, labels, x, y, eps, uint8(minLuminecense), minSamples)
			if numN < minSamples {
				labels[x][y] = NOISE
				continue
			}

			C = C + 1
			labels[x][y] = C

			S := N
			idxS := -1
			for len(S) > 0 {
				idxS++
				if idxS >= len(S) {
					break
				}

				Q := S[idxS]

				if gray[Q[X]][Q[Y]] < uint8(minLuminecense) {
					continue
				}

				if labels[Q[X]][Q[Y]] == NOISE {
					labels[Q[X]][Q[Y]] = C
				}
				if labels[Q[X]][Q[Y]] != UNDEFINED {
					continue
				}

				labels[Q[X]][Q[Y]] = C
				// Get all the neighbors
				N, xnumN := rangeQuery(gray, labels, Q[X], Q[Y], eps, uint8(minLuminecense), minSamples)
				if xnumN >= minSamples {
					S = append(S, N...)
				}
			}
		}
	}

	return labels
}

func rangeQuery(img [][]uint8, labels [][]int, x int, y int, maxDistance int, minLum uint8, minSamples int) ([][]int, int) {
	width := len(img)
	height := len(img[0])

	if img[x][y] == uint8(minLum) {
		return make([][]int, 0), 0
	}

	xStart := max(0, x-maxDistance-1)
	xEnd := min(width, x+maxDistance+1)
	yStart := max(0, y-maxDistance-1)
	yEnd := min(height, y+maxDistance+1)

	// Pass 1: Get the number of neighbors to determine how
	// many neighbors the point has
	numNeighbors := 0
	for pX := xStart; pX < xEnd; pX++ {
		for pY := yStart; pY < yEnd; pY++ {

			if labels[pX][pY] > NOISE || img[pX][pY] < minLum {
				continue
			}

			if distance(x, y, pX, pY) <= maxDistance {
				numNeighbors += 1
			}
		}
	}

	if numNeighbors < minSamples {
		return [][]int{}, 0
	}

	// Pass 2: Add all the neighbors to an array
	// Avoids many expensive array/slice grows
	neighbors := make([][]int, numNeighbors)
	numNeighbors = 0
	for pX := xStart; pX < xEnd; pX++ {
		for pY := yStart; pY < yEnd; pY++ {

			if labels[pX][pY] > NOISE || img[pX][pY] < minLum {
				continue
			}

			if distance(x, y, pX, pY) <= maxDistance {
				//neighbors = append(neighbors, []int{pX, pY})
				neighbors[numNeighbors] = []int{pX, pY}
				numNeighbors += 1
			}
		}
	}

	return neighbors, numNeighbors
}

func distance(x1, y1, x2, y2 int) int {
	//return max(max(x2, x1)-min(x2, x1), max(y2, y1)-min(y2, y1))
	dist := ((x2 - x1) - (y2 - y1)) * ((x2 - x1) - (y2 - y1))
	mask := dist >> INT_SIZE
	return (dist ^ mask) - mask
}

func convertToBrightnesMap(img image.Image, minLum int) [][]uint8 {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	brightnes := make([][]uint8, width)
	for x, _ := range brightnes {
		brightnes[x] = make([]uint8, height)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			brightnes[x][y] = uint8((0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 256.0)
			if brightnes[x][y] < uint8(minLum) {
				brightnes[x][y] = 0
			}
		}
	}

	return brightnes
}
