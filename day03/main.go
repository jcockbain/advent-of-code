package main

import (
	"math"
	"strconv"
	"time"

	input "github.com/jcockbain/advent-of-code-2021/inpututils"

	"fmt"
)

func main() {
	start := time.Now()
	input := input.GetInputPath()

	fmt.Println("--- Part One ---")
	fmt.Println(Part1(input))
	elapsed := time.Since(start)
	fmt.Printf("%s took %s seconds \n", "Part 1", elapsed)

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2(input))
	fmt.Printf("%s took %s seconds \n", "Part 2", time.Since(start)-elapsed)
}

func Part1(filename string) int {
	lines := input.ReadLines(filename)
	ones := map[int]int{}
	gammaRate, epsilonRate := 0, 0
	for _, line := range lines {
		for pos := len(line) - 1; pos >= 0; pos -= 1 {
			if getInt(string(line[pos])) == 1 {
				ones[len(line)-1-pos] += 1
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

func Part2(filename string) int {
	input := input.ReadLines(filename)

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
				target += 1
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
