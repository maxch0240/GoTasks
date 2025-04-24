package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	runes := []rune(a)

	for i := 0; i < len(runes)/2; i++ {
		buff := runes[len(runes)-i-1]
		runes[len(runes)-i-1] = runes[i]
		runes[i] = buff
	}

	fmt.Println(string(runes))
}
