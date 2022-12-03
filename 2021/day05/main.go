package main

import (
	_ "embed"

	"regexp"
	"strconv"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	re        = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
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

type pos struct{ x, y int }

func (p pos) getHorizontalsWith(p2 pos) (positions []pos) {
	if p.y != p2.y {
		return
	}
	min, max := getMinMax(p.x, p2.x)
	for x := min; x <= max; x++ {
		positions = append(positions, pos{x, p.y})
	}
	return
}

func (p pos) getVerticalsWith(p2 pos) (positions []pos) {
	if p.x != p2.x {
		return
	}
	min, max := getMinMax(p.y, p2.y)
	for y := min; y <= max; y++ {
		positions = append(positions, pos{p.x, y})
	}
	return
}

func (p pos) getDiagsWith(p2 pos) (positions []pos) {
	if p2.x == p.x {
		return
	}
	grad := (p2.y - p.y) / (p2.x - p.x)
	if abs(grad) != 1 {
		return
	}
	if p2.x > p.x {
		for x, y := p.x, p.y; x <= p2.x; x, y = x+1, y+grad {
			positions = append(positions, pos{x, y})
		}
	} else {
		for x, y := p2.x, p2.y; x <= p.x; x, y = x+1, y+grad {
			positions = append(positions, pos{x, y})
		}
	}
	return
}

type board map[pos]int

func (b board) getPoints() (points int) {
	for _, v := range b {
		if v >= 2 {
			points++
		}
	}
	return
}

func part1() int {
	inp := utils.GetLines(input)
	b := board{}
	for _, line := range inp {
		parts := re.FindStringSubmatch(line)
		x1, y1, x2, y2 := strToInt(parts[1]), strToInt(parts[2]), strToInt(parts[3]), strToInt(parts[4])
		p1, p2 := pos{x1, y1}, pos{x2, y2}
		travelledOver := append(p1.getHorizontalsWith(p2), p1.getVerticalsWith(p2)...)
		for _, p := range travelledOver {
			b[p]++
		}
	}
	return b.getPoints()
}

func part2() int {
	inp := utils.GetLines(input)
	b := board{}
	for _, line := range inp {
		parts := re.FindStringSubmatch(line)
		x1, y1, x2, y2 := strToInt(parts[1]), strToInt(parts[2]), strToInt(parts[3]), strToInt(parts[4])
		p1, p2 := pos{x1, y1}, pos{x2, y2}
		travelledOver := append(append(p1.getHorizontalsWith(p2), p1.getVerticalsWith(p2)...), p1.getDiagsWith(p2)...)
		for _, p := range travelledOver {
			b[p]++
		}
	}
	return b.getPoints()
}

func getMinMax(a int, b int) (int, int) {
	if a >= b {
		return b, a
	}
	return a, b
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
