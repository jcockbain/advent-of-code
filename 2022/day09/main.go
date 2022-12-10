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

type pos struct{ x, y int }

func (p *pos) add(p2 pos) {
	p.x += p2.x
	p.y += p2.y
}

func (p *pos) getTailPos(tailStart pos) pos {
	dx := p.x - tailStart.x
	dy := p.y - tailStart.y
	if abs(dx) <= 1 && abs(dy) <= 1 {
		return tailStart
	}
	if abs(dx) == 2 && abs(dy) == 2 {
		if dx > 0 && dy > 0 {
			return pos{p.x - 1, p.y - 1}
		}
		if dx > 0 && dy < 0 {
			return pos{p.x - 1, p.y + 1}
		}
		if dx < 0 && dy > 0 {
			return pos{p.x + 1, p.y - 1}
		}
		return pos{p.x + 1, p.y + 1}
	}

	if abs(dx) == 2 {
		if dx > 0 {
			return pos{p.x - 1, p.y}
		}
		return pos{p.x + 1, p.y}
	}
	if abs(dy) == 2 {
		if dy > 0 {
			return pos{p.x, p.y - 1}
		}
		return pos{p.x, p.y + 1}

	}
	panic(fmt.Sprintf("rope %v -> %v is f****d", p, tailStart))
}

func (p *pos) diff(p2 pos) pos {
	return pos{abs(p2.x - p.x), abs(p2.y - p.y)}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

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
	visited := map[pos]struct{}{}
	head, tail := pos{}, pos{}
	for _, line := range lines {
		spl := strings.Split(line, " ")
		direction, distance := spl[0], toInt(spl[1])
		var move pos
		switch direction {
		case "U":
			move = pos{0, 1}
		case "R":
			move = pos{1, 0}
		case "D":
			move = pos{0, -1}
		case "L":
			move = pos{-1, 0}
		}
		for i := 0; i < distance; i++ {
			head.add(move)
			tail = head.getTailPos(tail)
			visited[tail] = struct{}{}
		}

	}
	return len(visited)
}

func part2() int {
	lines := utils.GetLines(input)
	visited := map[pos]struct{}{}
	knots := [10]pos{}
	for _, line := range lines {
		spl := strings.Split(line, " ")
		direction, distance := spl[0], toInt(spl[1])
		var move pos
		switch direction {
		case "U":
			move = pos{0, 1}
		case "R":
			move = pos{1, 0}
		case "D":
			move = pos{0, -1}
		case "L":
			move = pos{-1, 0}
		}
		for i := 0; i < distance; i++ {
			knots[0].add(move)
			for j := 1; j < 10; j++ {
				knots[j] = knots[j-1].getTailPos(knots[j])
			}
			visited[knots[len(knots)-1]] = struct{}{}
		}
	}
	return len(visited)
}

func printKnots(p [10]pos) {
	fmt.Println("<------->")
	set := map[pos]int{}
	for i := 9; i >= 0; i-- {
		k := p[i]
		set[k] = i
	}
	for y := 5; y >= 0; y-- {
		s := strings.Builder{}
		for x := 0; x < 5; x++ {
			if k, in := set[pos{x, y}]; in {
				s.WriteString(strconv.FormatInt(int64(k), 10))
			} else {
				s.WriteString(".")
			}
		}
		fmt.Println(s.String())
	}
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
