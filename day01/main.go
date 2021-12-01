package main

import (
	input "aoc2021/inpututils"
	"time"

	"fmt"
)

func main() {
	start := time.Now()
	input := input.GetInputPath()

	fmt.Println("--- Part One ---")
	fmt.Println(Part1(input))
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 1", elapsed)

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2(input))
	fmt.Printf("%s took %s seconds \n", "Part 2", time.Since(start)-elapsed)
}

func Part1(filename string) int {
	nums := input.ReadNumbers(filename)
	increased := 0
	last := nums[0]
	for _, i := range nums[1:] {
		if i > last {
			increased += 1
		}
		last = i
	}
	return increased
}

func Part2(filename string) int {
	nums := input.ReadNumbers(filename)
	increased := 0
	last := sum(nums[0:3])
	numsLength := len(nums) - 2
	for i := 3; i < numsLength; i++ {
		next := sum(nums[i : i+3])
		if next > last {
			increased += 1
		}
		last = next
	}
	return increased
}

func sum(nums []int) int {
	ans := 0
	for _, a := range nums {
		ans += a
	}
	return ans
}
