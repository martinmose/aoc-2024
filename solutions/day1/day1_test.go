package day1

import (
	"testing"
)

func TestPart1Puzzle(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected int
	}{
		{
			name:     "Examples from the website",
			list1:    []int{3, 4, 2, 1, 3, 3},
			list2:    []int{4, 3, 5, 3, 9, 3},
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := part1Puzzle(tt.list1, tt.list2)
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
		list1    []int
		list2    []int
		expected int
	}{
		{
			name:     "Examples from the website",
			list1:    []int{3, 4, 2, 1, 3, 3},
			list2:    []int{4, 3, 5, 3, 9, 3},
			expected: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := part2Puzzle(tt.list1, tt.list2)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
