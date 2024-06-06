package binsearch

import (
	"fmt"
	"testing"
)

var list = []int{-9, -1, 0, 2, 8, 9, 13, 15}

func TestOne(t *testing.T) {
	target := 9
	expectedIndex := 5

	testInternal(target, expectedIndex, t)
}

func TestTwo(t *testing.T) {
	target := -1
	expectedIndex := 1

	testInternal(target, expectedIndex, t)
}

func testInternal(target int, expectedIndex int, t *testing.T) {
	actualIndex := BinSearch(list, target)

	if actualIndex != expectedIndex {
		t.Errorf(fmt.Sprintf("Expected to find %d at index %d, got %d", target, expectedIndex, actualIndex))
	}
}
