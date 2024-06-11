package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func calculateStrings(str1, str2 string) string {
	len1, len2 := len(str1), len(str2)
	result := make([]int, len1+len2)

	for i := len1 - 1; i >= 0; i-- {
		carry := 0
		num1 := int(str1[i] - '0')

		for j := len2 - 1; j >= 0; j-- {
			num2 := int(str2[j] - '0')

			sum := num1*num2 + result[i+j+1] + carry
			carry = sum / 10
			result[i+j+1] = sum % 10
		}

		result[i] += carry
	}
	strResult := strings.TrimLeft(strings.TrimRight(fmt.Sprint(result), "0"), "[")
	strResult = strings.ReplaceAll(strResult, " ", "")
	strResult = strResult[:len(strResult)-1]
	if string(strResult[0]) == "0" {
		strResult = strResult[1:]
	}
	return strResult
}

func extraLongFactorials(n int32) {
	// Write your code here
	var result = "1"
	for i := int32(1); i <= n; i++ {
		result = calculateStrings(result, fmt.Sprintf("%v", i))
	}
	fmt.Println(result)
}

func main() {
	n := int32(3)
	extraLongFactorials(n)
}
