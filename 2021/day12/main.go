package main

import (
	_ "embed"
	"regexp"
	"strings"

	"fmt"

	"github.com/jcockbain/advent-of-code/utils"
)

var (
	benchmark = false
	re        = regexp.MustCompile(`(.+)-(.+)`)
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

type graph map[string]*node

func (g graph) getNeighbours(s string) []*node {
	return g[s].edges
}

type node struct {
	id    string
	edges []*node
}

func part1() int {
	g := graph{}
	for _, line := range utils.GetLines(input) {
		parts := re.FindStringSubmatch(line)
		start, end := parts[1], parts[2]

		if _, in := g[start]; !in {
			g[start] = &node{
				start,
				[]*node{},
			}
		}
		if _, in := g[end]; !in {
			g[end] = &node{
				end,
				[]*node{},
			}
		}

		g[start].edges = append(g[start].edges, g[end])
		g[end].edges = append(g[end].edges, g[start])
	}

	var dfs func(string, []string)
	var allPaths [][]string

	dfs = func(s string, currentPath []string) {
		if s == "end" {
			allPaths = append(allPaths, currentPath)
		}

		for _, n := range g.getNeighbours(s) {
			if isUpperCase(n.id) || !inPath(currentPath, n.id) {
				dfs(n.id, append(currentPath, n.id))
			}
		}
	}

	dfs("start", []string{"start"})
	return len(allPaths)
}

func inPath(path []string, s string) bool {
	for _, p := range path {
		if p == s {
			return true
		}
	}
	return false
}

func isUpperCase(s string) bool {
	return (strings.ToUpper(s) == s)
}

func part2() int {
	g := graph{}
	for _, line := range utils.GetLines(input) {
		parts := re.FindStringSubmatch(line)
		start, end := parts[1], parts[2]

		if _, in := g[start]; !in {
			g[start] = &node{
				start,
				[]*node{},
			}
		}
		if _, in := g[end]; !in {
			g[end] = &node{
				end,
				[]*node{},
			}
		}

		g[start].edges = append(g[start].edges, g[end])
		g[end].edges = append(g[end].edges, g[start])
	}

	var dfs func(string, []string)
	var allPaths [][]string

	dfs = func(s string, currentPath []string) {
		if s == "end" {
			allPaths = append(allPaths, currentPath)
			return
		}

		for _, n := range g.getNeighbours(s) {
			if isUpperCase(n.id) || validSecondSmallCave(currentPath, n.id) {
				dfs(n.id, append(currentPath, n.id))
			}
		}
	}

	dfs("start", []string{"start"})
	return len(allPaths)
}

func validSecondSmallCave(path []string, s string) bool {
	caveCount := count(path, s)
	if caveCount == 0 {
		return true
	}
	if caveCount > 1 {
		return false
	}
	if s == "start" || s == "end" {
		return false
	}
	// check for other duplicate
	alreadySeen := map[string]bool{}
	for _, p := range path {
		if !isUpperCase(p) {
			if _, in := alreadySeen[p]; in {
				return false
			}
			alreadySeen[p] = true
		}
	}
	return true
}

func count(sl []string, t string) (c int) {
	for _, s := range sl {
		if s == t {
			c += 1
		}
	}
	return c
}
