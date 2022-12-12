package main

import (
	_ "embed"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type pos struct{ x, y int }

type qItem struct {
	p    pos
	dist int
}

func solve() (int, int) {
	lines := utils.GetLines(input)

	height := len(lines)
	width := len(lines[0])

	var start pos
	var end pos
	heightMap := map[pos]int{}

	for y, l := range lines {
		for x, c := range l {
			p := pos{x, y}
			if string(c) == "S" {
				start = p
				heightMap[p] = 0
			} else if string(c) == "E" {
				end = p
				heightMap[p] = int('z') - int('a')
			} else {
				heightMap[p] = int(c) - int('a')
			}
		}
	}

	var bfs func(pos) int
	bfs = func(start pos) int {
		queue := []qItem{{start, 0}}
		visited := map[pos]struct{}{}

		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			if _, in := visited[c.p]; in {
				continue
			}
			if c.p == end {
				return c.dist
			}
			visited[c.p] = struct{}{}
			for _, dc := range []pos{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
				n := pos{c.p.x + dc.x, c.p.y + dc.y}
				if 0 <= n.x && n.x < width && n.y >= 0 && n.y < height {
					if heightMap[n] <= heightMap[c.p]+1 {
						queue = append(queue, qItem{n, c.dist + 1})
					}
				}
			}
		}
		return 10000000
	}
	p1 := bfs(start)
	p2 := p1
	for k, v := range heightMap {
		if v == 0 && k != start {
			b := bfs(k)
			if b < p2 {
				p2 = b
			}
		}
	}
	return p1, p2
}
