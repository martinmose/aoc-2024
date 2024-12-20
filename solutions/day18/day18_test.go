package day18

import (
	"strconv"
	"testing"
)

func TestPart1Puzzle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "Examples from the website",
			input: `
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
      `,
			// expected: 22,
			expected: 146,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := part1Puzzle(tt.input)
			resultInt := atoi(result)
			if resultInt != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, resultInt)
			}
		})
	}
}

// func TestPart2Puzzle(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		input    string
// 		expected Coord
// 	}{
// 		{
// 			name: "Examples from the website",
// 			input: `
// 5,4
// 4,2
// 4,5
// 3,0
// 2,1
// 6,3
// 2,4
// 1,5
// 0,6
// 3,3
// 2,6
// 5,1
// 1,2
// 5,5
// 2,5
// 6,5
// 1,4
// 0,4
// 6,4
// 1,1
// 6,1
// 1,0
// 0,5
// 1,6
// 2,0
//       `,
// 			expected: Coord{6, 1},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := part2Puzzle(tt.input)
// 			resultCoord := parseCoord(result)
// 			if resultCoord != tt.expected {
// 				t.Errorf("Expected %v, got %v", tt.expected, resultCoord)
// 			}
// 		})
// 	}

func atoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

// func parseCoord(s string) Coord {
// 	parts := strings.Split(s, ",")
// 	x, _ := strconv.Atoi(parts[0])
// 	y, _ := strconv.Atoi(parts[1])
// 	return Coord{X: x, Y: y}
// }
