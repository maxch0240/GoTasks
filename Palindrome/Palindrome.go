package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	runes := []rune(a)

	var isPalindrome bool = true

	for i := range len(runes) / 2 {
		if runes[len(runes)-i-1] != runes[i] {
			isPalindrome = false
		}

	}

	fmt.Println(isPalindrome)
}
