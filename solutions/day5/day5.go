package day5

import (
	"aoc_2024/utils"
	"fmt"
	"strconv"
	"strings"
)

// Run runs the day 5 challenge
func Run() error {
	fmt.Println("Day 5:")

	dayPath := "5/input"
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
	rulesInput, updatesInput := splitInput(input)
	rules := parseRules(rulesInput)
	updates := parseUpdates(updatesInput)
	total := processUpdates(updates, rules)
	return total
}

func part2Puzzle(input string) int {
	rulesInput, updatesInput := splitInput(input)
	rules := parseRules(rulesInput)
	updates := parseUpdates(updatesInput)
	total := 0

	for _, update := range updates {
		if !isUpdateValid(update, rules) {
			correctedOrder := reorderUpdate(update, rules)
			total += findMiddlePage(correctedOrder)
		}
	}

	return total
}

func splitInput(input string) ([]string, []string) {
	sections := strings.Split(input, "\n\n")
	rules := strings.Split(strings.TrimSpace(sections[0]), "\n")
	updates := strings.Split(strings.TrimSpace(sections[1]), "\n")
	return rules, updates
}

func parseRules(rules []string) map[int][]int {
	orderingRules := make(map[int][]int)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		before, err1 := strconv.Atoi(parts[0])
		after, err2 := strconv.Atoi(parts[1])
		if err1 == nil && err2 == nil {
			orderingRules[before] = append(orderingRules[before], after)
		}
	}
	return orderingRules
}

func parseUpdates(updates []string) [][]int {
	parsed := [][]int{}
	for _, update := range updates {
		parts := strings.Split(update, ",")
		updatePages := make([]int, 0, len(parts))
		for _, part := range parts {
			page, err := strconv.Atoi(part)
			if err == nil {
				updatePages = append(updatePages, page)
			}
		}
		parsed = append(parsed, updatePages)
	}
	return parsed
}

func findMiddlePage(update []int) int {
	mid := len(update) / 2
	return update[mid]
}

func isUpdateValid(update []int, ruleMap map[int][]int) bool {
	pagePosition := make(map[int]int)
	for pos, page := range update {
		pagePosition[page] = pos
	}

	for page1, dependencies := range ruleMap {
		pos1, exists1 := pagePosition[page1]
		if !exists1 {
			continue
		}

		for _, page2 := range dependencies {
			pos2, exists2 := pagePosition[page2]
			if !exists2 {
				continue
			}

			if pos1 >= pos2 {
				return false
			}
		}
	}

	return true
}

func processUpdates(updates [][]int, ruleMap map[int][]int) int {
	total := 0
	for _, update := range updates {
		if isUpdateValid(update, ruleMap) {
			total += findMiddlePage(update)
		}
	}
	return total
}

func reorderUpdate(update []int, rules map[int][]int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)
	for _, page := range update {
		graph[page] = []int{}
		inDegree[page] = 0
	}

	for page1, dependencies := range rules {
		for _, page2 := range dependencies {
			if contains(update, page1) && contains(update, page2) {
				graph[page1] = append(graph[page1], page2)
				inDegree[page2]++
			}
		}
	}

	var sorted []int
	var queue []int

	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func contains(slice []int, element int) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}
