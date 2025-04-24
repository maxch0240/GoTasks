package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	a = factorial(a)
	fmt.Printf("factorial of a: %v", a)
}

func factorial(a int) int {
	if a == 1 || a == 0 {
		return 1
	} else {
		return a * factorial(a-1)
	}
}
