package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	fmt.Println("============================")
	// Slices are like references to arrays, A slice does not store any data, it just describes a section of an
	// underlying array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a2 := names[0:2]
	b := names[1:3]
	fmt.Println(a2, b)

	b[0] = "XXX"
	fmt.Println(a2, b)
	fmt.Println(names)

	fmt.Println("============================")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("q:", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("r:", r)

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)

	fmt.Println("============================")
	s2 := []int{2, 3, 5, 7, 11, 13}

	s2 = s2[1:4]
	fmt.Println(s2) // 3 ,5, 7

	// The default is zero for the low bound and the length of the slice for the high bound
	s2 = s2[:2]
	fmt.Println(s2) // 3, 5

	s2 = s2[1:]
	fmt.Println(s2) // 5

	fmt.Println("============================")
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	s3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s3) // 6, 6, [2, 3, 5, 7, 11, 13]

	// Slice the slice to give it zero length.
	s3 = s3[:0]
	printSlice(s3) // 0, 6, []

	// Extend its length.
	s3 = s3[:4]
	printSlice(s3) // 4, 6, [2, 3, 5, 7]

	// Drop its first two values.
	s3 = s3[2:]
	printSlice(s3) // 2, 4, [5, 7]

	fmt.Println("============================")
	var zeroSlice []int
	fmt.Println(zeroSlice, len(zeroSlice), cap(zeroSlice))
	if zeroSlice == nil {
		fmt.Println("nil!")
	}

	fmt.Println("============================")
	aSliceMade := make([]int, 5)
	printSlice2("aSliceMade", aSliceMade)

	bSliceMade := make([]int, 0, 5)
	printSlice2("bSliceMade", bSliceMade)

	cSlice := bSliceMade[:2]
	printSlice2("cSlice", cSlice)

	dSlice := cSlice[2:5]
	printSlice2("dSlice", dSlice)

	fmt.Println("============================")
	// Slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("=============Appending to a slice===============")
	var sliceToAppend []int
	printSlice(sliceToAppend)

	// append works on nil slices.
	sliceToAppend = append(sliceToAppend, 0)
	printSlice(sliceToAppend)

	// The slice grows as needed.
	sliceToAppend = append(sliceToAppend, 1)
	printSlice(sliceToAppend)

	// We can add more than one element at a time.
	sliceToAppend = append(sliceToAppend, 2, 3, 4)
	printSlice(sliceToAppend) // len=5 cap=6 [0 1 2 3 4]

	sliceToAppend = append(sliceToAppend, 5)
	printSlice(sliceToAppend) // len=6 cap=6 [0 1 2 3 4 5]

	sliceToAppend = append(sliceToAppend, 6)
	printSlice(sliceToAppend) // len=7 cap=12 [0 1 2 3 4 5 6]

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
