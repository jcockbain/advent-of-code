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

const (
	WinPoints  = 6
	DrawPoints = 3
	LossPoints = 0
)

//go:embed input.txt
var input string

type weapon interface {
	score() int
	fight(w weapon) int
	// p2 - takes the char for the result and returns a weapon that would produce that res AGAINST this weapon
	forceResult(inst string) weapon
}

func newWeapon(s string) weapon {
	switch s {
	case "R":
		return rock{}
	case "P":
		return paper{}
	case "S":
		return scissors{}
	}
	panic("invalid weapon")
}

type rock struct{}

func (r rock) score() int { return 1 }

func (r rock) fight(w weapon) int {
	switch w.(type) {
	case paper:
		return LossPoints
	case scissors:
		return WinPoints
	default:
		return DrawPoints
	}
}

func (r rock) forceResult(inst string) weapon {
	switch inst {
	case "Z":
		return paper{}
	case "Y":
		return rock{}
	case "X":
		return scissors{}
	}
	panic("invalid instruction")
}

type paper struct{}

func (p paper) score() int { return 2 }

func (p paper) fight(w weapon) int {
	switch w.(type) {
	case paper:
		return DrawPoints
	case scissors:
		return LossPoints
	default:
		return WinPoints
	}
}

func (p paper) forceResult(inst string) weapon {
	switch inst {
	case "Z":
		return scissors{}
	case "Y":
		return paper{}
	case "X":
		return rock{}
	}
	panic("invalid instruction")
}

type scissors struct{}

func (s scissors) score() int { return 3 }

func (s scissors) fight(w weapon) int {
	switch w.(type) {
	case paper:
		return WinPoints
	case scissors:
		return DrawPoints
	default:
		return LossPoints
	}
}

func (s scissors) forceResult(inst string) weapon {
	switch inst {
	case "Z":
		return rock{}
	case "Y":
		return scissors{}
	case "X":
		return paper{}
	}
	panic("invalid instruction")
}

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

func part1() (sum int) {
	turns := utils.GetLines(input)
	for _, t := range turns {
		turn := strings.Split(t, " ")
		opp, you := turn[0], turn[1]
		oppWeapon, youWeapon := newWeapon(oppTurns[opp]), newWeapon(youTurns[you])
		sum += youWeapon.score() + youWeapon.fight(oppWeapon)
	}
	return
}

func part2() (sum int) {
	turns := utils.GetLines(input)
	for _, t := range turns {
		turn := strings.Split(t, " ")
		opp, instruction := turn[0], turn[1]
		oppWeapon := newWeapon(oppTurns[opp])
		youWeapon := oppWeapon.forceResult(instruction)
		sum += youWeapon.score() + youWeapon.fight(oppWeapon)
	}
	return
}
