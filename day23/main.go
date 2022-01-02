package main

import (
	_ "embed"
	"strings"

	"fmt"

	"github.com/jcockbain/advent-of-code-2021/utils"
)

//go:embed input.txt
var input string

var (
	benchmark      = false
	corridorTop    = 2
	corridorBottom = 3
	allRoutes      = map[[2]pos][]pos{}
	p2Lines        = []string{"  #D#C#B#A#", "  #D#B#A#C#"}
)

const (
	space  = byte('.')
	amber  = byte('A')
	bronze = byte('B')
	copper = byte('C')
	desert = byte('D')
)

var corridors = map[byte]int{
	amber:  3,
	bronze: 5,
	copper: 7,
	desert: 9,
}
var energies = map[byte]int{
	amber:  1,
	bronze: 10,
	copper: 100,
	desert: 1000,
}

// check order from middle first, as more likely to move amps there
var validHallwayPos = []pos{{1, 6}, {1, 4}, {1, 8}, {1, 2}, {1, 1}, {1, 10}, {1, 11}}

type pos struct{ r, c int }

func (p pos) add(p2 pos) pos { return pos{p.r + p2.r, p.c + p2.c} }

type burrow map[pos]byte

func (b burrow) drawMap() {
	fmt.Println("<----->")
	for r := 0; r < corridorBottom+2; r++ {
		s := strings.Builder{}
		for c := 0; c <= 12; c++ {
			if _, ok := b[pos{r, c}]; ok {
				s.WriteByte(b[pos{r, c}])
			} else {
				s.WriteString("#")
			}
		}
		fmt.Println(s.String())
	}
}

func (b burrow) in(p pos) bool {
	if _, ok := b[p]; ok {
		return true
	}
	return false
}

func (b burrow) lineComplete(a byte) bool {
	c := corridors[a]
	for r := corridorTop; r <= corridorBottom; r++ {
		if b[pos{r, c}] != a {
			return false
		}
	}
	return true
}

func (b burrow) isSolution() bool {
	for a := range corridors {
		if !b.lineComplete(a) {
			return false
		}
	}
	return true
}

func (b burrow) isBlockingCorridor(p pos) bool {
	a := b[p]
	for r := p.r; r <= corridorBottom; r++ {
		amp := b[pos{r, p.c}]
		if (amp != space) && (amp != a) {
			return true
		}
	}
	return false
}

func (b burrow) routeIsClear(p []pos) bool {
	for _, p := range p[1:] {
		if b[p] != space {
			return false
		}
	}
	return true
}

// return if can enter corridor, and highest available pos
func (b burrow) canEnterCorridor(a byte) (bool, pos) {
	c := corridors[a]
	if b[pos{corridorTop, c}] != space {
		return false, pos{0, 0}
	}
	var topSpace pos
	for r := corridorTop; r <= corridorBottom; r++ {
		amp := b[pos{r, c}]
		if amp == space {
			topSpace = pos{r, c}
		} else if amp != a {
			return false, pos{0, 0}
		}
	}
	return true, topSpace
}

// return first amp from top of corridor (and bool for whether one is found)
func (b burrow) getFirstCorridorAmp(c int) (pos, bool) {
	for r := corridorTop; r <= corridorBottom; r++ {
		a := b[pos{r, c}]
		if isAmp(a) {
			return pos{r, c}, true
		}
	}
	return pos{0, 0}, false
}

func (b burrow) getPossibleMoves() []move {
	moves := []move{}
	// move from hallway
	for _, p := range validHallwayPos {
		a := b[p]
		if isAmp(a) {
			canMove, dest := b.canEnterCorridor(a)
			if canMove {
				r := allRoutes[[2]pos{p, dest}]
				if b.routeIsClear(r) {
					moves = append(moves, move{
						a,
						p,
						dest,
						(len(r) - 1) * energies[a],
					})
				}
			}
		}
	}

	// move into hallway
	for a, c := range corridors {
		p, hasAmp := b.getFirstCorridorAmp(c)
		if hasAmp {
			amp := b[p]
			if !(amp == a && !b.isBlockingCorridor(p)) {
				for _, dest := range validHallwayPos {
					r := allRoutes[[2]pos{p, dest}]
					if b.routeIsClear(r) {
						moves = append(moves, move{
							amp,
							p,
							dest,
							(len(r) - 1) * energies[amp],
						})
					}
				}
			}
		}
	}

	return moves
}

