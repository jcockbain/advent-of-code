package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"fmt"

	"github.com/jcockbain/advent-of-code-2021/utils"
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

func part1() int {
	lines := utils.GetLines(input)
	s := lines[0]
	for _, newS := range lines[1:] {
		s = transformPair(s, newS)
	}
	return calcMag(s)
}

func part2() int {
	max := 0
	lines := utils.GetLines(input)
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i != j {
				addVal := calcMag(transformPair(lines[i], lines[j]))
				if addVal > max {
					max = addVal
				}
			}
		}
	}
	return max
}

func transformPair(s1, s2 string) string {
	return transform(fmt.Sprintf("[%s,%s]", s1, s2))
}

func transform(s string) string {
	for {
		trans := explode(s)
		if trans != s {
			s = trans
			continue
		}
		spt := split(s)
		if spt == s {
			return s
		}
		s = spt
	}
}

func calcMag(s string) int {
	stack := stringToStack(s)
	for len(stack) > 1 {
		for i, _ := range stack {
			if isPairStack(stack[i : i+3]) {
				newVal := (3 * toInt(stack[i])) + (2 * toInt(stack[i+2]))
				stack = append(append(stack[:i-1], fmt.Sprint(newVal)), stack[i+4:]...)
			}
		}
	}
	return toInt(stack[0])
}

func explode(s string) string {
	stack := stringToStack(s)
	originalStack := make([]string, len(stack))
	nesting := 0
	copy(originalStack, stack)
	for i, c := range stack {
		if c == "[" {
			nesting += 1
		}
		if c == "]" {
			nesting -= 1
		}
		if (nesting >= 5) && isPairStack(stack[i:i+3]) {
			leftNumber := toInt(stack[i])
			rightNumber := toInt(stack[i+2])
			for right := i + 3; right < len(stack); right++ {
				rNum := stack[right]
				if isNumber(rNum) {
					stack[right] = fmt.Sprint(rightNumber + toInt(rNum))
					break
				}
			}
			for left := i - 1; left >= 0; left-- {
				lNum := stack[left]
				if isNumber(lNum) {
					stack[left] = fmt.Sprint(leftNumber + toInt(lNum))
					break
				}
			}
			stack = append(append(stack[:i-1], "0"), stack[i+4:]...)
			break
		}
	}

	res := strings.Builder{}
	for _, f := range stack {
		res.WriteString(f)
	}
	return res.String()
}

func split(s string) string {
	stack := stringToStack(s)
	originalStack := make([]string, len(stack))
	copy(originalStack, stack)
	for i, c := range stack {
		if isNumber(c) && toInt(c) >= 10 {
			lhs, rhs := createSplitPair(toInt(c))
			pair := []string{"[", fmt.Sprint(lhs), ",", fmt.Sprint(rhs), "]"}
			stack = append(append(stack[:i], pair...), originalStack[i+1:]...)
			break
		}
	}
	res := strings.Builder{}
	for _, f := range stack {
		res.WriteString(f)
	}
	return res.String()
}

func stringToStack(s string) []string {
	stack := []string{}
	i := 0
	for i < len(s) {
		c := string(s[i])
		if isNumber(c) {
			newNumber := ""
			for isNumber(string(s[i])) {
				newNumber += string(s[i])
				i++
			}
			stack = append(stack, newNumber)
		} else {
			if c == "[" {
				stack = append(stack, c)
			}
			if c == "]" {
				stack = append(stack, c)
			}
			if c == "," {
				stack = append(stack, ",")
			}
			i++
		}
	}
	return stack
}

func createSplitPair(x int) (int, int) {
	lhs := x / 2
	rhs := int(math.Ceil(float64(x) / float64(2)))
	return lhs, rhs
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isPairStack(s []string) bool {
	return (isNumber(s[0]) && (s[1] == ",") && isNumber(s[2]))
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