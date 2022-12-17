package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
	nRe       = regexp.MustCompile(`(\d+,\d+)`)
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

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

type pos struct{ x, y int }

func (p pos) add(p2 pos) pos { return pos{p.x + p2.x, p.y + p2.y} }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getPos(s string) pos {
	c := strings.Split(s, ",")
	x, y := toInt(c[0]), toInt(c[1])
	return pos{x, y}
}

func getPosMap() (map[pos]int, int) {
	lines := utils.GetLines(input)

	// 0 = air
	// 1 = rock
	// 2 = sand

	posMap := map[pos]int{}
	maxY := 0
	for _, l := range lines {
		parts := nRe.FindAllString(l, -1)
		last := getPos(parts[0])
		for _, p := range parts[1:] {
			next := getPos(p)
			for x := min(next.x, last.x); x <= max(next.x, last.x); x++ {
				posMap[pos{x, next.y}] = 1
			}
			for y := min(next.y, last.y); y <= max(next.y, last.y); y++ {
				posMap[pos{next.x, y}] = 1
			}
			if next.y > maxY {
				maxY = next.y
			}
			last = next
		}
	}
	return posMap, maxY
}

func part1() int {
	posMap, maxY := getPosMap()
	sandGrains := 0
	movements := []pos{{0, 1}, {-1, 1}, {1, 1}}
	for true {
		sandPos := pos{500, 0}
		settled := false
		for settled == false {
			if sandPos.y > maxY {
				return sandGrains
			}
			for _, m := range movements {
				finalPos := sandPos.add(m)
				if posMap[finalPos] == 0 {
					sandPos = finalPos
					posMap[sandPos] = 2
					settled = true
					sandGrains++
				}
			}
		}
	}
	panic("error!")
}

func printMap(pm map[pos]int, maxY int) {
	for y := 0; y < maxY+10; y++ {
		sb := strings.Builder{}
		for x := 490; x < 520; x++ {
			val, in := pm[pos{x, y}]
			if !in || val == 0 {
				sb.WriteString(".")
			} else if val == 1 {
				sb.WriteString("#")
			} else {
				sb.WriteString("o")
			}
		}
		fmt.Println(sb.String())
	}
}

func part2() int {
	posMap, maxY := getPosMap()
	floor := maxY + 2
	sandGrains := 0
	movements := []pos{{0, 1}, {-1, 1}, {1, 1}}
	for true {
		sandPos := pos{500, 0}
		settled := false
		for settled == false {
			nextSpot := false
			if sandPos.y == floor-1 {
				posMap[sandPos] = 2
				settled = true
				sandGrains++
			}
			for _, m := range movements {
				finalPos := sandPos.add(m)
				if posMap[finalPos] == 0 {
					sandPos = finalPos
					nextSpot = true
					break
				}
			}

			if nextSpot == false {
				posMap[sandPos] = 2
				settled = true
				sandGrains++
				if sandPos.y == 0 && sandPos.x == 500 {
					return sandGrains
				}
			}
		}
	}
	panic("error!")
}
