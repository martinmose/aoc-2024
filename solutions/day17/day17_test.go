package day17

import (
	"reflect"
	"strings"
	"testing"
)

func TestPart1Puzzle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name: "Examples from the website",
			input: `
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
      `,
			expected: []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseOutput(part1Puzzle(tt.input))
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestPart2Puzzle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "Examples from the website",
			input: `
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0
      `,
			expected: 117440,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := part2Puzzle(tt.input)
			resultInt := parseInt(result)
			if resultInt != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, resultInt)
			}
		})
	}
}

func parseOutput(output string) []int {
	var result []int
	for _, part := range strings.Split(output, ",") {
		result = append(result, parseInt(part))
	}
	return result
}
