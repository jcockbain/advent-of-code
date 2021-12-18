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

func TestExplode(t *testing.T) {
	tests := []struct {
		inp  string
		want string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
		{"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
		{"[[[[0,7],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[15,[0,13]]],[1,1]]"},
	}
	for _, test := range tests {
		got := explode(test.inp)
		assert.Equal(t, test.want, got)
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		inp  string
		want string
	}{
		{"[[[[0,7],4],[15,[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"},
	}
	for _, test := range tests {
		got := split(test.inp)
		assert.Equal(t, test.want, got)
	}
}

func TestTransform(t *testing.T) {
	tests := []struct {
		inp  string
		want string
	}{
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
		{"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"},
	}
	for _, test := range tests {
		got := transform(test.inp)
		assert.Equal(t, test.want, got)
	}
}

func TestTransformPair(t *testing.T) {
	tests := []struct {
		inp1 string
		inp2 string
		want string
	}{
		{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
		{"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]", "[2,9]", "[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]"},
	}
	for _, test := range tests {
		got := transformPair(test.inp1, test.inp2)
		assert.Equal(t, test.want, got)
	}
}

func TestCalcMag(t *testing.T) {
	tests := []struct {
		inp  string
		want int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
	}
	for _, test := range tests {
		got := calcMag(test.inp)
		assert.Equal(t, test.want, got)
	}
}
