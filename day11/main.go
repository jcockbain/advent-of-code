package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	input "github.com/jcockbain/advent-of-code-2021/inpututils"

	"fmt"
)

func main() {
	start := time.Now()
	input := input.GetInputPath()
	steps := 10
	if len(os.Args) >= 3 {
		steps = strToInt(os.Args[2])
	}
	fmt.Println("--- Part One ---")
	fmt.Println(part1(input, steps))
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 1", elapsed)

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(input))
	fmt.Printf("%s took %s seconds \n", "Part 2", time.Since(start)-elapsed)
}

type pos struct{ r, c int }

func (p pos) getNeighbours() []pos {
	n := []pos{}
	for _, dr := range []int{-1, 0, 1} {
		for _, dc := range []int{-1, 0, 1} {
			if !((dr == 0) && (dc == 0)) {
				n = append(n, pos{p.r + dr, p.c + dc})
			}
		}
	}
	return n
}

type octopi map[pos]int

func (o octopi) height() (h int) {
	for pos, _ := range o {
		if pos.r > h {
			h = pos.r
		}
	}
	return h + 1
}

func (o octopi) width() (w int) {
	for pos, _ := range o {
		if pos.c > w {
			w = pos.c
		}
	}
	return w + 1
}

func (o octopi) grow() {
	for p, _ := range o {
		o[p]++
	}
}

func (o octopi) flash() (flashes int) {
	for p, v := range o {
		if v >= 10 {
			flashes++
			for _, n := range p.getNeighbours() {
				if o[n] != 0 {
					o[n]++
				}
			}
			o[p] = 0
		}
	}
	return flashes
}

func (o octopi) allZero() bool {
	for _, v := range o {
		if v != 0 {
			return false
		}
	}
	return true
}

func (o octopi) drawMap() {
	for r := 0; r < o.height(); r++ {
		line := []string{}
		for c := 0; c < o.width(); c++ {
			line = append(line, fmt.Sprint(o[pos{r, c}]))
		}
		fmt.Println(strings.Join(line, ""))
	}
}

func part1(filename string, steps int) int {
	lines := input.ReadLines(filename)
	o := octopi{}
	for r, l := range lines {
		for c, v := range l {
			o[pos{r, c}] = strToInt(string(v))
		}
	}
	totalFlashes := 0
	for s := 0; s < steps; s++ {
		fmt.Printf("\n<--Step %s -->\n\n", fmt.Sprint(s))
		o.drawMap()
		o.grow()
		for {
			newFlashes := o.flash()
			totalFlashes += newFlashes
			if newFlashes == 0 {
				break
			}
		}
	}
	return totalFlashes
}

func part2(filename string) int {
	lines := input.ReadLines(filename)
	o := octopi{}
	for r, l := range lines {
		for c, v := range l {
			o[pos{r, c}] = strToInt(string(v))
		}
	}
	totalFlashes := 0
	for s := 0; s < 1000; s++ {
		fmt.Printf("\n<--Step %s -->\n\n", fmt.Sprint(s))
		o.drawMap()
		o.grow()
		for {
			newFlashes := o.flash()
			if newFlashes == 0 {
				if o.allZero() {
					return s + 1
				}
				break
			}
		}
	}
	return totalFlashes
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
