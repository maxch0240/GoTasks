package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "104"
	var num int = 35

	fmt.Println("Строка к числу")
	num2, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}
	fmt.Println(num2)

	fmt.Println("Число к строке")
	str2 := strconv.Itoa(num)
	fmt.Println(str2)
}
