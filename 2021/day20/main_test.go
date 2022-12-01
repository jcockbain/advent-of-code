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

func TestBinaryToInt(t *testing.T) {
	tests := []struct {
		inp  string
		want int
	}{
		{"000100010", 34},
	}
	for _, test := range tests {
		got := binaryToInt(test.inp)
		assert.Equal(t, test.want, got)
	}
}
