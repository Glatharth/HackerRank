package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func bigSorting(items []string) []string {
	if len(items) < 2 {
		return items
	}
	first := bigSorting(items[:len(items)/2])
	second := bigSorting(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []string, b []string) []string {
	var final []string
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var unsorted []string

	for i := 0; i < int(n); i++ {
		unsortedItem := readLine(reader)
		unsorted = append(unsorted, unsortedItem)
	}

	result := bigSorting(unsorted)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
