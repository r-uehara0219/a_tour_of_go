package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := x / 2
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
	return z, nil
}

func main() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
