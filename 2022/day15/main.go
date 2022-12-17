package main

import (
	_ "embed"
	"regexp"

	"fmt"
	"strconv"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
	nRe       = regexp.MustCompile(`(-?\d+)`)
)

//go:embed test1.txt
var input string

const (
	P1TESTY = 10
	P1Y     = 2000000
)

func main() {
	p1 := part1()
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type pos struct {
	x, y int
}

type beacon pos

type sensor struct {
	closestBeacon pos
	p             pos
	mh            int
}

func toInt(s string) int {
	negative := false
	if string(s[0]) == "-" {
		negative = true
		s = s[1:]
	}
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	if negative {
		return -1 * x
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func getManhatten(p1, p2 pos) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func part1() int {
	lines := utils.GetLines(input)

	sensors := []sensor{}

	knownBeaconSet := map[pos]struct{}{}

	for _, l := range lines {
		s := sensor{}
		parts := nRe.FindAllString(l, -1)
		sx, sy, bx, by := toInt(parts[0]), toInt(parts[1]), toInt(parts[2]), toInt(parts[3])
		s.p = pos{sx, sy}
		s.closestBeacon.x = bx
		s.closestBeacon.y = by
		knownBeaconSet[s.closestBeacon] = struct{}{}
		s.mh = getManhatten(s.p, s.closestBeacon)
		sensors = append(sensors, s)
	}

	noBeaconsPossible := map[pos]struct{}{}
	filteredSensors := []sensor{}
	y := P1Y
	// filter only possible sensors
	for _, s := range sensors {
		if abs(s.p.y-P1Y) < s.mh {
			filteredSensors = append(filteredSensors, s)
		}
	}
	for _, s := range filteredSensors {
		for x := s.p.x - mh; x < s.p.x+mh; x++ {
			p := pos{x, y}
			if _, in := knownBeaconSet[p]; !in && getManhatten(s.p, p) <= s.mh {
				noBeaconsPossible[p] = struct{}{}
			}
		}
	}
	return len(noBeaconsPossible)
}

func part2() int {
	lines := utils.GetLines(input)
	sensors := []sensor{}
	knownBeaconSet := map[pos]struct{}{}

	for _, l := range lines {
		s := sensor{}
		parts := nRe.FindAllString(l, -1)
		sx, sy, bx, by := toInt(parts[0]), toInt(parts[1]), toInt(parts[2]), toInt(parts[3])
		s.p = pos{sx, sy}
		s.closestBeacon.x = bx
		s.closestBeacon.y = by
		knownBeaconSet[s.closestBeacon] = struct{}{}
		s.mh = getManhatten(s.p, s.closestBeacon)
		sensors = append(sensors, s)
	}
	possibleBc := map[pos]struct{}{}
	filteredSensors := []sensor{}
	maxY := 20
	maxX := 20
	// filter only possible sensors - x and y
	for _, s := range sensors {
		if abs(s.p.x-maxX) < s.mh && abs(s.p.y-maxY) < s.mh {
			filteredSensors = append(filteredSensors, s)
		}
	}

	fmt.Println(filteredSensors)

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			possibleBc[pos{x, y}] = struct{}{}
		}
	}

	for _, s := range filteredSensors {
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				p := pos{x, y}
				if getManhatten(s.p, p) <= s.mh {
					delete(possibleBc, p)
				}
			}
		}
	}
	fmt.Println(possibleBc)
	return 0

}
