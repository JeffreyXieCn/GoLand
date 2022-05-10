/*
Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with
import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you
need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
*/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14

const (
	// Big Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100

	// Small Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables declared without an explicit initial value are given their zero value
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// in Go assignment between items of different type requires an explicit conversion
	var x, y int = 3, 4
	var f2 float64 = math.Sqrt(float64(x*x + y*y))
	var z2 uint = uint(f)
	fmt.Println(x, y, f2, z2)

	v1 := 42           // int
	v2 := 3.142        // float64
	v3 := 0.867 + 0.5i // complex128
	fmt.Printf("v1 is of type %T\n", v1)
	fmt.Printf("v2 is of type %T\n", v2)
	fmt.Printf("v3 is of type %T\n", v3)

	// Constants cannot be declared using the := syntax
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	//fmt.Println(needInt(Big)) // constant 1267650600228229401496703205376 overflows int
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println("========Be careful of float32========")
	var f3 float32 = 16777216 // 1 << 24
	fmt.Println(f3 == f3+1)   // true !
	var f4 = f3 + 1
	fmt.Println(f3 == f4)   // true !
	fmt.Println(f3 == f3+2) // false !

	fmt.Println("========Complex number========")
	fmt.Println(1i * 1i) // (-1+0i)

	fmt.Println("========String========")
	str := "Hello 世界"
	fmt.Println(len(str)) // 12
	fmt.Println(str[6:9]) // 世
	fmt.Println(str[9:])  // 界

	str2 := "Hello \u4e16\u754c"
	fmt.Println(str2) // 12

	// var str3 string = nil // Can't use nil as the type of string

}
