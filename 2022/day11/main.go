package main

import (
	_ "embed"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"fmt"
)

var (
	benchmark = false
	numRe     = regexp.MustCompile(`\d+`)
	opRe      = regexp.MustCompile(`Operation: new = (old|\d+) (\+|\*) (old|\d+)`)
	testRe    = regexp.MustCompile(`Test: divisible by (\d+)`)
)

type monkey struct {
	items       queue
	operation   func(int) int
	test        func(int) int
	inspections int
}

type monkeySlice []*monkey

func (ms monkeySlice) Len() int { return len(ms) }

func (ms monkeySlice) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (ms monkeySlice) Less(i, j int) bool {
	return ms[i].inspections < ms[j].inspections
}

func (ms monkeySlice) print() {
	for _, m := range ms {
		fmt.Println(*m)
	}
}

type queue []int

func (q *queue) push(x int) {
	*q = append((*q), x)
}

func (q *queue) pop() int {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

//go:embed test1.txt
var input string

func main() {
	p1 := part1()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
	}
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func mapToInt(s []string) []int {
	res := []int{}
	for _, c := range s {
		res = append(res, toInt(c))
	}
	return res
}

func getNumInLine(s string) int {
	n := numRe.FindString(s)
	return toInt(n)
}

func getTest(s1, s2, s3 string) func(int) int {
	split := numRe.FindStringSubmatch(strings.TrimSpace(s1))
	divisibleBy := toInt(split[0])
	v1, v2 := getNumInLine(s2), getNumInLine(s3)
	return func(x int) int {
		if x%divisibleBy == 0 {
			return v1
		}
		return v2
	}
}

func getOp(s string) func(int) int {
	s = strings.TrimSpace(s)
	split := opRe.FindStringSubmatch(s)
	a, op, b := split[1], split[2], split[3]
	var baseFunc func(int, int) int
	if op == "*" {
		baseFunc = func(a, b int) int { return a * b }
	} else {
		baseFunc = func(a, b int) int { return a + b }
	}
	if a == "old" {
		if b == "old" {
			return func(n int) int { return baseFunc(n, n) }
		}
		return func(n int) int { return baseFunc(n, toInt(b)) }
	}
	if b == "old" {
		return func(n int) int { return baseFunc(toInt(a), n) }
	}
	return func(n int) int { return baseFunc(toInt(a), toInt(b)) }
}

func part1() int {
	monkeyStrings := strings.Split(strings.TrimSpace(input), "\n\n")
	monkeys := monkeySlice{}
	for _, ms := range monkeyStrings {
		m := &monkey{}
		lines := strings.Split(ms, "\n")
		startingItems := numRe.FindAllString(lines[1], -1)
		m.items = mapToInt(startingItems)
		m.operation = getOp(lines[2])
		m.test = getTest(lines[3], lines[4], lines[5])
		monkeys = append(monkeys, m)
	}

	var turns int
	for turns < 20 {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				w := m.items.pop()
				m.inspections++
				w = m.operation(w)
				w = w / 3
				targetMonkey := m.test(w)
				monkeys[targetMonkey].items.push(w)
			}
		}
		turns++
	}

	sort.Sort(monkeys)

	return monkeys[len(monkeys)-1].inspections * monkeys[len(monkeys)-2].inspections
}
