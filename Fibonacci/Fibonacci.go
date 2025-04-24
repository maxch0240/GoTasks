package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter Fibonachi sequence size")
	fmt.Scanln(&size)

	fmt.Println(fibonacci(size))
}

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	fib := make([]int, n)
	fib[0], fib[1] = 0, 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}
