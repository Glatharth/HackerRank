package main

import "fmt"

type Node struct {
	data int32
	side bool
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {

}

func main() {
	idx := [][]int32{{2, 3}, {4, -1}, {5, -1}, {6, -1}, {7, 8}, {-1, 9}, {-1, -1}, {10, 11}, {-1, -1}, {-1, -1}, {-1, -1}}
	query := []int32{2, 4}
	fmt.Println(swapNodes(idx, query))
}
