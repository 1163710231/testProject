package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word] += 1
	}
	return m
}

func main() {
	s := "aaa bbb aaa vvv ccc"
	fmt.Println(WordCount(s))
}
