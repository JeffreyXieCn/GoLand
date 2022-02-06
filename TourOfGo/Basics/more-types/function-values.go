package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("=============Function closures=============")
	// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Closures
	//	Closures are useful because they let you associate data (the lexical environment) with a function that operates
	//	on that data. This has obvious parallels to object-oriented programming, where objects allow you to associate
	//	data (the object's properties) with one or more methods.
	//
	//	Consequently, you can use a closure anywhere that you might normally use an object with only a single method.

	pos, neg := adder(), adder() // pos and neg are closures with reference to its own sum
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
