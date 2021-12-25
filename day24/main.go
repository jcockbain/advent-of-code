package main

import (
	_ "embed"
	"regexp"
	"strconv"

	"fmt"

	"github.com/jcockbain/advent-of-code-2021/utils"
)

var (
	benchmark = false
	re        = regexp.MustCompile(`([a-z]+) ([a-z]) ?([a-z]|-?\d+)?`)
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := solve()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

type bucket struct {
	v1    int
	v2    int
	v3    int
	steps []string
}

var v1Line = 4
var v2Line = 5
var v3Line = 15

func solve() (int, int) {
	cmds := utils.GetLines(input)

	// buckets that "repeat" every 18 lines
	// these have 3 variables that differ, record these
	buckets := []bucket{}
	for i := 0; i < len(cmds); i += 18 {
		vs := make([]int, 3)
		for x, p := range []int{v1Line, v2Line, v3Line} {
			parts := re.FindStringSubmatch(cmds[i+p])
			vs[x] = toInt(parts[3])
		}
		buckets = append(buckets, bucket{
			v1:    vs[0],
			v2:    vs[1],
			v3:    vs[2],
			steps: cmds[i : i+18],
		})
	}
	// in each bucket - there are 3 variables

	// these were worked out by hand, based on the variables in each bucket
	highest := 99999795919456
	// fmt.Println(run(toIntArray(highest), cmds))

	lowest := 45311191516111
	// fmt.Println(run(toIntArray(lowest), cmds))

	return highest, lowest
}

func toIntArray(n int) []int {
	res := make([]int, 14)
	for p, i := range fmt.Sprint(n) {
		res[p] = toInt(string(i))
	}
	return res
}

func run(inp []int, cmds []string) (int, int, int, int) {
	res := map[string]int{
		"x": 0,
		"y": 0,
		"z": 0,
		"w": 0,
	}
	getVarOrInt := func(i string) int {
		if isNumber(i) {
			return toInt(i)
		}
		return res[i]
	}
	inpIdx := 0
	for _, c := range cmds {
		parts := re.FindStringSubmatch(c)
		mode := parts[1]
		switch mode {
		case "inp":
			v := parts[2]
			i := inp[inpIdx]
			inpIdx++
			res[v] = i
		case "add":
			v1, v2 := parts[2], getVarOrInt(parts[3])
			res[v1] = res[v1] + v2
		case "mul":
			v1, v2 := parts[2], getVarOrInt(parts[3])
			res[v1] = res[v1] * v2
		case "div":
			v1, v2 := parts[2], getVarOrInt(parts[3])
			res[v1] = res[v1] / v2
		case "mod":
			v1, v2 := parts[2], getVarOrInt(parts[3])
			res[v1] = res[v1] % v2
		case "eql":
			v1, v2 := parts[2], getVarOrInt(parts[3])
			if res[v1] == v2 {
				res[v1] = 1
			} else {
				res[v1] = 0
			}
		}
	}
	return res["x"], res["y"], res["z"], res["w"]
}

func isNumber(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}
	return false
}

func toInt(x string) int {
	i, err := strconv.Atoi(x)
	check(err)
	return i
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
