package main

import (
	"fmt"
	"slices"
)

func main() {
	weekDays := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	workingDays := slices.Clone(weekDays[0:5])

	weekDays = slices.Delete(weekDays, 0, 5)

	fmt.Printf("Рабочие дни: %v\n", workingDays)
	fmt.Printf("Выходные дни: %v", weekDays)

}
