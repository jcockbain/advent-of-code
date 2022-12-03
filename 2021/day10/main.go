package main

import (
	_ "embed"

	"sort"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
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

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *Stack) Peek() string {
	return (*s)[len(*s)-1]
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

type bracketsMap map[string]string

func (b bracketsMap) isOpeningBracket(s string) bool {
	for _, open := range b {
		if open == s {
			return true
		}
	}
	return false
}

var brackets = bracketsMap{
	"]": "[",
	"}": "{",
	")": "(",
	">": "<",
}

var value = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var value2 = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func part1() int {
	lines := utils.GetLines(input)
	total := 0
	for _, l := range lines {
		s := Stack{}
		for _, i := range l {
			c := string(i)
			if brackets.isOpeningBracket(c) {
				s.Push(c)
			} else if s.Peek() != brackets[c] {
				total += value[c]
				break
			} else {
				s.Pop()
			}
		}
	}
	return total
}

func part2() int {
	lines := utils.GetLines(input)
	scores := []int{}
	for _, l := range lines {
		s := Stack{}
		score := 0
		incomplete := true
		for _, i := range l {
			c := string(i)
			if brackets.isOpeningBracket(c) {
				s.Push(c)
			} else if s.Peek() != brackets[c] {
				incomplete = false
				break
			} else {
				s.Pop()
			}
		}
		if incomplete {
			for _, c := range reverse(s) {
				score *= 5
				score += value2[c]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[(len(scores)-1)/2]
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
