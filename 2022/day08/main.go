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

type pos struct {
	y, x int
}

func part1() int {
	lines := utils.GetLines(input)
	width := len(lines[0])
	height := len(lines)
	treeMap := map[pos]int{}
	seen := map[pos]struct{}{}

	for y, line := range lines {
		for x, c := range line {
			treeMap[pos{y, x}] = toInt(string(c))
		}
	}

	// look down
	for x := 0; x < width; x++ {
		currentMax := -1
		for y := 0; y < height; y++ {
			p := pos{y, x}
			v := treeMap[p]
			if v > currentMax {
				seen[p] = struct{}{}
			}
			currentMax = max(currentMax, v)
		}
	}

	// look up
	for x := 0; x < width; x++ {
		currentMax := -1
		for y := height - 1; y >= 0; y-- {
			p := pos{y, x}
			v := treeMap[p]
			if v > currentMax {
				seen[p] = struct{}{}
			}
			currentMax = max(currentMax, v)
		}
	}

	// look right
	for y := 0; y < height; y++ {
		currentMax := -1
		for x := 0; x < width; x++ {
			p := pos{y, x}
			v := treeMap[p]
			if v > currentMax {
				seen[p] = struct{}{}
			}
			currentMax = max(currentMax, v)
		}
	}

	// look left
	for y := 0; y < height; y++ {
		currentMax := -1
		for x := width - 1; x >= 0; x-- {
			p := pos{y, x}
			v := treeMap[p]
			if v > currentMax {
				seen[p] = struct{}{}
			}
			currentMax = max(currentMax, v)
		}
	}

	return len(seen)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func part2() int {
	lines := utils.GetLines(input)
	treeMap := map[pos]int{}
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x, c := range line {
			treeMap[pos{y, x}] = toInt(string(c))
		}
	}
	currentMax := 0
	for p, _ := range treeMap {
		if p.x == 0 || p.x == width-1 || p.y == 0 || p.y == height-1 {
			continue
		}
		view := getView(treeMap, p, height, width)
		if view > currentMax {
			currentMax = view
		}
	}
	return currentMax
}

func getView(tm map[pos]int, p pos, h, w int) int {
	overallView := 1
	treeH := tm[p]
	// up
	for y := p.y - 1; y >= 0; y-- {
		v := tm[pos{y, p.x}]
		if v >= treeH || y == 0 {
			overallView *= (p.y - y)
			break
		}
	}

	// down
	for y := p.y + 1; y < h; y++ {
		v := tm[pos{y, p.x}]
		if v >= treeH || y == (h-1) {
			overallView *= (y - p.y)
			break
		}
	}

	// right
	for x := p.x + 1; x < w; x++ {
		v := tm[pos{p.y, x}]
		if v >= treeH || x == (w-1) {
			overallView *= (x - p.x)
			break
		}
	}

	// left
	for x := p.x - 1; x >= 0; x-- {
		v := tm[pos{p.y, x}]
		if v >= treeH || x == 0 {
			overallView *= (p.x - x)
			break
		}
	}

	return overallView

}