type queue [][]pos

func (q *queue) pop() []pos {
	queueItem := (*q)[0]
	(*q) = (*q)[1:len(*q)]
	return queueItem
}

func (q *queue) push(i []pos) {
	(*q) = append(*q, i)
}

type posSet map[pos]bool

func (ps posSet) in(p pos) bool {
	if _, in := ps[p]; in {
		return true
	}
	return false
}

// bfs to find routes between positions in map
func (b burrow) getRoute(p1 pos, p2 pos) []pos {
	q := queue{[]pos{p1}}
	visited := posSet{}
	for len(q) > 0 {
		path := q.pop()
		i := path[len(path)-1]
		visited[i] = true
		if i == p2 {
			return path
		}
		for _, p2 := range []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			newPos := i.add(p2)
			if b.in(newPos) && !visited.in(newPos) {
				newPath := make([]pos, len(path))
				copy(newPath, path)
				newPath = append(newPath, newPos)
				q.push(newPath)
			}
		}
	}
	panic("no route!")
}

// todo? test using this as only data structure format and whether that has efficiency gains
func (b burrow) toCacheKey() string {
	s := strings.Builder{}
	for c := 1; c <= 11; c++ {
		p := pos{1, c}
		s.WriteByte(b[p])
	}
	for _, c := range corridors {
		for r := 2; r <= corridorBottom; r++ {
			p := pos{r, c}
			s.WriteByte(b[p])
		}
	}
	return s.String()
}

func isAmp(b byte) bool {
	if _, in := corridors[b]; in {
		return true
	}
	return false
}

func (b burrow) move(m move) burrow {
	b[m.dest] = m.amp
	b[m.start] = space
	return b
}

type move struct {
	amp    byte
	start  pos
	dest   pos
	energy int
}

func (m move) reverse() move {
	startCp := m.start
	m.start = m.dest
	m.dest = startCp
	return m
}

func main() {
	p1 := getMinEnergy(false)
	p2 := getMinEnergy(true)

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type cache map[string]int

func (c cache) in(k string) bool {
	if _, in := c[k]; in {
		return true
	}
	return false
}

func getMinEnergy(p2 bool) int {
	b := parseBurrow(p2)
	parseRoutes(b)
	c := cache{}
	var findMinEnergyFromConfig func(burrow) int
	findMinEnergyFromConfig = func(b burrow) int {
		if b.isSolution() {
			return 0
		}
		cacheKey := b.toCacheKey()
		if !c.in(cacheKey) {
			energies := []int{}
			for _, mv := range b.getPossibleMoves() {
				energies = append(energies, mv.energy+findMinEnergyFromConfig(b.move(mv)))
				mv = mv.reverse()
				b.move(mv)
			}
			if len(energies) == 0 {
				// set as a large number to mark as impossible configuration
				c[cacheKey] = 1000000
			} else {
				c[cacheKey] = minSlice(energies)
			}
		}
		return c[cacheKey]
	}
	return findMinEnergyFromConfig(b)
}

func minSlice(s []int) int {
	min := s[0]
	for _, x := range s[1:] {
		if x < min {
			min = x
		}
	}
	return min
}

func parseBurrow(p2 bool) burrow {
	b := burrow{}
	lines := utils.GetLines(input)
	if p2 {
		newLines := []string{}
		newLines = append(newLines, lines[:3]...)
		newLines = append(newLines, p2Lines...)
		lines = append(newLines, lines[3:]...)
		corridorBottom = 5
	} else {
		corridorBottom = 3
	}
	for _, p := range getHallwayPos() {
		b[p] = '.'
	}
	for _, c := range corridors {
		for r := corridorTop; r <= corridorBottom; r++ {
			p := pos{r, c}
			b[p] = lines[r][c]
		}
	}
	return b
}

func getHallwayPos() (res []pos) {
	for c := 1; c <= 11; c++ {
		p := pos{1, c}
		res = append(res, p)
	}
	return res
}

func parseRoutes(b burrow) {
	for p1 := range b {
		for p2 := range b {
			if p2 != p1 {
				allRoutes[[2]pos{p1, p2}] = b.getRoute(p1, p2)
			}
		}
	}
}
