package main

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

const delta = 0.000001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e) // this will cause infinite loop
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var z = x / 2
	var prevZ float64
	for i := 1; ; i++ {
		prevZ = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)
		if math.Abs(z-prevZ) <= delta {
			break
		}
	}

	return z, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
