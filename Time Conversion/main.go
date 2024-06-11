package main

import (
	"fmt"
	"strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	if s[len(s)-2:] == "PM" {
		if s[:2] != "12" {
			convert, _ := strconv.Atoi(s[:2])
			convert += 12
			s = fmt.Sprintf("%d", convert) + s[2:]
		}
	} else if s[:2] == "12" {
		s = "00" + s[2:]
	}
	s = s[:len(s)-2]
	return s
}

func main() {
	fmt.Println(timeConversion("12:00:00PM"))
}
