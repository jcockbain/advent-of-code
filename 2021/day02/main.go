package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	re        = regexp.MustCompile(`(.+) (\d+)`)
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
	lines := utils.GetLines(input)
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

func part2() int {
	lines := utils.GetLines(input)
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
