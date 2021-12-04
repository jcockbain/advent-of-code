package main

import (
	"time"

	input "github.com/jcockbain/advent-of-code-2021/inpututils"

	"fmt"
)

func main() {
	start := time.Now()
	input := input.GetInputPath()

	fmt.Println("--- Part One ---")
	fmt.Println(part1(input))
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 1", elapsed)

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(input))
	fmt.Printf("%s took %s seconds \n", "Part 2", time.Since(start)-elapsed)
}

func part1(filename string) int {
	nums := input.ReadNumbers(filename)
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func part2(filename string) int {
	return 12
}
