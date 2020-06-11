package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    result, error:=sqrt(rand.Float64() * 100)

    if(error != nil) {
	fmt.Println(error)
    } else {
	fmt.Println(result)
    }
}

func sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, errors.New("Undefined for negative numbers")
    }

    return math.Sqrt(x), nil
}
