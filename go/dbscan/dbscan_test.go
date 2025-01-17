package dbscan

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
)

var debug = flag.Bool("debug", false, "Output debug information from tests")

func TestDBScan(t *testing.T) {
	rand.Seed(42)

	maxDistance := float32(4.0)
	minSamples := 3
	data, labels := getData("2d-10c")

	if data == nil || len(data) < 1 {
		t.Fatal("Expected data to have length")
	}

	points := DBScan(data, maxDistance, minSamples)

	numErrors := compareClusters(getLabelsInts(points), labels, t)

	if numErrors > 0 {
		t.Errorf("%d / %d failed", numErrors, len(points))
	}

	clusters := getLabelsInts(points)
	for idx, label := range labels {
		fmt.Println(clusters[idx], label)
	}

	writeToCsv(points)
}

func TestSimple(t *testing.T) {
	// t.Skip(*debug == false)
	rand.Seed(42)

	maxDistance := float32(5.0)
	minSamples := 3

	points := []DataPoint{
		{1.1, 1.1, -1},
		{1.2, 1.2, -1},
		{1.0, 1.0, -1},
		{0.9, 0.9, -1},
		{10.1, 10.1, -1},
	}

	result := DBScan(points, maxDistance, minSamples)

	fmt.Println("Finished")
	fmt.Println(result)
}

func TestPrintPoints(t *testing.T) {
	t.Skip(!*debug)
	rand.Seed(42)

	maxDistance := float32(5.0)
	minSamples := 3
	data, _ := getData("2d-10c")

	if len(data) < 1 {
		t.Fatal("Expected data to have length")
	}

	points := DBScan(data, maxDistance, minSamples)
	// labels := getLabels(points)

	// Debug
	fmt.Println("x,y,clusterId")
	for _, point := range points {
		fmt.Printf("%f,%f,%d\n", point[0], point[1], int(point[2]))
	}
}

func BenchmarkDBScanSimple(b *testing.B) {
	maxDistance := float32(2.0)
	minSamples := 3
	data, _ := getData("2d-10c")

	for n := 0; n < b.N; n++ {
		DBScan(data, maxDistance, minSamples)
	}
}

func TestDistance(t *testing.T) {
	point1 := [3]float32{15.7184, 62.5602, 0.0}
	point2 := [3]float32{15.2056, 63.5720, 0.0}

	dist := distance(point1, point2)

	if dist < 1.1343 || dist > 1.1344 {
		t.Errorf("Expected distance to be 1.134329... got %f", dist)
	}
}

func BenchmarkDistance(b *testing.B) {
	rand.Seed(42)

	var testCases [25][2][3]float32
	for caseIdx := range testCases {
		testCases[caseIdx] = [2][3]float32{
			{rand.Float32() * 10, rand.Float32() * 10, 0.0},
			{rand.Float32() * 10, rand.Float32() * 10, 0.0},
		}
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, testCase := range testCases {
			distance(testCase[0], testCase[1])
		}
	}
}

func getData(name string) ([]DataPoint, []int) {
	data, err := os.ReadFile("./data-" + name + ".csv")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(data), "\n")

	points := make([]DataPoint, len(rows))
	labels := make([]int, len(rows))

	x := 0.0
	y := 0.0
	cluster := -1
	for idx, row := range rows {
		if len(row) == 0 {
			continue
		}

		parts := strings.Split(row, ",")

		x, err = strconv.ParseFloat(parts[X], 32)
		if err != nil {
			panic(err)
		}

		y, err = strconv.ParseFloat(parts[Y], 32)
		if err != nil {
			panic(err)
		}

		cluster, err = strconv.Atoi(parts[CLUSTER])
		if err != nil {
			panic(err)
		}

		// if x > 30.0 && y > 50 {
		points[idx] = DataPoint{float32(x), float32(y), float32(-1)}
		labels[idx] = cluster
		//}
	}

	return points, labels
}

func getLabels(points []DataPoint) []float32 {
	labels := make([]float32, len(points))
	for n := range points {
		labels[n] = points[n][CLUSTER]
	}

	return labels
}

func getLabelsInts(points []DataPoint) []int {
	labels := make([]int, len(points))
	for n := range points {
		labels[n] = int(points[n][CLUSTER])
	}

	return labels
}

func clearLabels(points []DataPoint) []DataPoint {
	for n := range points {
		points[n][CLUSTER] = 0.0
	}

	return points
}

func getRandomTestCases(labels []int) [][2]int {
	randIdx := 0
	testCases := make([][2]int, 100)
	for idx := range testCases {
		randIdx = rand.Intn(len(labels))

		testCases[idx][0] = randIdx
		testCases[idx][1] = labels[randIdx]
	}

	return testCases
}

func findAllSimilar(points []DataPoint, labels []int, t *testing.T) int {
	var numErrors int
	conversionMap := make(map[int]int)

	for idx, point := range points {

		actualCluster := int(point[CLUSTER])
		expectedCluster := labels[idx]

		if _, ok := conversionMap[actualCluster]; !ok {
			conversionMap[actualCluster] = expectedCluster
		}

		if actualCluster != expectedCluster {
			t.Errorf("Expected cluster %d, got %d", expectedCluster, actualCluster)
			numErrors++
		}
	}

	return numErrors
}

func compareClusters(actual []int, expected []int, t *testing.T) int {
	failures := make([]int, 0)

	checked := make(map[int]bool, 0)
	for n := 0; n < len(expected); n++ {
		if val, _ := checked[expected[n]]; val {
			continue
		}

		for x := 0; x < len(expected); x++ {
			if expected[n] == expected[x] && actual[n] != actual[x] {
				failures = append(failures, x)
				break
			}
		}

		checked[n] = true
	}

	for idx := range failures {
		k := failures[idx]
		t.Errorf("Expected point %d to be cluster %d, got %d", k, expected[k], actual[k])
	}

	return len(failures)
}

func writeToCsv(points []DataPoint) {
	content := ""
	for _, point := range points {
		content += fmt.Sprintf("%.4f,%.4f,%d\n", point[X], point[Y], int(point[CLUSTER]))
	}

	err := os.WriteFile("./clusters.csv", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
