package main

import (
	"bufio"
	"fmt"
	"os"
)

func isBalanced(s string) string {
	if len(s)%2 == 1 {
		return "NO"
	}
	brackets := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	var stack []string
	for _, c := range s {
		char := string(c)
		if brackets[char] != "" {
			stack = append(stack, brackets[char])
		} else if len(stack) > 0 {
			for _, bracket := range brackets {
				if bracket == char && bracket == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
					break
				}
			}
		} else {
			return "NO"
		}
	}
	if len(stack) > 0 {
		return "NO"
	}
	return "YES"
}

func main() {
	var inputLines []string
	input, err := os.Open("input/input20.txt")
	if err != nil {
		panic(err)
	}
	inputScanner := bufio.NewScanner(input)
	inputScanner.Split(bufio.ScanLines)
	i := 1
	for inputScanner.Scan() {
		if i > 1 {
			inputLines = append(inputLines, isBalanced(inputScanner.Text()))
		}
		i++
	}
	input.Close()

	output, err := os.Open("output/output20.txt")
	if err != nil {
		panic(err)
	}
	outputScanner := bufio.NewScanner(output)
	outputScanner.Split(bufio.ScanLines)
	i = 0
	for outputScanner.Scan() {
		//fmt.Println(outputScanner.Text(), inputLines[i])
		if outputScanner.Text() != inputLines[i] {
			fmt.Println(i, inputLines[i])
		}
		i++
	}
	output.Close()
}
