package main

import (
	"fmt"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func diagonalDifference(arr [][]int32) int32 {
	length := len(arr)
	var leftDiagonal, rightDiagonal int32
	for i := 0; i < length; i++ {
		leftDiagonal += arr[i][i]
		rightDiagonal += arr[i][length-i-1]
	}
	return abs(leftDiagonal - rightDiagonal)
}

func main() {
	result := diagonalDifference([][]int32{
		{11, 2, 4, 10, 8},
		{4, 5, 6, 12, 2},
		{10, 8, -12, 5, 8},
		{40, 1, -5, 1, 8},
		{1, 2, 3, 4, 5},
	})
	fmt.Printf("%d\n", result)

}
