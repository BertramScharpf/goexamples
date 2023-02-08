package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Sqrt: negative number %g", e)
}

const delta = 1e-10

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}
	z := f
	for {
		n := z - (z*z-f)/(2*z)
		if math.Abs(n-z) < delta {
			break
		}
		z = n
	}
	return z, nil
}

func main() {
	res, err := Sqrt(2)
	fmt.Println(res, err)

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}