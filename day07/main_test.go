package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testpart1(t *testing.T) {
	assert.Equal(t, 15, part1("test1.txt"))
}

func Testpart2(t *testing.T) {
	assert.Equal(t, 12, part2("test1.txt"))
}
