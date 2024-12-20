package day19

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

// Run runs the day 19 challenge
func Run() error {
	fmt.Println("Day 19:")

	dayPath := "19/input"
	body, err := utils.HTTPGet(dayPath)
	if err != nil {
		return err
	}

	part1Result := part1Puzzle(body)
	fmt.Println("Part 1 result:", part1Result)

	part2Result := part2Puzzle(body)
	fmt.Println("Part 2 result:", part2Result)

	return nil
}

func part1Puzzle(input string) int {
	patterns, designs := parseInput(input)
	return solveDesigns(patterns, designs, true)
}

func part2Puzzle(input string) int {
	patterns, designs := parseInput(input)
	return solveDesigns(patterns, designs, false)
}

func parseInput(input string) ([]string, []string) {
	sections := strings.Split(input, "\n\n")
	patterns := strings.Split(strings.TrimSpace(sections[0]), ", ")
	designs := strings.Split(strings.TrimSpace(sections[1]), "\n")
	return patterns, designs
}

func solveDesigns(patterns []string, designs []string, countMatching bool) int {
	cache := make(map[string]int)
	out := 0

	var solve func(string) int
	solve = func(s string) int {
		if _, exists := cache[s]; !exists {
			if len(s) == 0 {
				return 1
			}
			res := 0
			for _, pattern := range patterns {
				if strings.HasPrefix(s, pattern) {
					res += solve(s[len(pattern):])
				}
			}
			cache[s] = res
		}
		return cache[s]
	}

	for _, design := range designs {
		if countMatching {
			if solve(design) > 0 {
				out++
			}
		} else {
			out += solve(design)
		}
	}

	return out
}
