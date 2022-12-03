package main

import (
	_ "embed"
	"strconv"

	"github.com/jcockbain/advent-of-code/utils"

	"fmt"
	"strings"
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
	inp := utils.GetLines(input)
	numbersCalled, boards := processInput(inp)

	for _, n := range numbersCalled {
		for _, b := range boards {
			b.processNum(n)
			if b.isComplete() {
				return n * b.getScore()
			}
		}
	}
	panic("no answer!")
}

func part2() int {
	inp := utils.GetLines(input)
	numbersCalled, boards := processInput(inp)

	for _, n := range numbersCalled {
		for _, b := range boards {
			b.processNum(n)
			if allBoardsComplete(boards) {
				return n * b.getScore()
			}
		}
	}
	panic("no answer!")
}

func processInput(input []string) ([]int, []*board) {
	numbersCalled := mapToInt(strings.Split(input[0], ","))
	boards, currentRows := []*board{}, [][]int{}
	for _, inp := range input[2:] {
		if inp == "" {
			boards = append(boards, newBoard(currentRows))
			currentRows = [][]int{}
		} else {
			currentRows = append(currentRows, mapToInt(strings.Fields(inp)))
		}
	}
	return numbersCalled, append(boards, newBoard(currentRows))
}

type board [][]*bingoVal

type bingoVal struct {
	val    int
	called bool
}

func newBoard(rows [][]int) *board {
	b := make(board, len(rows))
	for i := range b {
		b[i] = make([]*bingoVal, len(rows[0]))
	}
	for d, row := range rows {
		for w, val := range row {
			b[d][w] = &bingoVal{val, false}
		}
	}
	return &b
}

func (b board) processNum(i int) {
	for _, row := range b {
		for _, bv := range row {
			if bv.val == i {
				bv.called = true
			}
		}
	}
}

func (b board) isComplete() bool {
	for _, row := range b {
		if allValsCalled(row) {
			return true
		}
	}

	for c := 0; c < len(b[0]); c++ {
		if allValsCalled(b.getColumn(c)) {
			return true
		}
	}

	return false
}

func (b board) getColumn(index int) (c []*bingoVal) {
	for _, row := range b {
		for col, bv := range row {
			if col == index {
				c = append(c, bv)
			}
		}
	}
	return
}

func (b board) getScore() (s int) {
	for _, row := range b {
		for _, bv := range row {
			if !bv.called {
				s += bv.val
			}
		}
	}
	return
}

func allValsCalled(b []*bingoVal) bool {
	for _, val := range b {
		if !val.called {
			return false
		}
	}
	return true
}

func allBoardsComplete(bds []*board) bool {
	for _, b := range bds {
		if !b.isComplete() {
			return false
		}
	}
	return true
}

func mapToInt(vs []string) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = getInt(v)
	}
	return vsm
}

func getInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
