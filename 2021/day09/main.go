package main

import (
	_ "embed"
	"sort"
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
	p1, p2 := part1()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type pos struct{ r, c int }

type posMap map[pos]int

func (m posMap) isDeepPoint(p pos) bool {
	neighbourPos := getNeighbourPos(p)
	for _, dp := range neighbourPos {
		if _, in := m[dp]; in {
			if m[p] >= m[dp] {
				return false
			}
		}
	}
	return true
}

func getNeighbourPos(p pos) (res []pos) {
	for _, dp := range []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		res = append(res, pos{p.r + dp.r, p.c + dp.c})
	}
	return
}

func part1() (int, int) {
	lines := utils.GetLines(input)
	m := posMap{}
	for r, line := range lines {
		for c, v := range line {
			m[pos{r, c}] = strToInt(string(v))
		}
	}

	total := 0
	deepPoints := []pos{}
	for pos, val := range m {
		if m.isDeepPoint(pos) {
			total += (val + 1)
			deepPoints = append(deepPoints, pos)
		}
	}
	basins := []int{}
	visited := map[pos]bool{}

	var dfs func(start pos) int

	dfs = func(s pos) int {
		total := 1
		visited[s] = true
		for _, n := range getNeighbourPos(s) {
			if _, in := m[n]; in {
				if _, in := visited[n]; !in {
					if (m[n] != 9) && (m[n] > m[s]) {
						total += dfs(n)
					}
				}
			}
		}
		return total
	}

	for _, s := range deepPoints {
		basins = append(basins, dfs(s))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	return len(deepPoints), product(basins[:3])
}

func product(l []int) int {
	total := 1
	for _, i := range l {
		total *= i
	}
	return total
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
