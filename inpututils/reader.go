package inpututils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads from the filepath and outputs array of lines as strings
func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var lines []string
	for Scanner.Scan() {
		lines = append(lines, Scanner.Text())
	}
	return lines
}

// ReadNumbers reads from the filepath and attempts to convert each line to an int
func ReadNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var numbers []int
	for Scanner.Scan() {
		numbers = append(numbers, toInt(Scanner.Text()))
	}
	return numbers
}

// ReadRaw returns the content of a text file as a string
func ReadRaw(filename string) string {
	content, err := ioutil.ReadFile(filename)
	check(err)
	return string(content)
}

// ReadSlice attempts to create a slice of strings for a comma-seperated txt file
func ReadSlice(filename string) []string {
	content := ReadRaw(filename)
	return append([]string{}, strings.Split(content, ",")...)
}

// GetInputPath gets the input path
func GetInputPath() string {
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		input = "input.txt"
	}
	return input
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
