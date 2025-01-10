package dbscan

import "math"

const (
	X       = 0
	Y       = 1
	CLUSTER = 2
)

// This data type holds the following values:
// x, y, category
type DataPoint [3]float64

func DBScan(dataPoints []DataPoint, maxDistance float64, minSamples int) []DataPoint {
	clusterIds := make([]float64, 0)
	var clusterId float64

	for n := 0; n < len(dataPoints); n++ {
		if dataPoints[n][CLUSTER] == 0.0 {
			continue
		}

		clusterId = float64(len(clusterIds) + 1)
		clusterIds = append(clusterIds, clusterId)

		dataPoints = paintCluster(dataPoints, n, maxDistance, minSamples, clusterId)
	}

	return dataPoints
}

func paintCluster(points []DataPoint, start int, maxDistance float64, minSamples int, clusterId float64) []DataPoint {
	stack := make([]int, len(points))
	// Prefill with -1
	for n := range stack {
		stack[n] = -1
	}
	stack[0] = start
	currentStackPos := 0
	lastStackElementPos := 0

	var closePoints []int
	var currentPointIdx int

	visited := make([]int, len(points))

	for stack[currentStackPos] > -1 || currentStackPos < len(stack) {

	CONTINUE_LABEL:
		currentPointIdx = stack[currentStackPos]
		if currentPointIdx == -1 {
			break
		}
		currentStackPos++

		// Make sure we have not visited this point before
		for n := 0; n < len(visited); n++ {
			if visited[n] == currentPointIdx {
				goto CONTINUE_LABEL
			}
		}
		visited = append(visited, currentPointIdx)

		// Holds the indexes of all close neighbouring points
		closePoints = make([]int, 0)

		for n, neighbour := range points {

			// Avoid nodes that already have a category set
			if neighbour[CLUSTER] != 0.0 {
				continue
			}

			// Avoid checking distance to self
			if n == currentPointIdx {
				continue
			}

			// Check distance to neighbour
			if distance(points[currentPointIdx], neighbour) <= maxDistance {
				// If it is close enough view it as a 'close point'
				closePoints = append(closePoints, n)
			}
		}

		if len(closePoints) >= minSamples {
			// Is core point
			if points[currentPointIdx][CLUSTER] == 0.0 {
				points[currentPointIdx][CLUSTER] = clusterId
			}
			for _, closePointIdx := range closePoints {
				if points[closePointIdx][CLUSTER] == 0.0 {
					points[closePointIdx][CLUSTER] = clusterId
				}

				stack[lastStackElementPos+1] = closePointIdx
				lastStackElementPos++
			}
		}

	}

	return points
}

// Calculates the distance between 2 points in 2d space
func distance(point1, point2 DataPoint) float64 {
	return float64(math.Abs(math.Sqrt(math.Pow(float64(point2[Y]-point1[Y]), 2) + math.Pow(float64(point2[X]-point1[X]), 2))))
}