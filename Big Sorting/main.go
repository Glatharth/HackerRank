package main

import (
	"fmt"
)

/*
 * Complete the 'bigSorting' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY unsorted as parameter.
 */

//	Brute force solution
//	func bigSorting(unsorted []string) []string {
//		size := len(unsorted)
//		for i := 0; i < size-1; i++ {
//			for j := i + 1; j < size; j++ {
//				sortedI, sortedJ := 0, 0
//				if len(unsorted[i]) > len(unsorted[j]) {
//					sortedI = 1
//				} else if len(unsorted[i]) < len(unsorted[j]) {
//					sortedJ = 1
//				} else {
//					for k, _ := range unsorted[i] {
//						if unsorted[i][k] > unsorted[j][k] {
//							sortedI = 1
//							break
//						} else if unsorted[i][k] < unsorted[j][k] {
//							sortedJ = 1
//							break
//						}
//					}
//				}
//				if sortedI > sortedJ {
//					unsorted[i], unsorted[j] = unsorted[j], unsorted[i]
//				}
//			}
//		}
//		return unsorted
//	}

func verify(a string, b string) bool {
	if len(a) > len(b) {
		return false
	} else if len(a) < len(b) {
		return true
	}
	for i := range a {
		if a[i] > b[i] {
			return false
		} else if a[i] < b[i] {
			return true
		}
	}
	return true
}

func bigSorting(items []string) []string {
	if len(items) < 2 {
		return items
	}
	first := bigSorting(items[:len(items)/2])
	second := bigSorting(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []string, b []string) []string {
	final := make([]string, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if verify(a[i], b[j]) {
			final[k] = a[i]
			i++
		} else {
			final[k] = b[j]
			j++
		}
		k++
	}
	for ; i < len(a); i++ {
		final[k] = a[i]
		k++
	}
	for ; j < len(b); j++ {
		final[k] = b[j]
		k++
	}
	return final[:k]
}

func main() {

	result := bigSorting([]string{"31415926535897932384626433832795", "1", "3", "10", "3", "5"})

	for i, resultItem := range result {
		fmt.Printf("%s", resultItem)

		if i != len(result)-1 {
			fmt.Printf("\n")
		}
	}
}
