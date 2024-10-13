package day7

import (
	"fmt"
	"strings"
)

type Dir struct {
	Parent   *Dir
	Children []*Dir
	Size     int
}

func buildTree(input string) *Dir {
	lines := strings.Split(input, "\n")

	root := Dir{}

	currentDir := &root
	var childDir *Dir

	for _, line := range lines {
		if line[0] == '$' {
			if line[2] == 'c' {
				// Back out of dir
				if line[5] == '.' {
					currentDir.Parent.Size += currentDir.Size
					currentDir = currentDir.Parent
				} else {
					// Go deeper into structure
					childDir = &Dir{Parent: currentDir}
					currentDir.Children = append(currentDir.Children, childDir)
					currentDir = childDir
				}
			}
			// If line starts with an integer it is a file size
		} else if isInt(line[0]) {
			// Assume no file size is larger than 999999
			// Use conditions instead of for loop for performance
			if line[6] == ' ' {
				currentDir.Size += strToInt(line[:6])
			} else if line[5] == ' ' {
				currentDir.Size += strToInt(line[:5])
			} else if line[4] == ' ' {
				currentDir.Size += strToInt(line[:4])
			} else if line[3] == ' ' {
				currentDir.Size += strToInt(line[:3])
			} else if line[2] == ' ' {
				currentDir.Size += strToInt(line[:2])
			} else {
				currentDir.Size += strToInt(line[:1])
			}
		}
	}

	root.Size = root.Children[0].Size

	return &root
}

func isInt(char byte) bool {
	return char >= 0x30 && char <= 0x39
}

func strToInt(number string) int {
	var integer int
	var x int

	for _, char := range number {
		x = int(char - 0x30)
		integer = integer*10 + x
	}

	return integer
}

func bfsLessThan100000(root *Dir) int {
	stack := []*Dir{root}
	idx := 0

	var totalSize int
	var current *Dir

	for idx < len(stack) {
		current = stack[idx]
		idx++

		if current.Size <= 100_000 {
			totalSize += current.Size
		}

		stack = append(stack, current.Children...)
	}

	return totalSize
}

func printNode(node *Dir, level int) {
	indentation := indentation(level)
	fmt.Printf("%s%d\n", indentation, node.Size)

	if len(node.Children) > 0 {
		for _, child := range node.Children {
			printNode(child, level+1)
		}
	}
}

func indentation(levels int) string {
	indent := ""
	for n := 0; n < levels; n++ {
		indent = indent + "  "
	}

	return indent
}
