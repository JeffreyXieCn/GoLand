//
// The Go Programming Language - Excercise 1.4
//
// Mingjie Li (limingjie@outlook.com)
// Jun 16, 2016
//

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines3(os.Stdin, "os.Stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines3(f, arg, counts)
			f.Close()
		}
	}
	for line, filenames := range counts {
		fileCount := len(filenames)
		if fileCount == 1 {
			total := 0
			for _, count := range filenames {
				total += count
			}
			if total <= 1 {
				continue
			}
		}

		fmt.Printf("[Found in %d file(s)]\t%s\n", fileCount, line)
		for name, count := range filenames {
			fmt.Printf("\t%d hit(s) in %s\n", count, name)
		}
	}
}

func countLines3(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
