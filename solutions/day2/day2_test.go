package day2

import (
	"testing"
)

func TestPart1Puzzle(t *testing.T) {
	tests := []struct {
		name     string
		levels   [][]int
		expected int
	}{
		{
			name: "Examples from the website",
			levels: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := part1Puzzle(tt.levels)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestPart2Puzzle(t *testing.T) {
	tests := []struct {
		name     string
		levels   [][]int
		expected int
	}{
		{
			name: "Examples from the website",
			levels: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := part2Puzzle(tt.levels)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
