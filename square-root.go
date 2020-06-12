package main

import (
	"fmt"
)

func computeSquareRoot(x float64) float64 {
	// 1 is always a good starting guess
	squareValue := 1.0

	for i:= 0; i < 10; i++ {
		squareValue -= (squareValue*squareValue - x) / (2*squareValue)
	}
	return squareValue
}

func main() {
	fmt.Println(computeSquareRoot(2))
}
