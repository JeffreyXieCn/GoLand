/*
If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to
be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing.
This general approach is called Newton's method. It works well for many functions but especially well for square root.
*/
package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)
	}

	return z
}

const delta = 0.000001

func Sqrt2(x float64) float64 {
	var z = 1.0
	var prevZ float64
	for i := 1; ; i++ {
		prevZ = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)
		if math.Abs(z-prevZ) <= delta {
			break
		}
	}

	return z
}

func Sqrt3(x float64) float64 {
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

	return z
}

func main() {
	for k := 1; k <= 5; k++ {
		fmt.Printf("Sqrt1(%d) = %g\n\n", k, Sqrt1(float64(k)))
	}

	for k := 1; k <= 5; k++ {
		fmt.Printf("Sqrt2(%d) = %g\n\n", k, Sqrt2(float64(k)))
	}

	for k := 1; k <= 5; k++ {
		fmt.Printf("Sqrt3(%d) = %g\n\n", k, Sqrt3(float64(k)))
	}
}
