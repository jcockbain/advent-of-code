package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"fmt"

	"container/heap"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
)

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	items []pos
	m     map[pos]int // value to index
	pr    map[pos]int // value to priority
}

func (pq *PriorityQueue) Len() int           { return len(pq.items) }
func (pq *PriorityQueue) Less(i, j int) bool { return pq.pr[pq.items[i]] < pq.pr[pq.items[j]] }
func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.m[pq.items[i]] = i
	pq.m[pq.items[j]] = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(pos)
	pq.m[item] = n
	pq.items = append(pq.items, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.m[item] = -1
	pq.items = old[0 : n-1]
	return item
}

// update modifies the priority of an item in the queue.
func (pq *PriorityQueue) update(item pos, priority int) {
	pq.pr[item] = priority
	heap.Fix(pq, pq.m[item])
}
func (pq *PriorityQueue) addWithPriority(item pos, priority int) {
	heap.Push(pq, item)
	pq.update(item, priority)
}

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

type board map[pos]int

func getNeighbours(p pos, h int, w int) []pos {
	n := []pos{}
	for _, dp := range []pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		np := pos{p.x + dp.x, p.y + dp.y}
		if (np.x >= 0) && (np.x < w) && (np.y >= 0) && (np.y < h) {
			n = append(n, np)
		}
	}
	return n
}

func (b board) getMinPathDijkstra(start pos, target pos) (int, map[pos]pos) {
	dist := map[pos]int{}
	prev := map[pos]pos{}
	q := &PriorityQueue{
		items: make([]pos, 0, len(b)),
		m:     make(map[pos]int, len(b)),
		pr:    make(map[pos]int, len(b)),
	}
	for p, _ := range b {
		if p != start {
			dist[p] = math.MaxInt
		}
		q.addWithPriority(p, dist[p])
	}
	for len(q.items) != 0 {
		u := heap.Pop(q).(pos)
		for _, v := range getNeighbours(u, target.y+1, target.x+1) {
			alt := dist[u] + b[v]
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				q.update(v, alt)
			}
		}
	}
	return dist[target], prev
}

func part1() int {
	lines := utils.GetLines(input)
	b := board{}
	height := len(lines)
	width := len(lines[0])
	for y, l := range lines {
		for x, c := range l {
			b[pos{x, y}] = toInt(string(c))
		}
	}
	dist, _ := b.getMinPathDijkstra(pos{0, 0}, pos{width - 1, height - 1})
	// printPath(pos{0, 0}, pos{width - 1, height - 1}, path)
	return dist
}

func part2() int {
	lines := utils.GetLines(input)
	b := board{}
	height := len(lines) * 5
	width := len(lines[0]) * 5
	for sy := 0; sy < 5; sy++ {
		for sx := 0; sx < 5; sx++ {
			for y, l := range lines {
				for x, c := range l {
					y2 := y + (sy * len(lines))
					x2 := x + (sx * len(lines[0]))
					b[pos{x2, y2}] = getWrapped(toInt(string(c)) + sx + sy)
				}
			}
		}
	}
	dist, _ := b.getMinPathDijkstra(pos{0, 0}, pos{width - 1, height - 1})
	// printPath(pos{0, 0}, pos{width - 1, height - 1}, path)
	return dist
}

func printPath(start pos, end pos, prev map[pos]pos) {
	path := []pos{}
	pathsVisited := map[pos]bool{}
	for end != start {
		path = append(path, end)
		end = prev[end]
	}
	for _, p := range path {
		pathsVisited[p] = true
	}
	for y := 0; y < path[0].y+1; y++ {
		line := []string{}
		for x := 0; x < path[0].x+1; x++ {
			if _, in := pathsVisited[pos{x, y}]; in {
				line = append(line, "#")
			} else {
				line = append(line, ".")
			}
		}
		fmt.Println(strings.Join(line, ""))
	}
}

func getWrapped(c int) int {
	if c == 9 {
		return c
	}
	return c % 9
}

func toInt(x string) int {
	i, err := strconv.Atoi(x)
	check(err)
	return i
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
