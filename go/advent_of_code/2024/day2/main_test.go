package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestExample1(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	numValidReports := validateReports(reports)

	if numValidReports != 2 {
		t.Fatalf("Expected 2 got %d", numValidReports)
	}
}

func TestNumValidReports(t *testing.T) {
	inputs := readInput("input.csv")
	reports := formatReports(inputs)

	numValidReports := validateReports(reports)

	fmt.Println("Num valid reports ", numValidReports)
}

func validateReports(reports [][]int) int {
	var numValidReports int
	for _, report := range reports {
		if validateReport(report) {
			numValidReports++
			fmt.Println("v", report)
		} else {
			fmt.Println(" ", report)
		}
	}

	return numValidReports
}

func validateReport(report []int) bool {
	var lastDirection int
	var difference int

	if len(report) == 0 {
		return false
	}

	for n := 0; n < len(report)-1; n++ {
		difference = report[n] - report[n+1]

		if difference > 3 || difference < -3 {
			return false
		}

		if report[n]-report[n+1] == 0 {
			return false
		} else if report[n] < report[n+1] {
			// Acending
			if lastDirection == -1 {
				return false
			}
			lastDirection = 1
		} else if report[n] > report[n+1] {
			// Decending
			if lastDirection == 1 {
				return false
			}
			lastDirection = -1
		}
	}

	return true
}

func formatReports(input []string) [][]int {
	var level int
	var err error

	reports := make([][]int, len(input))

	for reportIdx, line := range input {

		if len(line) == 0 {
			continue
		}

		levels := strings.Split(line, " ")

		report := make([]int, len(levels))
		for levelIdx, levelStr := range levels {
			// convert to integer
			level, err = strconv.Atoi(levelStr)
			if err != nil {
				panic(err)
			}
			// add to report at[lineIdx][levelIdx]
			report[levelIdx] = level
		}
		reports[reportIdx] = report
	}

	return reports
}

func readInput(filePath string) []string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	return lines
}
