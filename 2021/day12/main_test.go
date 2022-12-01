package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T){
	assert.Equal(t, 5756, part1())
}

func TestPart2(t *testing.T){
	assert.Equal(t, 144603, part2())
}

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}
