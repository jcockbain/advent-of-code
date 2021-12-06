package main

import (
	"strconv"
	"strings"
	"time"

	input "github.com/jcockbain/advent-of-code-2021/inpututils"

	"fmt"
)

func main() {
	start := time.Now()
	input := input.GetInputPath()
	fmt.Println("--- Part One ---")
	fmt.Println(part1(input, 80))
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 1", elapsed)

	fmt.Println("--- Part Two ---")
	fmt.Println(part1(input, 256))
	fmt.Printf("%s took %s seconds \n", "Part 2", time.Since(start)-elapsed)
}

// same code used for part one and two
func part1(filename string, days int) int {
	inp := input.ReadRaw(filename)
	fish := map[int]int{}
	for _, c := range strings.Split(inp, ",") {
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
