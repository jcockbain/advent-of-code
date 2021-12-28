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

func TestGetRoute(t *testing.T) {
	b := parseBurrow()
	type test struct {
		start    pos
		dest     pos
		expected []pos
	}

	tests := []test{
		{
			pos{2, 3},
			pos{2, 5},
			[]pos{{2, 3}, {1, 3}, {1, 4}, {1, 5}, {2, 5}},
		},
		{
			pos{1, 1},
			pos{2, 3},
			[]pos{{1, 1}, {1, 2}, {1, 3}, {2, 3}},
		},
		{
			pos{1, 1},
			pos{1, 5},
			[]pos{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}},
		},
		{
			pos{3, 5},
			pos{5, 3},
			[]pos{{3, 5}, {2, 5}, {1, 5}, {1, 4}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}},
		},
		{
			pos{1, 3},
			pos{1, 6},
			[]pos{{1, 3}, {1, 4}, {1, 5}, {1, 6}},
		},
		{
			pos{1, 6},
			pos{5, 5},
			[]pos{{1, 6}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5}},
		},
	}
	for _, test := range tests {
		route := b.getRoute(test.start, test.dest)
		assert.Equal(t, test.expected, route)
	}
}
