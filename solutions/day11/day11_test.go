package day11

import (
	"math/big"
	"testing"
)

func TestPart1Puzzle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *big.Int
	}{
		{
			name:     "Examples from the website - 25 blinks",
			input:    "125 17",
			expected: big.NewInt(55312),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := part1Puzzle(tt.input)
			if result.Cmp(tt.expected) != 0 {
				t.Errorf("Expected %s, got %s", tt.expected.String(), result.String())
			}
		})
	}
}

func TestPart2Puzzle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *big.Int
	}{
		{
			name:     "Examples from the website - 75 blinks",
			input:    "125 17",
			expected: big.NewInt(65601038650482),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := part2Puzzle(tt.input)
			if result.Cmp(tt.expected) != 0 {
				t.Errorf("Expected %s, got %s", tt.expected.String(), result.String())
			}
		})
	}
}
