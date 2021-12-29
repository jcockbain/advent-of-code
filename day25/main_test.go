package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkMain(b *testing.B) {
	benchmark = true
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestMain(t *testing.T) {
	assert.Equal(t, 549, part1())
}
