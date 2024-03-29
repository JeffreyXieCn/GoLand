package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)

	for _, word := range words {
		//ele, exists := m[word]
		//if !exists {
		//	m[word] = 1
		//} else {
		//	m[word] = ele + 1
		//}
		m[word] = m[word] + 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
