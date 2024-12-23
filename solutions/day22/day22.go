package day22

import (
	"aoc_2024/utils"
	"strconv"
	"strings"
)

const pruneNum = 16777216

// Run runs the day 22 challenge
func Run() error {
	dayPath := "22/input"
	body, err := utils.HTTPGet(dayPath)
	if err != nil {
		return err
	}

	part1Result := part1Puzzle(body)
	println("Part 1 result:", part1Result)

	part2Result := part2Puzzle(body)
	println("Part 2 result:", part2Result)

	return nil
}

func part1Puzzle(input string) int {
	nums := parseInput(input)
	sum := 0
	for _, num := range nums {
		for i := 0; i < 2000; i++ {
			num = getSecretNum(num)
		}
		sum += num
	}
	return sum
}

func part2Puzzle(input string) int {
	nums := parseInput(input)
	sequences := make(map[[4]int]int)

	for _, num := range nums {
		seen := make(map[[4]int]bool)
		last4 := [4]int{10, 10, 10, 10} // Start with impossible changes

		for i := 0; i < 2000; i++ {
			prev := num % 10
			num = getSecretNum(num)
			curr := num % 10

			// Shift the window and add new change
			last4[0] = last4[1]
			last4[1] = last4[2]
			last4[2] = last4[3]
			last4[3] = curr - prev

			// Update sequence map
			if !seen[last4] {
				seen[last4] = true
				sequences[last4] += curr
			}
		}
	}

	// Find max value in sequences
	maxValue := 0
	for _, sum := range sequences {
		if sum > maxValue {
			maxValue = sum
		}
	}
	return maxValue
}

func parseInput(input string) []int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ints := make([]int, len(lines))
	for i, line := range lines {
		ints[i], _ = strconv.Atoi(line)
	}
	return ints
}

func getSecretNum(num int) int {
	num = prune(mix(num*64, num))
	num = prune(mix(num/32, num))
	num = prune(mix(num*2048, num))
	return num
}

// Helpers
func mix(a, b int) int { return a ^ b }
func prune(a int) int  { return a % pruneNum }
