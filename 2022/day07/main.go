package main

import (
	_ "embed"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
	cdRe      = regexp.MustCompile(`\$ (cd) ([a-z]|.)*`)
)

//go:embed input.txt
var input string

type dir struct {
	name     string
	parent   *dir
	children map[string]*dir
	files    []*file
}

type file struct {
	name string
	size int
}

func main() {
	p1 := part1()
	p2 := part2()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func getRoot() *dir {
	lines := utils.GetLines(input)
	root := &dir{
		children: map[string]*dir{},
	}
	current := root
	for _, line := range lines[1:] {
		if cdRe.Match([]byte(line)) {
			parts := strings.Split(line, " ")
			directive := parts[2]
			switch directive {
			case "..":
				current = current.parent
			default:
				current = current.children[directive]
			}
		} else if line != "$ ls" {
			split := strings.Split(line, " ")
			first, second := split[0], split[1]
			if first == "dir" {
				d := dir{
					name:     second,
					parent:   current,
					children: map[string]*dir{},
				}
				current.children[second] = &d
			} else {
				f := file{
					name: second,
					size: toInt(first),
				}
				current.files = append(current.files, &f)
			}
		}
	}
	return root
}

func part1() int {
	res := 0
	root := getRoot()
	var dfs func(n *dir) int
	dfs = func(n *dir) int {
		cumulative := 0
		for _, f := range n.files {
			cumulative += f.size
		}
		for _, c := range n.children {
			cumulative += dfs(c)
		}
		if cumulative <= 100000 {
			res += cumulative
		}
		return cumulative
	}
	dfs(root)
	return res
}

func part2() int {
	root := getRoot()
	allSizes := []int{}
	var dfs func(n *dir) int
	dfs = func(n *dir) int {
		cumulative := 0
		for _, f := range n.files {
			cumulative += f.size
		}
		for _, c := range n.children {
			cumulative += dfs(c)
		}
		allSizes = append(allSizes, cumulative)
		return cumulative
	}
	dfs(root)
	sort.Ints(allSizes)
	currentSize := allSizes[len(allSizes)-1]
	targetToRemove := currentSize - 40000000
	for _, size := range allSizes {
		if size > targetToRemove {
			return size
		}
	}
	panic("something's gone wrong!")
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
