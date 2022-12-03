package main

import (
	_ "embed"
	"regexp"
	"strconv"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
	re        = regexp.MustCompile(`--- scanner (\d+) ---`)
	lineRe    = regexp.MustCompile(`(-?\d+),(-?\d+),(-?\d+)`)
)

//go:embed input.txt
var input string
var orMap = getOrMap()

func main() {
	p1, p2 := part1()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type pos struct{ x, y, z int }
type orientation [3]int

func (p pos) subtract(p2 pos) pos { return pos{p.x - p2.x, p.y - p2.y, p.z - p2.z} }
func (p pos) mag() int            { return abs(p.x) + abs(p.y) + abs(p.z) }

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type scanner struct {
	number          int
	absoluteBeacons []pos
	absolutePos     pos
}

func part1() (int, int) {
	lines := utils.GetLines(input)
	scs := []scanner{}
	currentScanner := scanner{}
	for _, line := range lines[1:] {
		if re.MatchString(line) {
			parts := re.FindStringSubmatch(line)
			scs = append(scs, currentScanner)
			currentScanner = scanner{number: toInt(parts[1])}
		} else if line != "" {
			p := lineRe.FindStringSubmatch(line)
			x, y, z := toInt(p[1]), toInt(p[2]), toInt(p[3])
			currentScanner.absoluteBeacons = append(currentScanner.absoluteBeacons, pos{x, y, z})
		}
	}
	scs = append(scs, currentScanner)
	fixedScanners := []scanner{scs[0]}
	unfixedScanners := scs[1:]

	// keep a set of fixed -> unfixed to avoid dupicating work
	tried := map[[2]int]struct{}{}
	for updated := true; updated && len(unfixedScanners) > 0; {
		updated = false
		for _, scA := range fixedScanners {
			for i := 0; i < len(unfixedScanners); i++ {
				unfixed := unfixedScanners[i]
				if _, in := tried[[2]int{scA.number, unfixed.number}]; !in {
					updatedSc, isUpdated := locateScanner(scA, unfixed)
					if isUpdated {
						updated = true
						fixedScanners = append(fixedScanners, updatedSc)
						unfixedScanners = append(unfixedScanners[:i], unfixedScanners[i+1:]...)
						i--
					}
					tried[[2]int{scA.number, unfixed.number}] = struct{}{}
				}
			}
		}
	}
	positions := map[pos]bool{}
	for _, u := range fixedScanners {
		for _, p := range u.absoluteBeacons {
			positions[p] = true
		}
	}
	maxDist := 0
	for _, a := range fixedScanners {
		for _, b := range fixedScanners {
			dist := a.absolutePos.subtract(b.absolutePos).mag()
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return len(positions), maxDist
}

func locateScanner(fixed, unfixed scanner) (scanner, bool) {
	for _, fixedBcn := range fixed.absoluteBeacons {
		for _, unfixedBeacon := range unfixed.absoluteBeacons {
			for or := range orMap {
				unfixedPos := fixedBcn.subtract(orMap[or](unfixedBeacon))
				positionsFromFixed := getBcnPositionsFromScanner(unfixed.absoluteBeacons, unfixedPos, or)
				commonBeacons := getCommonBcnPositions(positionsFromFixed, fixed.absoluteBeacons)
				if commonBeacons >= 12 {
					unfixed.absoluteBeacons = positionsFromFixed
					unfixed.absolutePos = unfixedPos
					return unfixed, true
				}
			}
		}
	}
	return unfixed, false
}

func getBcnPositionsFromScanner(bcs []pos, rp pos, or orientation) []pos {
	res := make([]pos, len(bcs))
	for i, b := range bcs {
		new := orMap[or](b)
		res[i] = pos{rp.x + new.x, rp.y + new.y, rp.z + new.z}
	}
	return res
}

func getCommonBcnPositions(bc1 []pos, bc2 []pos) int {
	var res int
	for _, b1 := range bc1 {
		for _, b2 := range bc2 {
			if b1 == b2 {
				res++
			}
		}
	}
	return res
}

type posMap func(pos) pos

// horrible way to gen rotations
func getOrMap() map[orientation]posMap {
	return map[orientation]posMap{
		{1, 2, 3}:    func(p pos) pos { return pos{p.x, p.y, p.z} },
		{1, -2, -3}:  func(p pos) pos { return pos{p.x, -p.y, -p.z} },
		{1, 3, -2}:   func(p pos) pos { return pos{p.x, p.z, -p.y} },
		{1, -3, 2}:   func(p pos) pos { return pos{p.x, -p.z, p.y} },
		{-1, -2, 3}:  func(p pos) pos { return pos{-p.x, -p.y, p.z} },
		{-1, -3, -2}: func(p pos) pos { return pos{-p.x, -p.z, -p.y} },
		{-1, 3, 2}:   func(p pos) pos { return pos{-p.x, p.z, p.y} },
		{-1, 2, -3}:  func(p pos) pos { return pos{-p.x, p.y, -p.z} },
		{2, -1, 3}:   func(p pos) pos { return pos{p.y, -p.x, p.z} },
		{2, 1, -3}:   func(p pos) pos { return pos{p.y, p.x, -p.z} },
		{2, 3, 1}:    func(p pos) pos { return pos{p.y, p.z, p.x} },
		{2, -3, -1}:  func(p pos) pos { return pos{p.y, -p.z, -p.x} },
		{-2, 3, -1}:  func(p pos) pos { return pos{-p.y, p.z, -p.x} },
		{-2, -3, 1}:  func(p pos) pos { return pos{-p.y, -p.z, p.x} },
		{-2, 1, 3}:   func(p pos) pos { return pos{-p.y, p.x, p.z} },
		{-2, -1, -3}: func(p pos) pos { return pos{-p.y, -p.x, -p.z} },
		{3, -1, -2}:  func(p pos) pos { return pos{p.z, -p.x, -p.z} },
		{3, -2, 1}:   func(p pos) pos { return pos{p.z, -p.y, p.x} },
		{3, 1, 2}:    func(p pos) pos { return pos{p.z, p.x, p.y} },
		{3, 2, -1}:   func(p pos) pos { return pos{p.z, p.y, -p.x} },
		{-3, 1, -2}:  func(p pos) pos { return pos{-p.z, p.x, -p.y} },
		{-3, 2, 1}:   func(p pos) pos { return pos{-p.z, p.y, p.x} },
		{-3, -2, -1}: func(p pos) pos { return pos{-p.z, -p.y, -p.x} },
		{-3, -1, 2}:  func(p pos) pos { return pos{-p.z, -p.x, p.y} },
	}
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
