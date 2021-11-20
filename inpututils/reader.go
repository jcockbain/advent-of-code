package inpututils

import (
	"bufio"
	"os"
	"strconv"
	"io/ioutil"
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

// ReadLines reads from the filepath and attempts to convert each line to an int
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