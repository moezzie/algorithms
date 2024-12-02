package day1

import (
	"fmt"
	"io/ioutil"
	"math"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	lists := [][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}}
	lists = sortLists(lists)

	totalDistance := 0
	for n := 0; n < len(lists[0]); n++ {
		totalDistance += int(math.Abs(float64(lists[0][n] - lists[1][n])))
	}

	if totalDistance != 11 {
		t.Fatalf("Expected 11, got %d", totalDistance)
	}
}

func TestPart1(t *testing.T) {
	contentLines := readInputFile("input.csv")
	lists := constructLists(contentLines)
	lists = sortLists(lists)

	totalDifference := 0
	for n := 0; n < len(lists[0]); n++ {
		totalDifference += int(math.Abs(float64(lists[0][n] - lists[1][n])))
	}

	fmt.Println("Total difference ", totalDifference)
}

func TestExamplePart2(t *testing.T) {
	lists := [][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}}
	lists = sortLists(lists)

	similarityScore := calculateSimilarityScore(lists)

	if similarityScore != 31 {
		t.Fatalf("Expected 31, got %d", similarityScore)
	}
}

func TestPart2(t *testing.T) {
	contentLines := readInputFile("input.csv")
	lists := constructLists(contentLines)
	lists = sortLists(lists)

	similarityScore := calculateSimilarityScore(lists)

	fmt.Println("Similarity score ", similarityScore)
}

func readInputFile(filePath string) []string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	return lines
}

func constructLists(lines []string) [][]int {
	numLines := len(lines)
	left := make([]int, numLines)
	right := make([]int, numLines)

	var parts []string
	for n := 0; n < numLines; n++ {
		if len(lines[n]) > 0 {
			parts = strings.Split(lines[n], "   ")
			left[n], _ = strconv.Atoi(parts[0])
			right[n], _ = strconv.Atoi(parts[1])
		}
	}

	output := make([][]int, 2)
	output[0] = left
	output[1] = right

	return output
}

func sortLists(lists [][]int) [][]int {
	slices.Sort(lists[0])
	slices.Sort(lists[1])

	return lists
}

func calculateSimilarityScore(lists [][]int) int {
	var similarityScore int
	var current int
	for n := 0; n < len(lists[0]); n++ {
		current = lists[0][n]
		for i := 0; i < len(lists[1]); i++ {
			if current == lists[1][i] {
				similarityScore += current
			}
		}
	}

	return similarityScore
}
