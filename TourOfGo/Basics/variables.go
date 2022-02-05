package main

import "fmt"

var c, python, java bool

var k, j int = 1, 2

func main() {
	var i int
	fmt.Println(i, c, python, java)

	var c2, python2, java2 = true, false, "no!"
	fmt.Println(k, j, c2, python2, java2)

	var i2, j2 int = 1, 2
	k2 := 3
	c3, python3, java3 := true, false, "no!"

	fmt.Println(i2, j2, k2, c3, python3, java3)
}
