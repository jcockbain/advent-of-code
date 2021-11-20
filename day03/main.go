package main

import (
	"fmt"
	"math"
)

// var shortestDistance = map[int]int{}

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1(289326))
}

func Part1(input int) int {
	perfectSquares := map[int]int{}

	if input == 1 {
		return 0
	}

	max := math.Sqrt(float64(input)) + float64(2)
	for i := 1; float64(i) < max; i += 2 {
		perfectSquares[i*i] = int(math.Floor(float64(i) / 2))
	}

	var sideLength, endOfLayer, layer int

	i := input
	for {
		if l, ok := perfectSquares[i]; ok {
			layer = l
			sideLength = 1 + (2 * layer)
			endOfLayer = i
			break
		}
		i += 1
	}

	sideDistance := sideLength - 1
	middleOfBottom := endOfLayer - int((sideLength-1)/2)

	middlePoints := []int{
		middleOfBottom,
		middleOfBottom - sideDistance,
		middleOfBottom - (2 * sideDistance),
		middleOfBottom - (3 * sideDistance),
	}

	shortestPath := input
	for _, m := range middlePoints {
		shortestPath = min(shortestPath, abs(input-m))
	}
	return layer + shortestPath
}

func Part2(input int) int {

	type pos struct {
		x, y int
	}

	vals := map[pos]int{
		pos{0, 0}: 1,
	}
	getNeighbours := func(p pos) (v int) {
		for dx := range []int{-1, 0, 1} {
			for dy := range []int{-1, 0, 1} {
				if dx != 0 || dy != 0 {
					newPos := pos{x: p.x + dx, y: p.y + dy}
					if val, ok := vals[newPos]; ok {
						v += val
					}
				}
			}
		}
		return v
	}

	current_dir := 0
	steps := 1
	stepsTaken := 0

	current_pos := pos{1, 0}
	for {
		v := getNeighbours(current_pos)
		
		if v > input {
			return v
		}

	}
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Part2(filename string) int {
	return 12
}
