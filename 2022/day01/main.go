package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"

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
	nums := getCalories(input)
	return max(nums)
}

func part2() int {
	nums := getCalories(input)
	sort.Ints(nums)
	return sum(nums[len(nums)-3:])
}

func toInt(s string) int {
	converted, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return converted
}

func getCalories(input string) []int {
	var calories []int
	elves := strings.Split(input, "\n\n")
	for _, e := range elves {
		elfTotal := 0
		for _, c := range strings.Split(strings.Trim(e, " "), "\n") {
			if c != "" {
				elfTotal += toInt(c)
			}
		}
		calories = append(calories, elfTotal)
	}
	return calories
}

func max(input []int) int {
	var m int
	for _, i := range input {
		if i > m {
			m = i
		}
	}
	return m
}

func sum(input []int) int {
	var r int
	for _, s := range input {
		r += s
	}
	return r
}
