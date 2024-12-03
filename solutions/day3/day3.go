package day3

import (
	"aoc_2024/utils"
	"fmt"
	"strconv"
)

// Run runs the day 3 challenge
func Run() error {
	fmt.Println("Day 3:")

	dayPath := "3/input"
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
	total := 0
	i := 0

	for i < len(input) {
		product, newIndex, success := parseMulInstruction(input, i)

		if success {
			total += product
			i = newIndex
			continue
		}

		i++
	}
	return total
}

func part2Puzzle(input string) int {
	total := 0
	i := 0
	mulEnabled := true

	for i < len(input) {
		// Check for do() or don't()
		if len(input) >= i+4 && input[i:i+4] == "do()" {
			mulEnabled = true
			i += 4 // Move past "do()"
			continue
		} else if len(input) >= i+7 && input[i:i+7] == "don't()" {
			mulEnabled = false
			i += 7 // Move past "don't()"
			continue
		}

		if mulEnabled {
			product, newIndex, success := parseMulInstruction(input, i)
			if success {
				total += product
				i = newIndex
				continue
			}
		}

		i++
	}

	return total
}

func parseMulInstruction(input string, startIndex int) (product int, newIndex int, success bool) {
	i := startIndex

	// Ensure "mul(" prefix exists
	if len(input) < i+4 || input[i:i+4] != "mul(" {
		return 0, startIndex, false
	}
	i += 4 // Move past "mul("

	// Extract the first number
	start := i
	for i < len(input) && (input[i] >= '0' && input[i] <= '9') {
		i++
	}
	if i == start {
		return 0, startIndex, false // No valid number found
	}
	num1, err := strconv.Atoi(input[start:i])
	if err != nil {
		return 0, startIndex, false
	}

	// Ensure there is a comma next
	if i >= len(input) || input[i] != ',' {
		return 0, startIndex, false
	}
	i++ // Move past ','

	// Extract the second number
	start = i
	for i < len(input) && (input[i] >= '0' && input[i] <= '9') {
		i++
	}
	if i == start {
		return 0, startIndex, false // No valid number found
	}
	num2, err := strconv.Atoi(input[start:i])
	if err != nil {
		return 0, startIndex, false
	}

	// Ensure the sequence ends with a closing parenthesis
	if i >= len(input) || input[i] != ')' {
		return 0, startIndex, false
	}
	i++ // Move past ')'

	// Return the product of the two numbers and the updated index
	return num1 * num2, i, true
}
