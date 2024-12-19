package day11

import (
	"aoc_2024/utils"
	"fmt"
	"math/big"
	"strings"
)

// Run executes the Day 11 challenge
func Run() error {
	fmt.Println("Day 11:")

	dayPath := "11/input"
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

func part1Puzzle(input string) *big.Int {
	stones := parseInput(input)
	processed := processStones(stones, 25)
	return processed
}

func part2Puzzle(input string) *big.Int {
	stones := parseInput(input)
	processed := processStones(stones, 75)
	return processed
}

func parseInput(input string) map[string]*big.Int {
	stones := map[string]*big.Int{}
	parts := strings.Fields(strings.TrimSpace(input))
	for _, part := range parts {
		if stones[part] == nil {
			stones[part] = big.NewInt(0)
		}
		stones[part].Add(stones[part], big.NewInt(1))
	}
	return stones
}

func processStones(stones map[string]*big.Int, blinks int) *big.Int {
	for i := 0; i < blinks; i++ {
		newStones := map[string]*big.Int{}

		for value, count := range stones {
			n := big.NewInt(0)
			n.SetString(value, 10)

			if n.Cmp(big.NewInt(0)) == 0 {
				addToMap(newStones, "1", count)
			} else if numDigits(n)%2 == 0 {
				power := power10(numDigits(n) / 2)
				left := big.NewInt(0).Div(n, power)
				right := big.NewInt(0).Mod(n, power)

				addToMap(newStones, left.String(), count)
				addToMap(newStones, right.String(), count)
			} else {
				multiplied := big.NewInt(0).Mul(n, big.NewInt(2024))
				addToMap(newStones, multiplied.String(), count)
			}
		}

		stones = newStones
	}

	totalCount := big.NewInt(0)
	for _, count := range stones {
		totalCount.Add(totalCount, count)
	}

	return totalCount
}

func numDigits(n *big.Int) int {
	return len(n.String())
}

func power10(n int) *big.Int {
	result := big.NewInt(1)
	for i := 0; i < n; i++ {
		result.Mul(result, big.NewInt(10))
	}
	return result
}

func addToMap(m map[string]*big.Int, key string, count *big.Int) {
	if m[key] == nil {
		m[key] = big.NewInt(0)
	}
	m[key].Add(m[key], count)
}
