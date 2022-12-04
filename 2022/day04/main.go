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
	re        = regexp.MustCompile(`(\d+-\d+),(\d+-\d+)`)
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

func part1() (sum int) {
	pairs := utils.GetLines(input)
	for _, pair := range pairs {
		parts := re.FindStringSubmatch(pair)
		section1, section2 := newSection(parts[1]), newSection(parts[2])
		if sectionFullyInside(section1, section2) || sectionFullyInside(section2, section1) {
			sum += 1
		}

	}
	return
}

func part2() (sum int) {
	pairs := utils.GetLines(input)
	for _, pair := range pairs {
		parts := re.FindStringSubmatch(pair)
		section1, section2 := newSection(parts[1]), newSection(parts[2])
		if anyOverlap(section1, section2) || anyOverlap(section2, section1) {
			sum += 1
		}

	}
	return
}

func newSection(s string) section {
	pair := strings.Split(s, "-")
	return section{toInt(pair[0]), toInt(pair[1])}
}

type section struct {
	start, stop int
}

func sectionFullyInside(r1, r2 section) bool {
	return (r1.start >= r2.start) && (r1.stop <= r2.stop)
}

func anyOverlap(r1, r2 section) bool {
	return (r2.start <= r1.stop) && (r1.start <= r2.start)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
