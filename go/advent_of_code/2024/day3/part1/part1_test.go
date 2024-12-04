package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestExamplePuzzle1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	total := calculateTotalFromInput(input)

	if total != 161 {
		t.Fatalf("Expected 161, got %d", total)
	}
}

func TestPuzzle1(t *testing.T) {
	input := readInput("../input.txt")

	total := calculateTotalFromInput(input)

	fmt.Println("Total: ", total)
}

func calculateTotalFromInput(memory string) int {
	instructions := findMulInstructionsWithRE(memory)

	multiplications := convertInstructions(instructions)

	return calculateInstructions(multiplications)
}

func findMulInstructionsWithRE(memory string) []string {
	pattern, err := regexp.Compile("(mul\\([0-9]+,[0-9]+\\))")
	checkErr(err)

	instructions := pattern.FindAllString(memory, -1)

	return instructions
}

func convertInstructions(instructions []string) [][2]int {
	multiplications := make([][2]int, len(instructions))
	for idx, instruction := range instructions {
		multiplications[idx] = convertInstruction(instruction)
	}

	return multiplications
}

func convertInstruction(instruction string) [2]int {
	var err error

	startPos := strings.Index(instruction, "(")
	endPos := strings.Index(instruction, ")")

	digits := instruction[startPos+1 : endPos]
	parts := strings.Split(digits, ",")

	multiplication := [2]int{}
	for d, digit := range parts {
		multiplication[d], err = strconv.Atoi(digit)
		checkErr(err)
	}

	return multiplication
}

func calculateInstructions(multiplications [][2]int) int {
	var total int
	for _, multiplication := range multiplications {
		total += multiplication[0] * multiplication[1]
	}

	return total
}

func readInput(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
