package main

import (
	_ "embed"

	"math"
	"strconv"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var benchmark = false

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
	ones := map[int]int{}
	gammaRate, epsilonRate := 0, 0
	for _, line := range lines {
		for pos := len(line) - 1; pos >= 0; pos-- {
			if getInt(string(line[pos])) == 1 {
				ones[len(line)-1-pos]++
			}
		}
	}
	for pos, count := range ones {
		if count > (len(lines) - count) {
			gammaRate += 1 << pos
		} else {
			epsilonRate += 1 << pos
		}
	}

	return gammaRate * epsilonRate
}

func part2() int {
	input := utils.GetLines(input)

	getRating := func(compFunc func(int, int) bool) int {
		lines := make([]string, len(input))
		copy(lines, input)
		for b := 0; b < len(input[0]); b++ {
			newLines := []string{}
			ones, target := 0, 0
			for _, l := range lines {
				ones += getInt(string(l[b]))
			}
			if compFunc(ones, len(lines)-ones) {
				target++
			}
			for _, l := range lines {
				if getInt(string(l[b])) == target {
					newLines = append(newLines, l)
				}
			}
			if len(newLines) == 1 {
				return convertBinaryToDecimal(getInt(newLines[0]))
			}
			lines = newLines
		}
		panic("no answer found")
	}

	oxygenRating := getRating(func(i1, i2 int) bool { return i1 >= i2 })
	co2Rating := getRating(func(i1, i2 int) bool { return i1 < i2 })
	return oxygenRating * co2Rating
}

func getInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func convertBinaryToDecimal(number int) int {
	decimal := 0
	counter := 0.0
	remainder := 0

	for number != 0 {
		remainder = number % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		number = number / 10
		counter++
	}
	return decimal
}
