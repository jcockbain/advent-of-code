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

var width int
var height int

type pos struct{ x, y int }

type board map[pos]byte

func (b board) drawMap() {
	fmt.Println("<---->")
	fmt.Println(width, height)
	for r := 0; r < height; r++ {
		s := strings.Builder{}
		for c := 0; c < width; c++ {
			p := pos{r, c}
			s.WriteString(string(b[p]))
		}
		fmt.Println(s.String())
	}
}

func (b board) moveEast() (board, bool) {
	moved := false
	newBoard := board{}
	toEmpty := map[pos]bool{}
	toFilled := map[pos]bool{}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if b[pos{r, c}] == '>' {
				p := pos{r, c}
				nextPos := pos{r, (c + 1) % width}
				if b[nextPos] == '.' {
					moved = true
					toEmpty[p] = true
					toFilled[nextPos] = true
				}
			}
		}
	}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			p := pos{r, c}
			if toFilled[p] {
				newBoard[p] = '>'
			} else if toEmpty[p] {
				newBoard[p] = '.'
			} else {
				newBoard[p] = b[p]
			}
		}
	}
	return newBoard, moved
}

func (b board) moveSouth() (board, bool) {
	moved := false
	newBoard := board{}
	toEmpty := map[pos]bool{}
	toFilled := map[pos]bool{}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if b[pos{r, c}] == 'v' {
				p := pos{r, c}
				nextPos := pos{(r + 1) % height, c}
				if b[nextPos] == '.' {
					moved = true
					toEmpty[p] = true
					toFilled[nextPos] = true
				}
			}
		}
	}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			p := pos{r, c}
			if toFilled[p] {
				newBoard[p] = 'v'
			} else if toEmpty[p] {
				newBoard[p] = '.'
			} else {
				newBoard[p] = b[p]
			}
		}
	}
	return newBoard, moved
}

func parse() board {
	lines := utils.GetLines(input)
	b := board{}
	width = len(lines[0])
	height = len(lines)
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			m := lines[r][c]
			b[pos{r, c}] = m
		}
	}
	return b
}

func main() {
	p1 := part1()
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() int {
	b := parse()
	moves := 0
	for moves < 1000 {
		w := b
		w, movedEast := w.moveEast()
		w, movedSouth := w.moveSouth()
		if !(movedSouth || movedEast) {
			break
		}
		b = w
		moves++
	}
	return moves + 1
}

func part2() int {
	return 12
}
