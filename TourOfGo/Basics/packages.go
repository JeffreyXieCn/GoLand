package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMilli())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(math.Pi)
}
