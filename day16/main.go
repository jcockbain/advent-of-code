package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"fmt"
)

var (
	benchmark = false
)

//go:embed input.txt
var input string

func main() {
	p1, p2 := part1()

	if !benchmark {
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	}
}

func part1() (int, int) {
	b := hexToBin(input)
	total, pVersion, current := 0, 0, 0
	for current < len(b)-6 {
		count, version, finalIdx := parsePacket(b[current:])
		current += finalIdx
		total += count
		pVersion += version
	}
	return pVersion, total
}

func parsePacket(b string) (int, int, int) {
	typeID := binaryToInt(b[3:6])
	current := 0
	pVersion, total := 0, 0
	if typeID == 4 {
		count, version, finalIdx := parseLiteralValue(b[current:])
		pVersion += version
		total += count
		current += finalIdx
	} else {
		count, version, finalIdx := parseOperatorValue(b[current:])
		pVersion += version
		total += count
		current += finalIdx
	}
	return total, pVersion, current
}

// return value, version, final index
func parseLiteralValue(b string) (int, int, int) {
	pVersion := binaryToInt(b[:3])
	totalBin := strings.Builder{}
	startBit := 6
	for string(b[startBit]) == "1" {
		totalBin.WriteString(b[startBit+1 : startBit+5])
		startBit += 5
	}
	totalBin.WriteString(b[startBit+1 : startBit+5])
	endBit := startBit + 5
	paddedBits := (endBit - 6) % 5
	return binaryToInt(totalBin.String()), pVersion, endBit + paddedBits
}

// return value, version, final index
func parseOperatorValue(b string) (int, int, int) {
	pVersion := binaryToInt(b[:3])
	lengthID := string(b[6])
	typeID := binaryToInt(b[3:6])
	if lengthID == "0" {
		length := binaryToInt(b[7:22])
		current := 22
		packets := []int{}
		for current < 22+length {
			count, version, finalIdx := parsePacket(b[current:])
			packets = append(packets, count)
			pVersion += version
			current += finalIdx
		}
		return doOperator(packets, typeID), pVersion, current
	}
	numberOfPackets := binaryToInt(b[7:18])
	packets := [][]int{}
	current := 18
	for len(packets) < numberOfPackets {
		total, version, finalIdx := parsePacket(b[current:])
		current += finalIdx
		packets = append(packets, []int{total, version})
	}
	packetValues := getVals(packets)
	return doOperator(packetValues, typeID), pVersion + sumByIdx(packets, 1), current
}

func sum(p []int) (i int) {
	for _, x := range p {
		i += x
	}
	return
}

func product(p []int) int {
	i := 1
	for _, x := range p {
		i *= x
	}
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSlice(s []int) int {
	m := math.MaxInt
	for _, x := range s {
		m = min(m, x)
	}
	return m
}

func maxSlice(s []int) int {
	m := 0
	for _, x := range s {
		if x > m {
			m = x
		}
	}
	return m
}

func equal(p []int) int {
	if p[0] == p[1] {
		return 1
	}
	return 0
}

func less(p []int) int {
	if p[0] < p[1] {
		return 1
	}
	return 0
}

func great(p []int) int {
	if p[0] > p[1] {
		return 1
	}
	return 0
}

func doOperator(packets []int, code int) int {
	switch code {
	case 0:
		return sum(packets)
	case 1:
		return product(packets)
	case 2:
		return minSlice(packets)
	case 3:
		return maxSlice(packets)
	case 5:
		return great(packets)
	case 6:
		return less(packets)
	case 7:
		return equal(packets)
	}
	panic("no!")
}

func getVals(x [][]int) []int {
	vals := []int{}
	for _, p := range x {
		vals = append(vals, p[0])
	}
	return vals
}

func sumByIdx(x [][]int, i int) int {
	s := 0
	for _, p := range x {
		s += p[i]
	}
	return s
}

func hexToBin(hex string) string {
	s := strings.Builder{}
	for _, c := range hex {
		s.WriteString(getBinFromHex(string(c)))
	}
	return s.String()
}

func getBinFromHex(hex string) string {
	ui, err := strconv.ParseUint(hex, 16, 64)
	check(err)
	format := fmt.Sprintf("%%0%db", len(hex)*4)
	return fmt.Sprintf(format, ui)
}

func binaryToInt(bin string) int {
	i, err := strconv.ParseInt(bin, 2, 64)
	check(err)
	return int(i)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
