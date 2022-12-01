package main

import (
	_ "embed"

	"fmt"

	"github.com/jcockbain/advent-of-code-2021/utils"
)

var (
	benchmark = false
)

//go:embed input.txt
var input string

func main() {
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type Cube struct {
	x1, x2, y1, y2, z1, z2 int
	isOn                   bool
}

func (c Cube) intersection(other *Cube) *Cube {
	x1 := max(c.x1, other.x1)
	x2 := min(c.x2, other.x2)
	if x2 < x1 {
		return nil
	}

	y1 := max(c.y1, other.y1)
	y2 := min(c.y2, other.y2)
	if y2 < y1 {
		return nil
	}

	z1 := max(c.z1, other.z1)
	z2 := min(c.z2, other.z2)
	if z2 < z1 {
		return nil
	}

	return &Cube{
		x1: x1,
		x2: x2,
		y1: y1,
		y2: y2,
		z1: z1,
		z2: z2,
	}

}

func (cube *Cube) volume() int {
	vol := ((cube.x2 + 1) - cube.x1) *
		((cube.y2 + 1) - cube.y1) *
		((cube.z2 + 1) - cube.z1)
	if vol < 0 {
		return -vol
	}
	return vol
}

type Cubes []*Cube

func (cubes *Cubes) count() (sum int) {
	// resolve intersections backwards
	for i := len(*cubes) - 1; i >= 0; i-- {
		cube := (*cubes)[i]

		if !cube.isOn {
			continue
		}

		intersections := &Cubes{}
		for _, next := range (*cubes)[i+1:] {
			intersection := cube.intersection(next)
			if intersection == nil {
				continue
			}
			intersection.isOn = true
			*intersections = append(*intersections, intersection)
		}

		sum += cube.volume()
		sum -= intersections.count()
	}
	return
}

func part2() int {
	lines := utils.GetLines(input)
	cubes := &Cubes{}
	for _, l := range lines {
		var x1, x2, y1, y2, z1, z2 int
		var on string
		fmt.Sscanf(l, "%s x=%d..%d,y=%d..%d,z=%d..%d", &on, &x1, &x2, &y1, &y2, &z1, &z2)
		*cubes = append(*cubes, &Cube{
			x1,
			x2,
			y1,
			y2,
			z1,
			z2,
			on == "on",
		})
	}
	return cubes.count()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
