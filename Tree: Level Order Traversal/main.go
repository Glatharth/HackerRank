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
	bst := BST{}
	// 1 2 5 3 6 4
	//bst.Insert(1)
	//bst.Insert(2)
	//bst.Insert(5)
	//bst.Insert(3)
	//bst.Insert(6)
	//bst.Insert(4)
	//bst.Insert(9)
	//bst.Insert(7)
	//bst.Insert(8)
	//bst.Insert(10)
	bst.Insert(6)
	bst.Insert(98)
	bst.Insert(18)
	bst.Insert(13)
	bst.Insert(42)
	bst.Insert(23)
	bst.Insert(69)
	bst.Insert(26)
	bst.Insert(68)
	bst.Insert(77)
	bst.Insert(41)
	bst.Insert(46)
	bst.Insert(73)
	bst.Insert(85)
	bst.Insert(89)
	// 6 98 18 13 42 23 69 26 41 68 77 46 73 85 89
	fmt.Print(bst.levelOrder())
}
