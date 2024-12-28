package day25

import (
	"aoc_2024/utils"
	"strings"
)

// Run runs the day 25 challenge
func Run() error {
	dayPath := "25/input"
	body, err := utils.HTTPGet(dayPath)
	if err != nil {
		return err
	}

	part1Result := part1Puzzle(body)
	println("Part 1 result:", part1Result)

	return nil
}

func part1Puzzle(input string) int {
	locks, keys := parseInput(input)
	matches := 0

	for _, lock := range locks {
		for _, key := range keys {
			if checkMatch(lock, key) {
				matches++
			}
		}
	}

	return matches
}

func parseInput(input string) ([][]int, [][]int) {
	var locks, keys [][]int
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i := 0; i < len(lines); i += 8 {
		if i+7 > len(lines) {
			break
		}

		heights := make([]int, 5)
		isLock := false

		for row := 0; row < 7; row++ {
			for col, char := range lines[i+row] {
				if char == '#' {
					heights[col]++
				}
			}
			if row == 0 && lines[i][0] == '#' {
				isLock = true
			}
		}

		for i := range heights {
			heights[i]--
		}

		if isLock {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	return locks, keys
}

func checkMatch(lock, key []int) bool {
	for i := 0; i < 5; i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}
