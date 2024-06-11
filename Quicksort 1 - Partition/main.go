package main

import "fmt"

/*
 * Complete the 'quickSort' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func quickSort(arr []int32) []int32 {
	var left, right []int32
	p := arr[0]
	for i := 1; i <= len(arr)-1; i++ {
		value := arr[i]
		if value > p {
			right = append(right, value)
		} else {
			left = append(left, value)
		}
	}
	arr = []int32{}
	arr = append(arr, left...)
	arr = append(arr, p)
	arr = append(arr, right...)
	return arr
}

func main() {
	result := quickSort([]int32{4, 5, 3, 7, 2})
	fmt.Println(result)
}
