package main

import (
	_ "embed"
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

var oppTurns = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
}

var youTurns = map[string]string{
	"X": "R",
	"Y": "P",
	"Z": "S",
}

var turnScores = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

var resScores = map[string]int{
	"L": 0,
	"D": 3,
	"W": 6,
}

func part1() int {
	turns := utils.GetLines(input)
	sum := 0

	for _, t := range turns {
		turn := strings.Split(t, " ")
		opp, you := turn[0], turn[1]
		oppC := oppTurns[opp]
		youC := youTurns[you]

		if oppC == youC {
			sum += resScores["D"]
			sum += turnScores[youC]
		} else {
			if youC == "R" {
				sum += turnScores["R"]
				if oppC == "P" {
					sum += resScores["L"]
				} else {
					sum += resScores["W"]
				}
			} else if youC == "P" {
				sum += turnScores["P"]
				if oppC == "S" {
					sum += resScores["L"]
				} else {
					sum += resScores["W"]
				}
			} else {
				sum += turnScores["S"]
				if oppC == "R" {
					sum += resScores["L"]
				} else {
					sum += resScores["W"]
				}
			}
		}
	}
	return sum
}

func part2() int {
	turns := utils.GetLines(input)
	sum := 0

	for _, t := range turns {
		turn := strings.Split(t, " ")
		opp, you := turn[0], turn[1]
		oppC := oppTurns[opp]

		if you == "Y" {
			sum += resScores["D"]
			sum += turnScores[oppC]
		} else if you == "X" {
			sum += resScores["L"]
			if oppC == "P" {
				sum += turnScores["R"]
			} else if oppC == "R" {
				sum += turnScores["S"]
			} else {
				sum += turnScores["P"]
			}
		} else {
			sum += resScores["W"]
			if oppC == "P" {
				sum += turnScores["S"]
			} else if oppC == "R" {
				sum += turnScores["P"]
			} else {
				sum += turnScores["R"]
			}
		}
	}
	return sum
}
