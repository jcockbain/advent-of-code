package main

import (
	_ "embed"

	"regexp"
	"sort"
	"strings"

	"github.com/jcockbain/advent-of-code/utils"

	"fmt"
	"strconv"
)

var (
	re        = regexp.MustCompile(`(.+)\|(.+)`)
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

var display = map[string]string{
	"abcefg":  "0",
	"cf":      "1",
	"acdeg":   "2",
	"acdfg":   "3",
	"bcdf":    "4",
	"abdfg":   "5",
	"abdefg":  "6",
	"acf":     "7",
	"abcdefg": "8",
	"abcdfg":  "9",
}

var numLightsToVals = map[int]int{
	2: 1,
	3: 7,
	4: 4,
	7: 8,
}

func part1() int {
	lines := utils.GetLines(input)
	totalUnique := 0
	for _, l := range lines {
		parts := re.FindStringSubmatch(l)
		_, out := strings.Split(parts[1], " "), strings.Split(parts[2], " ")
		for _, o := range out {
			l := len(o)
			if isUnique(l) {
				totalUnique += 1
			}
		}
	}
	return totalUnique
}

func isUnique(x int) bool {
	for l, _ := range numLightsToVals {
		if x == l {
			return true
		}
	}
	return false
}

type positionMap map[string]string

func (p positionMap) getKeys() (res []string) {
	for k, _ := range p {
		res = append(res, k)
	}
	return
}

func (p positionMap) getKey(val string) string {
	for k, v := range p {
		if v == val {
			return k
		}
	}
	panic("no key!")
}

func part2() int {
	lines := utils.GetLines(input)
	total := 0
	for _, l := range lines {
		parts := re.FindStringSubmatch(l)
		in, out := strings.Split(strings.TrimSpace(parts[1]), " "), strings.Split(strings.TrimSpace(parts[2]), " ")
		in = sortAllStrings(in)
		out = sortAllStrings(out)

		positionMapping := positionMap{}
		numberMapping := map[int]string{}

		for _, i := range in {
			if len(i) == 2 {
				numberMapping[1] = i
			}
			if len(i) == 3 {
				numberMapping[7] = i
			}
			if len(i) == 4 {
				numberMapping[4] = i
			}
			if len(i) == 7 {
				numberMapping[8] = i
			}
		}

		positionMapping[getExtraChar(numberMapping[1], numberMapping[7])] = "a"
		possibles := []string{}

		for _, x := range numberMapping[7] {
			s := string(x)
			if _, in := positionMapping[s]; !in {
				possibles = append(possibles, string(x))
			}
		}

		sixes := getStringsOfLen(in, 6)
		fives := getStringsOfLen(in, 5)
		for i, p := range possibles {
			for _, s := range sixes {
				if !strings.Contains(s, p) {
					positionMapping[p] = "c"
					positionMapping[possibles[1-i]] = "f"
				}
			}
		}

		requiredForThree := []string{positionMapping.getKey("a"), positionMapping.getKey("c"), positionMapping.getKey("f")}

		for _, f := range fives {
			if containsAllStrings(f, requiredForThree) {
				possibles = getExtraChars(strings.Join(requiredForThree, ""), f)
			}
		}

		for _, s := range sixes {
			for i, p := range possibles {
				if !strings.Contains(s, p) {
					positionMapping[p] = "d"
					positionMapping[possibles[1-i]] = "g"
				}
			}
		}

		possibles = getExtraChars(strings.Join(positionMapping.getKeys(), ""), "abcdefg")

		for _, s := range sixes {
			for i, p := range possibles {
				if !strings.Contains(s, p) {
					positionMapping[p] = "e"
					positionMapping[possibles[1-i]] = "b"
				}
			}
		}
		total += calculateTotal(out, positionMapping)
	}

	return total
}

func containsAllStrings(s string, strs []string) bool {
	for _, str := range strs {
		if !strings.Contains(s, str) {
			return false
		}
	}
	return true
}

func getStringsOfLen(strs []string, l int) (res []string) {
	for _, s := range strs {
		if len(s) == l {
			res = append(res, s)
		}
	}
	return
}

func calculateTotal(out []string, p positionMap) (total int) {
	vals := []string{}
	for _, s := range out {
		og := []string{}
		for _, c := range s {
			og = append(og, p[string(c)])
		}
		ogString := sortString(strings.Join(og, ""))

		vals = append(vals, (display[ogString]))
	}
	return stringToInt(strings.Join(vals, ""))
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func sortAllStrings(l []string) []string {
	for i, s := range l {
		l[i] = sortString(s)
	}
	return l
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func getExtraChars(a, b string) (res []string) {
	aChars := map[string]bool{}
	for _, s := range a {
		aChars[string(s)] = true
	}
	for _, s := range b {
		if _, in := aChars[string(s)]; !in {
			res = append(res, string(s))
		}
	}
	return
}

func getExtraChar(a, b string) string {
	aChars := map[string]bool{}
	for _, s := range a {
		aChars[string(s)] = true
	}
	for _, s := range b {
		if _, in := aChars[string(s)]; !in {
			return string(s)
		}
	}
	panic("no extra char")
}
