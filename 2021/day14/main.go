package main

import (
	_ "embed"
	"math"
	"regexp"

	"fmt"

	"github.com/jcockbain/advent-of-code-2021/utils"
)

var (
	benchmark = false
	re        = regexp.MustCompile(`(.+) -> (.+)`)
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := part2(10), part2(40)

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part2(steps int) int {
	lines := utils.GetLines(input)
	template := lines[0]
	rules := map[string]string{}
	for _, line := range lines[2:] {
		parts := re.FindStringSubmatch(line)
		rules[parts[1]] = parts[2]
	}
	// these need to be recorded for calculating the character count at the end
	first, last := string(template[0]), string(template[len(template)-1])
	c := getPairsCounter(template)

	for s := 0; s < steps; s++ {
		newC := counter{}
		for k, v := range c {
			newC[string(k[0])+rules[k]] += v
			newC[rules[k]+string(k[1])] += v
		}
		c = newC
	}
	charCounter := getCharCounterFromPairs(c, first, last)
	return charCounter.getMostCommonVal() - charCounter.getLeastCommonVal()
}

type counter map[string]int

func getPairsCounter(s string) counter {
	c := counter{}
	for i := 0; i < len(s)-1; i++ {
		c[s[i:i+2]]++
	}
	return c
}

func getCharCounterFromPairs(c counter, first string, last string) counter {
	m := counter{}
	for k, v := range c {
		for _, char := range k {
			m[string(char)] += v
		}
	}
	for k, v := range m {
		if (k == first) || (k == last) {
			m[k] = v + 1
		}
		m[k] = m[k] / 2
	}
	return m
}

func (c counter) getMostCommonVal() int {
	mv := 0
	for _, v := range c {
		if v > mv {
			mv = v
		}
	}
	return mv
}

func (c counter) getLeastCommonVal() int {
	mv := math.MaxInt
	for _, v := range c {
		if v < mv {
			mv = v
		}
	}
	return mv
}
