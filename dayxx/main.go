package main

import (
	_ "embed"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
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
	nums := utils.GetInts(input)
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func part2() int {
	return 12
}
