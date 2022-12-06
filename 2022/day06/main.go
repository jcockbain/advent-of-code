package main

import (
	_ "embed"

	"fmt"
)

var (
	benchmark = false
)

//go:embed input.txt
var input string

func main() {
	p1 := solve(4)
	p2 := solve(14)

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func solve(msgLen int) int {
	for i := 0; i < len(input)-msgLen; i++ {
		if !repeats(input[i : i+msgLen]) {
			return i + msgLen
		}
	}
	panic("no repeats")
}

func repeats(s string) bool {
	set := map[rune]bool{}
	for _, c := range s {
		if _, ok := set[c]; ok {
			return true
		}
		set[c] = true
	}
	return false
}
