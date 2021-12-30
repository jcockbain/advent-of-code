package main

import (
	_ "embed"
	"strings"

	"fmt"

	"github.com/jcockbain/advent-of-code-2021/utils"
)

//go:embed test1.txt
var input string

var (
	benchmark      = false
	hallway        = 1
	corridorTop    = 2
	corridorBottom = 3
	space          = byte('.')
	allRoutes      = map[[2]pos][]pos{}
	p2Lines        = []string{"  #D#C#B#A#", "  #D#B#A#C#"}
)

const (
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

type pos struct{ r, c int }

func (p pos) add(p2 pos) pos  { return pos{p.r + p2.r, p.c + p2.c} }
func (p pos) inHallway() bool { return p.r == hallway }

func (p pos) inCorridor() bool {
	for _, cp := range corridors {
		if p.c == cp {
			return true
		}
	}
	return false
}

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
		p := pos{r, c}
		if b[p] != a {
			return false
		}
	}
	return true
}

func (b burrow) inOwnCorridor(p pos) bool { return p.c == corridors[b[p]] }

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

func (b burrow) spaceBelow(p pos) bool {
	if p.r == corridorBottom {
		return false
	}
	for r := p.r + 1; r <= corridorBottom; r++ {
		amp := b[pos{r, p.c}]
		if amp == space {
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

func (b burrow) canEnterCorridor(p1 pos, p2 pos) bool {
	a := b[p1]
	if corridors[a] != p2.c {
		return false
	}
	for r := corridorTop; r <= corridorBottom; r++ {
		amp := b[pos{r, p2.c}]
		if (amp != space) && (amp != a) {
			return false
		}
	}
	return true
}

func (b burrow) getPossibleMoves() []move {
	moves := []move{}
	for p1, a := range b {
		if isAmp(a) {
			for p2, a2 := range b {
				if (p1 != p2) && (a2 == space) {
					// rule 1: no stopping above a corridor
					if p2.inCorridor() && p2.inHallway() {
						continue
					}
					// rule 2: must only enter own corridor, with own amps
					if p2.inCorridor() && !b.canEnterCorridor(p1, p2) {
						continue
					}
					// rule 3: no moving within hallways
					if p1.inHallway() && p2.inHallway() {
						continue
					}
					// no reason to move upwards within a corridor
					if p1.c == p2.c && p2.r-p1.r < 0 {
						continue
					}
					// only move out of own corridor if blocking other amp
					if p1.inCorridor() && p1.c == corridors[a] && p2.inHallway() && !b.isBlockingCorridor(p1) {
						continue
					}
					// don't leave space below when entering coridor
					if p2.inCorridor() && p1.inHallway() && p2.r-p1.r > 0 && b.spaceBelow(p2) {
						continue
					}
					// stay at the bottom of own corrdor
					if b.inOwnCorridor(p1) && p1.r == corridorBottom {
						continue
					}
					r := allRoutes[[2]pos{p1, p2}]
					if b.routeIsClear(r) {
						moves = append(moves, move{a, p1, p2, (len(r) - 1) * energies[a]})
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

// bfs to find route
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
	newB := burrow{}
	for pos, a := range b {
		newB[pos] = a
	}
	newB[m.dest] = m.amp
	newB[m.start] = space
	return newB
}

type move struct {
	amp    byte
	start  pos
	dest   pos
	energy int
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
	}
	for c := 1; c <= 11; c++ {
		p := pos{1, c}
		b[p] = '.'
	}
	for _, c := range corridors {
		for r := 2; r <= corridorBottom; r++ {
			p := pos{r, c}
			b[p] = lines[r][c]
		}
	}
	return b
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
