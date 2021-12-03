package main

import (
	input "aoc2021/inpututils"
	"math"
	"strconv"
	"time"

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
			c := string(line[pos])
			if getInt(c) == 1 {
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
	lines := make([]string, len(input))
	copy(lines, input)
	oxygenRating, co2Rating := 0, 0
	bitSize := len(input[0])
	for b := 0; b < bitSize; b++ {
		newLines := []string{}
		ones, mostCommon := 0, 0
		for _, l := range lines {
			c := string(l[b])
			ones += getInt(c)
		}
		if ones >= len(lines)-ones {
			mostCommon += 1
		}
		for _, l := range lines {
			c := string(l[b])
			if getInt(c) == mostCommon {
				newLines = append(newLines, l)
			}
		}
		if len(newLines) == 1 {
			oxygenRating = convertBinaryToDecimal(getInt(newLines[0]))
		}
		lines = newLines
	}
	lines = make([]string, len(input))
	copy(lines, input)
	for b := 0; b < bitSize; b++ {
		newLines := []string{}
		ones, leastCommon := 0, 1
		for _, l := range lines {
			c := string(l[b])
			ones += getInt(c)
		}
		if ones >= len(lines)-ones {
			leastCommon -= 1
		}
		for _, l := range lines {
			c := string(l[b])
			if getInt(c) == leastCommon {
				newLines = append(newLines, l)
			}
		}
		if len(newLines) == 1 {
			co2Rating = convertBinaryToDecimal(getInt(newLines[0]))
		}
		lines = newLines
	}
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
