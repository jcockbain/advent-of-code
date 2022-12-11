package main

import (
	_ "embed"
	"strconv"
	"strings"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func toInt(s string) int {
	negative := string(s[0]) == "-"
	if negative {
		s = s[1:]
	}
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if negative {
		return -1 * x
	}
	return x
}

func part1() int {
	lines := utils.GetLines(input)
	signals := map[int]int{}
	cycle, reg := 0, 1

	for _, line := range lines {
		cycle++
		signals[cycle] = reg
		if line != "noop" {
			spl := strings.Split(line, " ")
			dist := toInt(spl[1])
			cycle++
			reg += dist
			signals[cycle] = reg
		}
	}

	// part 1
	s := 0
	for x := 20; x <= 220; x += 40 {
		s += x * signals[x-1]
	}
	// part 2
	for y := 0; y < 6; y++ {
		s := strings.Builder{}
		for x := 0; x < 40; x++ {
			idx := (40 * y) + x
			if abs(signals[idx]-x) <= 1 {
				s.WriteString("#")
			} else {
				s.WriteString(".")
			}
		}
		// fmt.Println(s.String())
	}

	return s
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
