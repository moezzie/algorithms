package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack based implementation of DFS
func DFS(root *TreeNode) []int {
	stack := []*TreeNode{root}
	var current *TreeNode
	output := make([]int, 0)

	for len(stack) > 0 {
		// Get the last element and
		// 'pop' the it from the stack (LIFO)
		current = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		output = append(output, current.Val)

		// Add the right child first
		// This preserves the order of check Left, then Right
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		// Add the left child last.
		// We are using a Stack(LIFO), so it will get
		// traversed first
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return output
}

// Recursion based implementaiton of DFS
func DFSRecursive(node *TreeNode, output *[]int) {
	*output = append(*output, node.Val)
	if node.Left != nil {
		DFSRecursive(node.Left, output)
	}
	if node.Right != nil {
		DFSRecursive(node.Right, output)
	}
}
