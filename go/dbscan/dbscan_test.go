package dbscan

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDBScan(t *testing.T) {
	maxDistance := float64(2.0)
	minSamples := 3
	data, err := getData("2d-10c")
	if err != nil {
		t.Fatal(err)
	}

	if len(data) < 1 {
		t.Fatal("Expected data to have length")
	}

	points := DBScan(data, maxDistance, minSamples)

	testCases := [][]int{
		{0, 0},
		{957, 2},
		{1999, 8},
		{2260, 8},
		{2989, 9},
		{2990, 0},
	}

	for _, testCase := range testCases {
		if int(points[testCase[0]][2]) != testCase[1] {
			t.Errorf("Expected index %d to be cluster %d, got %d", testCase[0], testCase[1], int(points[testCase[0]][2]))
		}
	}

	// Debug
	/*fmt.Println("x,y,clusterId")
	for _, point := range points {
		fmt.Printf("%f,%f,%d\n", point[0], point[1], int(point[2]))
	}
	*/
}

func BenchmarkDBScanSimple(b *testing.B) {
	maxDistance := float64(2.0)
	minSamples := 3
	data, _ := getData("2d-10c")

	for n := 0; n < b.N; n++ {
		DBScan(data, maxDistance, minSamples)
	}
}

func getData(name string) ([]DataPoint, error) {
	data, err := os.ReadFile("./data-" + name + ".csv")
	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(data), "\n")

	points := make([]DataPoint, len(rows))

	var x float64
	var y float64
	var category float64
	for idx, row := range rows {
		if row == "" {
			continue
		}

		parts := strings.Split(row, ",")

		x, err = strconv.ParseFloat(parts[0], 32)
		if err != nil {
			return nil, err
		}

		y, err = strconv.ParseFloat(parts[1], 32)
		if err != nil {
			return nil, err
		}

		category, err = strconv.ParseFloat(parts[2], 32)
		if err != nil {
			return nil, err
		}

		points[idx] = DataPoint{float64(x), float64(y), float64(category)}
	}

	return points, nil
}
