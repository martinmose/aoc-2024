package day23

import (
	"strings"
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
kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
      `,
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := part1Puzzle(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestPart2Puzzle(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Examples from the website",
			input: `
ka-co
ta-co
de-co
ta-ka
de-ta
ka-de
      `,
			expected: "co,de,ka,ta",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := part2Puzzle(strings.TrimSpace(tt.input))
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}
