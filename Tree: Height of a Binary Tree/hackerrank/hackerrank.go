package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func height(root *Node) int {
	if root == nil {
		return -1
	} else {
		leftHeight := height(root.left)
		rightHeight := height(root.right)
		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
	}
}

func (bst *BST) Insert(node *Node, data int) *Node {
	if bst.root == nil {
		bst.root = &Node{data, nil, nil}
		return bst.root
	}
	if node == nil {
		return &Node{data, nil, nil}
	}
	if data <= node.data {
		node.left = bst.Insert(node.left, data)
	}
	if data > node.data {
		node.right = bst.Insert(node.right, data)
	}
	return node
}

func main() {
	bst := BST{}
	var n int
	fmt.Scanln(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	for _, data := range numbers {
		bst.Insert(bst.root, data)
	}

	fmt.Print(height(bst.root))
}
