package main

import "fmt"

func isBalanced(s string) string {
	brackets := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
	var stack []string
	for _, c := range s {
		char := string(c)
		if _, ok := brackets[char]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return "NO"
			}
			if brackets[stack[len(stack)-1]] == char {
				stack = stack[:len(stack)-1]
			} else {
				return "NO"
			}
		}
	}
	return "YES"
}

func main() {
	fmt.Print(isBalanced("{[()]}"))
	fmt.Print(isBalanced("{[(])}"))
	fmt.Print(isBalanced("{{[[(())]]}}"))
	fmt.Println()
	fmt.Print(isBalanced("{(([])[])[]}"))
	fmt.Print(isBalanced("{(([])[])[]]}"))
	fmt.Print(isBalanced("{(([])[])[]}[]"))
	fmt.Println()
	fmt.Print(isBalanced("{[{]}"))

}
