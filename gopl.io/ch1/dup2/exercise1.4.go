// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type void struct{}

var empty void

type LineCounter struct {
	n         int
	fileNames map[string]void
}

func main() {
	counts := make(map[string]LineCounter)
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines2(f, counts)
		f.Close()
	}

	fmt.Println(counts)
	for line, lc := range counts {
		if lc.n > 1 {
			fmt.Printf("%d\t%s\t%v\n", lc.n, line, mapKeyToString(lc.fileNames))
		}
	}
}

func countLines2(f *os.File, counts map[string]LineCounter) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		val, ok := counts[input.Text()] // is the val by reference or by value? It's by value!
		if ok {
			val.n = val.n + 1
			val.fileNames[f.Name()] = empty

			counts[input.Text()] = val // this is important to update the value in the map
		} else {
			set := make(map[string]void)
			set[f.Name()] = empty
			counts[input.Text()] = LineCounter{1, set}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func mapKeyToString(m map[string]void) string {
	keys := make([]string, len(m))

	i := 0
	for key := range m { // more efficient than append
		keys[i] = key
		i++
	}
	//for key := range m {
	//	keys = append(keys, key)
	//}

	return strings.Join(keys, " ")
}

//!-
