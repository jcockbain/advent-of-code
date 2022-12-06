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
	p1 := part1()
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	for i := 0; i < len(input)-4; i++ {
		if !repeats(input[i : i+4]) {
			return i + 4
		}
	}
	panic("no repeats")
}

func part2() int {
	for i := 0; i < len(input)-14; i++ {
		if !repeats(input[i : i+14]) {
			return i + 14
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
