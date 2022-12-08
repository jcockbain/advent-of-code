package main

import (
	_ "embed"
	"strconv"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
)

//go:embed test1.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type pos struct {
	x, y int
}

func part1() int {
	lines := utils.GetLines(input)
	width := len(lines[0])
	height := len(lines)

	treeMap := map[pos]int{}

	bounds := map[pos][]int{}

	sum := 0
	for y, line := range lines {
		for x, c := range line {
			treeMap[pos{y, x}] = toInt(string(c))
			bounds[pos{y, x}] = []int{0, 0, 0, 0}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			t := treeMap[pos{y, x}]
			if y > 0 {
				bounds[pos{y, x}][0] = bounds[pos{y - 1, x}][0] + t
			}
			if x > 0 {
				bounds[pos{y, x}][1] = bounds[pos{y, x - 1}][1] + t
			}
		}
	}

	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			t := treeMap[pos{y, x}]
			if y < height-1 {
				bounds[pos{y, x}][2] = bounds[pos{y + 1, x}][2] + t
			}
			if x < width-1 {
				bounds[pos{y, x + 1}][3] = bounds[pos{y, x + 1}][3] + t
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			hidden := 1
			if y > 0 && treeMap[pos{y, x}] > bounds[pos{y - 0, x}][i] {
				hidden = 0
			}
		}
		sum += hidden
	}

	return sum
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func part2() int {
	return 12
}
