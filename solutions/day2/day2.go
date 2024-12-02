package day2

import (
	"aoc_2024/utils"
	"fmt"
	"strconv"
	"strings"
)

// Run runs the day 2 challenge
func Run() error {
	fmt.Println("Day 1:")

	dayPath := "2/input"
	body, err := utils.HTTPGet(dayPath)
	if err != nil {
		return err
	}

	var levels [][]int
	lines := strings.Split(strings.TrimSpace(body), "\n")

	for _, line := range lines {
		// Split each line into fields (space-separated values)
		fields := strings.Fields(line)

		// Convert fields to integers
		var row []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return fmt.Errorf("failed to parse '%s' as integer: %w", field, err)
			}
			row = append(row, num)
		}

		levels = append(levels, row)
	}

	part1Result, err := part1Puzzle(levels)
	if err != nil {
		return err
	}

	fmt.Println("Part 1 result:", part1Result)

	part2Result := part2Puzzle(levels)

	fmt.Println("Part 2 result:", part2Result)

	return nil
}

func part1Puzzle(levels [][]int) (int, error) {
	var numberOfSafeLevels int

	for _, row := range levels {
		isSafeLevel := true
		direction := 0 // 1 for increasing, -1 for decreasing

		for i := 0; i < len(row)-1; i++ {
			diff := row[i+1] - row[i]

			// Convert to absolute value inline
			if diff < 0 {
				diff = -diff
			}

			if diff < 1 || diff > 3 {
				isSafeLevel = false
				break
			}

			if direction == 0 {
				if row[i+1] > row[i] {
					direction = 1
				} else if row[i+1] < row[i] {
					direction = -1
				}
			} else if (row[i+1] > row[i] && direction == -1) || (row[i+1] < row[i] && direction == 1) {
				isSafeLevel = false
				break
			}
		}

		if isSafeLevel {
			numberOfSafeLevels++
		}
	}

	return numberOfSafeLevels, nil
}

func isRowSafe(row []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 0; i < len(row)-1; i++ {
		diff := row[i+1] - row[i]

		if !(diff >= 1 && diff <= 3) {
			isIncreasing = false
		}

		if !(diff >= -3 && diff <= -1) {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func part2Puzzle(levels [][]int) int {
	var numberOfSafeLevels int

	for _, row := range levels {
		if isRowSafe(row) {
			numberOfSafeLevels++
			continue
		}

		isSafeWithOneRemoval := false
		for j := 0; j < len(row); j++ {
			temp := append([]int{}, row[:j]...)
			temp = append(temp, row[j+1:]...)

			if isRowSafe(temp) {
				isSafeWithOneRemoval = true
				break
			}
		}

		if isSafeWithOneRemoval {
			numberOfSafeLevels++
		}
	}

	return numberOfSafeLevels
}
