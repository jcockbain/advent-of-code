package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 15, Part1("test1.txt"))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 12, Part2("test1.txt"))
}
