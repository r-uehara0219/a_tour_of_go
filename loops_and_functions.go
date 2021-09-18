package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x/2
	loopCount := 0
	for {
		loopCount++
		zdiff := (z*z - x) / (2 * z)
		if math.Abs(zdiff) <= 1e-10 {
			break
		} else {
			z -= zdiff
		}
	}
	fmt.Println(loopCount)
	return z
}

func main() {
	fmt.Println(Sqrt(123412341234))
}
