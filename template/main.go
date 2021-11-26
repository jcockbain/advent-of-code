package main

import (
	input "aoc2021/inpututils"
	"os"
	"time"

	"fmt"
)

func main() {
	start := time.Now()
	input := getInput()

	fmt.Println("--- Part One ---")
	fmt.Println(Part1(input))
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 1", elapsed)

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2(input))
	fmt.Printf("%s took %s seconds \n", "Part 2", time.Since(start)-elapsed)
}

func getInput() string {
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		input = "input.txt"
	}
	return input
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
