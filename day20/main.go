package main

import (
	_ "embed"
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

type pos struct{ r, c int }

func (p pos) getNeighbours() []pos {
	n := []pos{}
	for _, dr := range []int{-1, 0, 1} {
		for _, dc := range []int{-1, 0, 1} {
			n = append(n, pos{p.r + dr, p.c + dc})
		}
	}
	return n
}

type image map[pos]bool

func (i image) drawImg() {
	minW, maxW, minH, maxH := i.getDims()
	fmt.Println("<--Image-->")
	for r := minH - 3; r <= maxH+3; r++ {
		s := strings.Builder{}
		for c := minW - 3; c <= maxW+3; c++ {
			if i[pos{r, c}] {
				s.WriteString("#")
			} else {
				s.WriteString(".")
			}
		}
		fmt.Println(s.String())
	}
}

func (i image) getDims() (int, int, int, int) {
	minH, maxH := 0, 0
	minW, maxW := 0, 0
	for pos := range i {
		if pos.r < minH {
			minH = pos.r
		}
		if pos.r > maxH {
			maxH = pos.r
		}
		if pos.c < minW {
			minW = pos.c
		}
		if pos.c > maxW {
			maxW = pos.c
		}
	}
	return minW, maxW, minH, maxH
}

type dimensions struct{ minW, maxW, minH, maxH int }

func isGreaterThanBorder(d dimensions, p pos) bool {
	return (p.r < d.minH) || (p.r > d.maxH) || (p.c < d.minW) || (p.c > d.maxW)
}

func part1() int {
	lines := utils.GetLines(input)
	boardInp := lines[2:]
	img := image{}
	algo := lines[0]

	toggle := false
	border := false
	if (algo[0] == byte('#')) && (algo[len(algo)-1] == byte('.')) {
		toggle = true
	}

	for r, inp := range boardInp {
		for c, s := range inp {
			if string(s) == "#" {
				img[pos{r, c}] = true
			} else {
				img[pos{r, c}] = false
			}
		}
	}

	for i := 0; i < 50; i++ {
		newImg := image{}
		minW, maxW, minH, maxH := img.getDims()
		d := dimensions{minW, maxW, minH, maxH}

		for r := minH - 1; r <= maxH+1; r++ {
			for c := minW - 1; c <= maxW+1; c++ {
				p := pos{r, c}
				nbours := p.getNeighbours()
				binaryString := strings.Builder{}
				for _, n := range nbours {
					if (border && isGreaterThanBorder(d, n)) || img[n] {
						binaryString.WriteString("1")
					} else {
						binaryString.WriteString("0")
					}
				}
				bs := binaryString.String()
				algoPos := binaryToInt(bs)
				if string(algo[algoPos]) == "#" {
					newImg[p] = true
				}
			}
		}
		img = newImg
		if toggle {
			border = !border
		}
		// img.drawImg()
	}

	res := 0
	for _, pixelOn := range img {
		if pixelOn {
			res++
		}
	}
	return res
}

func part2() int {
	return 12
}

func binaryToInt(binary string) int {
	output, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(output)
}
