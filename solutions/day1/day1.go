package day1

import (
	"aoc_2024/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Run runs the day 1 challenge
func Run() error {
	fmt.Println("Day 1:")

	dayPath := "1/input"
	body, err := utils.HTTPGet(dayPath)
	if err != nil {
		return err
	}

	lines := strings.Split(body, "\n")

	var list1, list2 []int
	for _, line := range lines {
		columns := strings.Fields(line)
		if len(columns) >= 2 {
			num1, err1 := strconv.Atoi(columns[0])
			num2, err2 := strconv.Atoi(columns[1])
			if err1 == nil && err2 == nil {
				list1 = append(list1, num1)
				list2 = append(list2, num2)
			}
		}
	}

	part1Result, err := part1Puzzle(list1, list2)
	if err != nil {
		return err
	}

	fmt.Println("Part 1 result:", part1Result)

	part2Result, err := part2Puzzle(list1, list2)
	if err != nil {
		return err
	}

	fmt.Println("Part 2 result:", part2Result)

	return nil
}

func part1Puzzle(list1 []int, list2 []int) (int, error) {
	sort.Ints(list1)
	sort.Ints(list2)

	var result int
	for i := 0; i < len(list1); i++ {
		result += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return result, nil
}

func part2Puzzle(list1 []int, list2 []int) (int, error) {
	elementCounts := make(map[int]int)
	for _, elem := range list2 {
		elementCounts[elem]++
	}

	var result int
	for _, element1 := range list1 {
		if similarity, exists := elementCounts[element1]; exists {
			result += element1 * similarity
		}
	}

	return result, nil
}
