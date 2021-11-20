package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	// assert.Equal(t, 0, Part1(1))
	assert.Equal(t, 3, Part1(12))
	assert.Equal(t, 2, Part1(23))
	assert.Equal(t, 31, Part1(1024))
}

// func TestPart2(t *testing.T) {
// 	assert.Equal(t, 12, Part2("input.txt"))
// }
