package main

import (
	"fmt"
)

func main() {
	workingDays := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
	}

	daysOff := []string{
		"Saturday",
		"Sunday",
	}

	weekDays := append(workingDays, daysOff...)

	fmt.Printf("Вся неделя: %v\n", weekDays)

}
