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

	var max int = arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}

	fmt.Println(max)
}
