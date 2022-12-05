package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	"fmt"
)

var (
	benchmark = false
	nRe       = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
)

//go:embed input.txt
var input string

func main() {
	p1 := part1()
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 1: %s\n", p1)
		fmt.Printf("Part 2: %s\n", p2)
	}
}

type crateStack []byte

func (cs *crateStack) size() int {
	return len(*cs)
}

func (cs *crateStack) pop() byte {
	last := cs.top()
	*cs = (*cs)[:cs.size()-1]
	return last
}

func (cs *crateStack) top() byte {
	return (*cs)[cs.size()-1]
}

func (cs *crateStack) push(r byte) {
	*cs = append(*cs, r)
}

type stackMap map[int]*crateStack

func newStackMap(stackList []string) (stackMap, int) {
	stackNs := stackList[len(stackList)-1]
	splitOnWs := strings.Fields(stackNs)
	maxN := toInt(splitOnWs[len(splitOnWs)-1])
	stackMap := map[int]*crateStack{}
	padLength := len(stackNs) + 1
	paddedStackList := make([]string, len(stackList)-1)
	for i, st := range stackList[:len(stackList)-1] {
		paddedStackList[i] = fmt.Sprintf("%-*s", padLength, st)
	}
	for i := 1; i <= maxN; i++ {
		stackMap[i] = &crateStack{}
		idx := 1 + (4 * (i - 1))
		for j := len(stackList) - 2; j >= 0; j-- {
			char := paddedStackList[j][idx]
			if string(paddedStackList[j][idx]) != " " {
				stackMap[i].push(char)
			}
		}
	}
	return stackMap, maxN
}

func part1() string {
	split := strings.Split(input, "\n\n")
	stacks, moves := split[0], split[1]
	stackList := strings.Split(stacks, "\n")
	movesList := strings.Split(moves, "\n")
	stackMap, maxN := newStackMap(stackList)

	for _, move := range movesList {
		parts := nRe.FindStringSubmatch(move)
		if len(parts) < 4 {
			continue
		}
		num, from, to := toInt(parts[1]), toInt(parts[2]), toInt(parts[3])
		fromStack := stackMap[from]
		toStack := stackMap[to]
		for i := 0; i < num; i++ {
			toStack.push(fromStack.pop())
		}
	}

	res := strings.Builder{}
	for i := 1; i <= maxN; i++ {
		st := stackMap[i]
		res.WriteByte(st.top())
	}

	return res.String()
}

func part2() string {
	split := strings.Split(input, "\n\n")
	stacks, moves := split[0], split[1]
	stackList := strings.Split(stacks, "\n")
	movesList := strings.Split(moves, "\n")
	stackMap, maxN := newStackMap(stackList)

	for _, move := range movesList {
		parts := nRe.FindStringSubmatch(move)
		if len(parts) < 4 {
			continue
		}
		num, from, to := toInt(parts[1]), toInt(parts[2]), toInt(parts[3])
		fromStack := stackMap[from]
		toStack := stackMap[to]
		extraStack := &crateStack{}
		for i := 0; i < num; i++ {
			extraStack.push(fromStack.pop())
		}
		for i := 0; i < num; i++ {
			toStack.push(extraStack.pop())
		}
	}

	res := strings.Builder{}
	for i := 1; i <= maxN; i++ {
		st := stackMap[i]
		res.WriteByte(st.top())
	}

	return res.String()
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
