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
	for _, line := range lines {
		for pos := len(line) - 1; pos >= 0; pos -= 1 {
			c := string(line[pos])
			val, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			if val == 1 {
				ones[len(line)-1-pos] += 1
			}
		}
	}

	length := len(lines)
	gammaRate, epsilonRate := 0, 0
	for pos, count := range ones {
		if count > (length - count) {
			gammaRate += 1 << pos
		} else {
			epsilonRate += 1 << pos
		}
	}

	return gammaRate * epsilonRate
}

func Part2(filename string) int {
	lines := input.ReadLines(filename)
	oxygenRating, co2Rating := "", ""
	for b := 0; b < len(lines[0]); b++ {
		newLines := []string{}
		ones, mostCommon := 0, 0
		for _, l := range lines {
			c := l[b]
			v, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			ones += v
		}
		if ones >= len(lines)-ones {
			mostCommon += 1
		}
		for _, l := range lines {
			c := l[b]
			v, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			if v == mostCommon {
				newLines = append(newLines, l)
			}
		}
		if len(newLines) == 1 {
			oxygenRating = newLines[0]
		}
		lines = newLines
	}

	lines2 := input.ReadLines(filename)
	for b := 0; b < len(lines[0]); b++ {
		newLines := []string{}
		ones, leastCommon := 0, 1
		for _, l := range lines2 {
			c := l[b]
			v, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			ones += v
		}
		if ones >= len(lines2)-ones {
			leastCommon -= 1
		}
		for _, l := range lines2 {
			c := l[b]
			v, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			if v == leastCommon {
				newLines = append(newLines, l)
			}
		}
		if len(newLines) == 1 {
			co2Rating = newLines[0]
		}
		lines2 = newLines
	}
	oxygenVal := convertBinaryToDecimal(getInt(oxygenRating))
	co2Val := convertBinaryToDecimal(getInt(co2Rating))
	return oxygenVal * co2Val
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
