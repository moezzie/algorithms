package day7

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	buildTree(input())
}

func TestFindAtMost100_000(t *testing.T) {
	root := buildTree(input())
	totalSize := bfsLessThan100000(root)
	expected := 1844187
	if totalSize != expected {
		t.Errorf("Expected: %d, got: %d", expected, totalSize)
	}
}

func Benchmark(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buildTree(input())
	}
}

func TestInteger(t *testing.T) {
	number := "123"

	actual := strToInt(number)
	if actual != 123 {
		t.Errorf("Expected 123, got %d", actual)
	}
}
