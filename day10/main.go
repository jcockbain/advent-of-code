package main

import (
	"time"

	input "github.com/jcockbain/advent-of-code-2021/inpututils"

	"fmt"
)

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

func main() {
	start := time.Now()
	input := input.GetInputPath()

	fmt.Println("--- Part One ---")
	fmt.Println(part1(input))
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 1", elapsed)

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(input))
	fmt.Printf("%s took %s seconds \n", "Part 2", time.Since(start)-elapsed)
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

func (b bracketsMap) isClosingBracket(s string) bool {
	for close, _ := range b {
		if close == s {
			return true
		}
	}
	return false
}

func (b bracketsMap) getClosingBracket(s string) string {
	for close, open := range b {
		if open == s {
			return close
		}
	}
	return ""
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

func part1(filename string) int {
	lines := input.ReadLines(filename)
	total := 0
	for _, l := range lines {
		s := Stack{}

		for _, i := range l {
			c := string(i)
			fmt.Println(s)
			if brackets.isOpeningBracket(c) {
				s.Push(c)
			} else if s.Peek() != brackets[c] {
				total += value[c]
			} else {
				s.Pop()
			}
		}
	}
	return total
}

func part2(filename string) int {
	return 12
}
