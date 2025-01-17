package dbscan

import (
	"fmt"
	"math"
)

const (
	CLUSTER_NOT_SET = -1.0
	X               = 0
	Y               = 1
	CLUSTER         = 2
)

// This data type holds the following values:
// x, y, category
type DataPoint [3]float32

func DBScan(dataPoints []DataPoint, maxDistance float32, minSamples int) []DataPoint {
	clusterIds := make([]float32, 0)
	var clusterId float32 = 0.0

	for n := 0; n < len(dataPoints); n++ {
		if dataPoints[n][CLUSTER] > CLUSTER_NOT_SET {
			continue
		}

		clusterId = float32(len(clusterIds) + 1)
		clusterIds = append(clusterIds, clusterId)

		dataPoints = paintCluster(dataPoints, n, maxDistance, minSamples, clusterId)

	}

	return dataPoints
}

func paintCluster(points []DataPoint, start int, maxDistance float32, minSamples int, clusterId float32) []DataPoint {
	stack := make([]int, len(points))
	stack[0] = start
	currentStackPos := 0
	nextStackElementPos := 1

	var closePoints []int
	var currentPointIdx int

	visited := make([]int, len(points))
	visitedInsertPos := 0

	for currentStackPos < nextStackElementPos {

	CONTINUE_LABEL:
		currentPointIdx = stack[currentStackPos]
		currentStackPos++

		// Make sure we have not visited this point before
		for n := 0; n < visitedInsertPos; n++ {
			if visited[n] == currentPointIdx {
				goto CONTINUE_LABEL
			}
		}
		visited[visitedInsertPos] = currentPointIdx
		visitedInsertPos++

		// Holds the indexes of all close neighbouring points
		closePoints = make([]int, 0)

		for n, neighbour := range points {

			// Avoid nodes that already have a category set
			// Avoid checking distance to self
			if neighbour[CLUSTER] > CLUSTER_NOT_SET || n == currentPointIdx {
				continue
			}

			dist := distance(points[currentPointIdx], neighbour)

			// Check distance to neighbour
			if dist < maxDistance {
				// If it is close enough view it as a 'close point'
				closePoints = append(closePoints, n)
			}
		}

		// Is core point
		if len(closePoints) >= minSamples {
			fmt.Println("Core point")
			// Set the current point clusterId
			if int8(points[currentPointIdx][CLUSTER]) == int8(CLUSTER_NOT_SET) {
				points[currentPointIdx][CLUSTER] = clusterId
			}

			// Set close points clusterId
			for _, closePointIdx := range closePoints {
				if int8(points[closePointIdx][CLUSTER]) == int8(CLUSTER_NOT_SET) {
					points[closePointIdx][CLUSTER] = clusterId
				}

				stack[nextStackElementPos] = closePointIdx
				nextStackElementPos++
			}
		}

	}

	return points
}

// Calculates the distance between 2 points in 2d space
func distance(point1, point2 DataPoint) float32 {
	return float32(math.Abs(
		math.Sqrt(
			float64(((point2[Y] - point1[Y]) * (point2[Y] - point1[Y])) + ((point2[X] - point1[X]) * (point2[X] - point1[X]))),
		)))
}
