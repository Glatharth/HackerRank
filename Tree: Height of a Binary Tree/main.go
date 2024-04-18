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

func (bst *BST) getHeight2() int {
	return bst.getHeightRec2(bst.root)
}

func (bst *BST) getHeightRec2(node *Node) int {
	if node == nil {
		return -1
	} else {
		leftHeight := bst.getHeightRec(node.left)
		rightHeight := bst.getHeightRec(node.right)
		if leftHeight > rightHeight {
			return leftHeight + 1
		} else {
			return rightHeight + 1
		}
	}
}

func (bst *BST) getHeight() int {
	return bst.getHeightRec(bst.root)
}

func (bst *BST) getHeightRec(node *Node) int {
	if node == nil {
		return -1
	} else {
		//leftHeight := bst.getHeightRec(node.left)
		//rightHeight := bst.getHeightRec(node.right)
		//if leftHeight > rightHeight {
		//	return leftHeight + 1
		//} else {
		//	return rightHeight + 1
		//}
		return max(bst.getHeightRec(node.left), bst.getHeightRec(node.right)) + 1
	}
}

func main() {
	bst := BST{}
	//bst.Insert(3) // return 0
	//bst.Insert(5)
	//bst.Insert(2)
	//bst.Insert(1)
	//bst.Insert(4)
	//bst.Insert(6)
	//bst.Insert(7) // return 3
	// ------------------------------
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(7)
	bst.Insert(5)
	bst.Insert(4) // return 3
	//bst.Inorder(bst.root)
	fmt.Print(bst.getHeight())
}
