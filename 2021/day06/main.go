package main

import (
	_ "embed"

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
	p1 := part1(80)
	p2 := part1(256)

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

// same code used for part one and two
func part1(days int) int {
	fish := map[int]int{}
	for _, c := range strings.Split(input, ",") {
		fish[stringToInt(c)]++
	}

	for d := 0; d < days; d++ {
		fishToAdd := fish[0]
		fish[0] = 0
		for i := 0; i <= 7; i++ {
			fish[i] = fish[i+1]
		}
		fish[8] = fishToAdd
		fish[6] = fish[6] + fishToAdd
	}
	ans := 0
	for _, f := range fish {
		ans += f
	}
	return ans
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
