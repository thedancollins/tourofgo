package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	//
	sm := strings.Fields(s)
	m := make(map[string]int)
	for _, value := range sm {
		//if value not in sm, add it with 1 or increment
		v, ok := m[value]
		if ok {
			m[value] = v + 1
		} else {
			m[value] = 1
		}

	}

	return m
}

func main() {
	wc.Test(WordCount)
}
