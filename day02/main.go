package main

import (
	"time"

	input "github.com/jcockbain/advent-of-code-2021/inpututils"

	"fmt"
	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`(.+) (\d+)`)
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
	lines := input.ReadLines(filename)
	depth, forward := 0, 0
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		w := stringToInt(parts[2])
		switch parts[1] {
		case "forward":
			forward += w
		case "up":
			depth -= w
		case "down":
			depth += w
		}
	}
	return depth * forward
}

func part2(filename string) int {
	lines := input.ReadLines(filename)
	depth, forward, aim := 0, 0, 0
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		w := stringToInt(parts[2])
		switch parts[1] {
		case "forward":
			forward += w
			depth += (aim * w)
		case "up":
			aim -= w
		case "down":
			aim += w
		}
	}
	return depth * forward
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
