package dfs

import (
	"reflect"
	"testing"
)

// TestDFSBasic tests BFS on a simple binary tree
func TestDFS(t *testing.T) {
	root := getTree()

	expected := []int{1, 2, 3, 4, 5, 6, 7, 9}
	result := DFS(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v; got %v", expected, result)
	}
}

func TestDFSRecurseive(t *testing.T) {
	root := getTree()

	expected := []int{1, 2, 3, 4, 5, 6, 7, 9}
	result := make([]int, 0)
	DFSRecursive(root, &result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v; got %v", expected, result)
	}
}

func BenchmarkDFS(t *testing.B) {
	root := getTree()
	for n := 0; n < 10_000; n++ {
		DFS(root)
	}
}

func BenchmarkDFSRecurse(t *testing.B) {
	root := getTree()

	output := make([]int, 0)
	for n := 0; n < 10_000; n++ {
		DFSRecursive(root, &output)
	}
}

func getTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 9}
	return root
}
