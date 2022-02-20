package main

import (
	"example/hello/morestrings"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"rsc.io/quote/v4"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Glass())

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
