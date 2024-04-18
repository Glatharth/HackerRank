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

func (bst *BST) Insert(data int) {
	bst.InsertRec(bst.root, data)
}
func (bst *BST) InsertRec(node *Node, data int) *Node {
	if bst.root == nil {
		bst.root = &Node{data, nil, nil}
		return bst.root
	}
	if node == nil {
		return &Node{data, nil, nil}
	}
	if data <= node.data {
		node.left = bst.InsertRec(node.left, data)
	}
	if data > node.data {
		node.right = bst.InsertRec(node.right, data)
	}
	return node
}

func (bst *BST) levelOrder() string {
	if bst.root == nil {
		return ""
	}
	queue := []*Node{bst.root}
	var result string

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		result += fmt.Sprintf("%d ", curr.data)
		if curr.left != nil {
			queue = append(queue, curr.left)
		}
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
	return result
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	bst := BST{}
	var n int
	fmt.Scanln(&n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}
	for _, data := range numbers {
		bst.Insert(data)
	}
	fmt.Print(bst.levelOrder())
}
