package main

import (
	"fmt"
	"math"
)

func powerNumber(numberToPower, power, limit float64) float64 {
	if result := math.Pow(numberToPower, power); result < limit {
		return result
	}
	return limit
}

func main() {
	fmt.Println(
		powerNumber(3, 2, 10),
		powerNumber(3, 3, 20),
	)
}
