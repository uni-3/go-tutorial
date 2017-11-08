package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	fields := strings.Fields(s)
	for _, field := range fields {
		counts[field] += 1
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
