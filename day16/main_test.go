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

func TestHexToBin(t *testing.T) {

	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "D2FE28", want: "110100101111111000101000"},
		{input: "38006F45291200", want: "00111000000000000110111101000101001010010001001000000000"},
		{input: "EE00D40C823060", want: "11101110000000001101010000001100100000100011000001100000"},
		{input: "8A004A801A8002F478", want: "100010100000000001001010100000000001101010000000000000101111010001111000"},
	}

	for _, test := range tests {
		got := hexToBin(test.input)
		assert.Equal(t, test.want, got)
	}
}

func TestParsePacket(t *testing.T) {
	total, _, _ := parsePacket("00111000000000000110111101000101001010010001001000000000")
	assert.Equal(t, 30, total)
	total, _, _ = parsePacket("11101110000000001101010000001100100000100011000001100000")
	assert.Equal(t, 6, total)
}

func TestParseLiteralValue(t *testing.T) {
	total, _, _ := parseLiteralValue("11010001010")
	assert.Equal(t, total, 10)

	b := "0101001000100100"
	total2, _, finalIdx := parseLiteralValue(b)
	assert.Equal(t, total2, 20)
	assert.Equal(t, len(b), finalIdx)

}

func TestParseOperatorValue(t *testing.T) {
	total, _, _ := parseOperatorValue("00111000000000000110111101000101001010010001001000000000")
	assert.Equal(t, 30, total)
}
