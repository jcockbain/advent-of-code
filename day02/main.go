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
	fmt.Println(Part1(input))
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 1", elapsed)

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2(input))
	fmt.Printf("%s took %s seconds \n", "Part 2", time.Since(start)-elapsed)
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	depth, forward := 0, 0
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		w, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		dir := parts[1]
		if dir == "forward" {
			forward += w
		}
		if dir == "up" {
			depth -= w
		}
		if dir == "down" {
			depth += w
		}
	}
	return depth * forward
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	depth, forward, aim := 0, 0, 0
	for _, line := range lines {
		parts := re.FindStringSubmatch(line)
		w, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		dir := parts[1]
		if dir == "forward" {
			forward += w
			depth += (aim * w)
		}
		if dir == "up" {
			aim -= w
		}
		if dir == "down" {
			aim += w
		}
	}
	return depth * forward
}
