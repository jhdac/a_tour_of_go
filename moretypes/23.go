package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	counts := make(map[string]int)
	for i := 0; i < len(fields); i++ {
		counts[fields[i]]++
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
