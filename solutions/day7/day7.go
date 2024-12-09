package day7

import (
	"aoc_2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Run runs the day 7 challenge
func Run() error {
	fmt.Println("Day 7:")

	dayPath := "7/input"
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
	lines := strings.Split(strings.TrimSpace(input), "\n")
	totalSum := 0

	for _, line := range lines {
		target, numbers := parseLine(line)
		if isSolvable(target, numbers, false) {
			totalSum += target
		}
	}

	return totalSum
}

func part2Puzzle(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	totalSum := 0

	for _, line := range lines {
		target, numbers := parseLine(line)
		if isSolvable(target, numbers, true) {
			totalSum += target
		}
	}

	return totalSum
}

func parseLine(line string) (int, []int) {
	parts := strings.Split(line, ":")
	target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	numberStrings := strings.Fields(strings.TrimSpace(parts[1]))
	numbers := make([]int, len(numberStrings))

	for i, numStr := range numberStrings {
		numbers[i], _ = strconv.Atoi(numStr)
	}

	return target, numbers
}

func isSolvable(target int, numbers []int, allowConcat bool) bool {
	return checkCombinations(target, numbers, allowConcat)
}

func checkCombinations(target int, numbers []int, allowConcat bool) bool {
	totalOperators := len(numbers) - 1
	maxCombinations := int(math.Pow(3, float64(totalOperators)))

	for combination := 0; combination < maxCombinations; combination++ {
		result := numbers[0]
		currentCombination := combination

		for i := 0; i < totalOperators; i++ {
			operator := currentCombination % 3 // 0 = +, 1 = *, 2 = concatenation
			currentCombination /= 3

			switch operator {
			case 0: // Add
				result += numbers[i+1]
			case 1: // Multiply
				result *= numbers[i+1]
			case 2: // Concatenate
				if allowConcat {
					result = concatenateNumbers(result, numbers[i+1])
				}
			}
		}

		if result == target {
			return true
		}
	}

	return false
}

func concatenateNumbers(a, b int) int {
	magnitude := int(math.Pow(10, float64(numDigits(b))))
	return a*magnitude + b
}

func numDigits(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(n))) + 1
}
