package main

import (
	"fmt"
	"math"
)

type AmericanVelocity float64

type EuropeanVelocity float64

func main() {
	var speed1 float64 = 102.4
	var speed2 float64 = 130

	var europeanSpeed EuropeanVelocity = EuropeanVelocity(round(speed1 * 3.6))
	var americanSpeed AmericanVelocity = AmericanVelocity(round(speed2 * 2.237))

	fmt.Printf("Евро скорость: %v\n", europeanSpeed)
	fmt.Printf("Американская скрость: %v", americanSpeed)
}

func round(num float64) float64 {
	return math.Round(num*100) / 100
}
