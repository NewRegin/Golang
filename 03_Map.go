//练习3.map exercise
//扩展了相关方法，未体现在代码上
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
