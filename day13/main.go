package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	"fmt"

	"github.com/jcockbain/advent-of-code-2021/utils"
)

var (
	benchmark = false
	re1       = regexp.MustCompile(`(\d+),(\d+)`)
	re2       = regexp.MustCompile(`fold along (x|y)=(\d+)`)
)

//go:embed input.txt
var input string

type pos struct{ x, y int }

func main() {
	p1 := part1()
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2 :\n %s", p2)
	}
}

func part1() int {
	lines := utils.GetLines(input)
	p, folds := loadData(lines)
	f := folds[0]
	if f.axes == "x" {
		p = p.foldAlongX(f.line)
	} else {
		p = p.foldAlongY(f.line)
	}
	return p.countDots()
}

func part2() string {
	lines := utils.GetLines(input)
	p, folds := loadData(lines)
	for _, f := range folds {
		if f.axes == "x" {
			p = p.foldAlongX(f.line)
		} else {
			p = p.foldAlongY(f.line)
		}
	}
	return p.string()
}

func loadData(s []string) (paper, []fold) {
	p, folds := paper{}, []fold{}
	for _, l := range s {
		if re1.MatchString(l) {
			parts := re1.FindStringSubmatch(l)
			p[pos{toInt(parts[1]), toInt(parts[2])}] = true
		}
		if re2.MatchString(l) {
			parts := re2.FindStringSubmatch(l)
			folds = append(folds, fold{toInt(parts[2]), parts[1]})
		}
	}
	return p, folds
}

type paper map[pos]bool

func (p paper) getHeight() (h int) {
	for pos := range p {
		if pos.y > h {
			h = pos.y
		}
	}
	return h + 1
}

func (p paper) getWidth() (w int) {
	for pos := range p {
		if pos.x > w {
			w = pos.x
		}
	}
	return w + 1
}

func (p paper) foldAlongX(line int) paper {
	r := paper{}
	for d := range p {
		r[pos{line - abs(line-d.x), d.y}] = true
	}
	return r
}

func (p paper) foldAlongY(line int) paper {
	r := paper{}
	for d := range p {
		r[pos{d.x, line - abs(line-d.y)}] = true
	}
	return r
}

func (p paper) countDots() (r int) {
	for _, dot := range p {
		if dot {
			r++
		}
	}
	return
}

func (p paper) string() string {
	var b strings.Builder
	for y := 0; y < p.getHeight(); y++ {
		line := []string{}
		for x := 0; x < p.getWidth(); x++ {
			char := "."
			if p[pos{x, y}] {
				char = "#"
			}
			line = append(line, char)
		}
		b.WriteString(fmt.Sprintf("%s\n", strings.Join(line, "")))
	}
	return b.String()
}

type fold struct {
	line int
	axes string
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
