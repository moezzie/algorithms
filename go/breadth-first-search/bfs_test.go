package main

import (
	"reflect"
	"testing"
)

func TestBFSSimple(t *testing.T) {
	root := getTree()

	expected := []int{3, 2, 1}
	result := BFS(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v; got %v", expected, result)
	}
}

func TestBFS2(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}

	expected := []int{1, 2, 3}
	result := BFS(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v; got %v", expected, result)
	}
}

func TestBFS3(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 1}

	expected := []int{2, 1, 4, 3}
	result := BFS(root)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v; got %v", expected, result)
	}
}

func BenchmarkDFSRecurse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BFS(getTree())
	}
}

func getTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	return root
}
