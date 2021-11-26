package main

import (
	input "aoc2021/inpututils"
	"os"
	"time"

	"fmt"
)

func main() {
	start := time.Now()
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		input = "input.txt"
	}

	fmt.Println("--- Part One ---")
	fmt.Println(Part1(input))
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 1", elapsed)

	start = time.Now()
	fmt.Println("--- Part Two ---")
	fmt.Println(Part2(input))
	elapsed = time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 2", elapsed)
}

func Part1(filename string) int {
	nums := input.ReadNumbers(filename)
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func Part2(filename string) int {
	return 12
}
