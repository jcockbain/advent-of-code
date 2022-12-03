package main

import (
	_ "embed"
	"fmt"

	utils "github.com/jcockbain/advent-of-code/utils"
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

func part1() (increased int) {
	nums := utils.GetInts(input)
	last := nums[0]
	for _, i := range nums[1:] {
		if i > last {
			increased++
		}
		last = i
	}
	return
}

func part2() (increased int) {
	nums := utils.GetInts(input)
	last, numsLength := sum(nums[0:3]), len(nums)
	for i := 3; i < numsLength-2; i++ {
		next := sum(nums[i : i+3])
		if next > last {
			increased++
		}
		last = next
	}
	return
}

func sum(nums []int) (ans int) {
	for _, a := range nums {
		ans += a
	}
	return
}
