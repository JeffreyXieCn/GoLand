package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs2(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale2(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex2{3, 4}
	v.Scale(10) // Go will interpret it as
	// (&v).Scale(10) // and vice-versa
	fmt.Println(v.Abs())

	//Here we see the Abs and Scale methods rewritten as functions
	v2 := Vertex2{3, 4}
	Scale2(&v2, 10)
	fmt.Println(Abs2(v2))
}
