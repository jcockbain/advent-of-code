package main

import (
	"strconv"
	"time"

	input "github.com/jcockbain/advent-of-code-2021/inpututils"

	"fmt"
)

const MaxUint = ^uint(0)
const maxInt = int(MaxUint >> 1)

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
	inp := input.ReadSlice(filename)
	nums := mapToInts(inp)
	ans := maxInt
	for pos := minSlice(nums); pos <= maxSlice(nums); pos++ {
		total := 0
		for _, start := range nums {
			total += abs(pos - start)
		}
		ans = min(ans, total)
	}
	return ans
}

func part2(filename string) int {
	inp := input.ReadSlice(filename)
	nums := mapToInts(inp)
	ans := maxInt
	for pos := minSlice(nums); pos <= maxSlice(nums); pos++ {
		total := 0
		for _, start := range nums {
			total += getP2Distance(start, pos)
		}
		ans = min(ans, total)
	}
	return ans
}

func getP2Distance(start, end int) int {
	diff := abs(start - end)
	ans, step := 0, 0
	for i := 0; i <= diff; i++ {
		ans += step
		step++
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSlice(s []int) int {
	m := int(maxInt)
	for _, x := range s {
		if x < m {
			m = x
		}
	}
	return m
}

func maxSlice(s []int) int {
	m := 0
	for _, x := range s {
		if x > m {
			m = x
		}
	}
	return m
}

func mapToInts(l []string) []int {
	a := make([]int, len(l))
	for i, s := range l {
		a[i] = stringToInt(s)
	}
	return a
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
