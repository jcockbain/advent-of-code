package utils

import (
	"strconv"
	"strings"
)

// GetLines splits on new lines
func GetLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

// GetInts takes a string, and converts each line to an int
func GetInts(str string) []int {
	var numbers []int
	for _, line := range strings.Split(str, "\n") {
		numbers = append(numbers, toInt(line))
	}
	return numbers
}

func toInt(s string) int {
	converted, err := strconv.Atoi(s)
	check(err)
	return converted
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
