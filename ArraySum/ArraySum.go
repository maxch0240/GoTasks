package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter arr size")
	fmt.Scanln(&size)

	var arr = make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter element %v\n", i)
		fmt.Scanln(&arr[i])
	}

	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)
}
