package main

import (
	"fmt"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	var pos, neg, zer float32
	for _, i := range arr {
		if i > 0 {
			pos++
		} else if i < 0 {
			neg++
		} else {
			zer++
		}
	}
	length := float32(len(arr))
	fmt.Printf("%.6f\n%.6f\n%.6f", pos/length, neg/length, zer/length)
}

func main() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}
