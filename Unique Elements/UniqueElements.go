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

	seen := make([]bool, size)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				seen[j] = true
			}
		}
	}

	result := []int{}

	for i := 0; i < len(arr); i++ {
		if !seen[i] {
			result = append(result, arr[i])
		}
	}

	fmt.Println(result)
}
