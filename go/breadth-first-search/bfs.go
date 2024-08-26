package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFS(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	output := make([]int, 0)
	queue := []*TreeNode{root}
	var current *TreeNode

	for len(queue) >= 1 {
		current = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		output = append([]int{current.Val}, output...)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return output
}
