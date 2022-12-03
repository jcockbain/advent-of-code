package main

import (
	_ "embed"

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

func part1() (sum int) {
	bags := utils.GetLines(input)
	for _, bag := range bags {
		sum += getSharedValue(bag)
	}
	return
}

func getSharedValue(bag string) int {
	seenSet := map[rune]struct{}{}
	bag1 := bag[0 : len(bag)/2]
	bag2 := bag[len(bag)/2 : len(bag)]
	for _, c := range bag1 {
		seenSet[c] = struct{}{}
	}
	for _, c := range bag2 {
		if _, in := seenSet[c]; in {
			return getRuneVal(c)
		}
	}
	panic(fmt.Sprintf("invalid bag %v", bag))
}

func part2() (sum int) {
	bags := utils.GetLines(input)
	for i := 0; i < len(bags); i += 3 {
		sum += getCommonItem(bags[i : i+3])
	}
	return
}

func getCommonItem(bags []string) int {
	counter := map[rune]int{}
	for _, s := range bags {
		// so we only count object once per bag
		seen := map[rune]struct{}{}
		for _, c := range s {
			if _, ok := seen[c]; !ok {
				counter[c] += 1
				seen[c] = struct{}{}
			}
		}
	}
	for c, count := range counter {
		if count == 3 {
			return getRuneVal(c)
		}
	}
	panic(fmt.Sprintf("invalid bags %v", bags))
}

func getRuneVal(c rune) int {
	if int(c) >= int('a') {
		return 1 + int(c) - int('a')
	}
	return 27 + int(c) - int('A')
}
