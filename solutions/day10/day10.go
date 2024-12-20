package day10

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

type positionType struct {
	x, y int
}

type stackNode struct {
	start positionType
	end   positionType
}

// Run runs the day 10 challenge
func Run() error {
	fmt.Println("Day 10:")

	dayPath := "10/input"
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
	grid := parseInput(input)
	uniqueCount, _ := exploreTrails(grid, true)
	return uniqueCount
}

func part2Puzzle(input string) int {
	grid := parseInput(input)
	_, totalCount := exploreTrails(grid, false)
	return totalCount
}

func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))
	for y, line := range lines {
		grid[y] = make([]int, len(line))
		for x, char := range line {
			grid[y][x] = int(char - '0')
		}
	}
	return grid
}

func positionsEqual(a, b positionType) bool {
	return a.x == b.x && a.y == b.y
}

func nodesEqual(a, b stackNode) bool {
	return positionsEqual(a.start, b.start) && positionsEqual(a.end, b.end)
}

func tile(grid [][]int, pos positionType) int {
	if pos.x < 0 || pos.x >= len(grid[0]) {
		return -1
	}
	if pos.y < 0 || pos.y >= len(grid) {
		return -1
	}
	return grid[pos.y][pos.x]
}

func exploreTrails(grid [][]int, countUnique bool) (int, int) {
	stack := make([]stackNode, 0, 5000)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 0 {
				pos := positionType{x, y}
				stack = append(stack, stackNode{start: pos, end: pos})
			}
		}
	}

	paths := make([]stackNode, 0, 2000)
	uniqueCount := 0
	totalCount := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		position := node.end
		currentHeight := tile(grid, position)

		if currentHeight == 9 {
			totalCount++
			if countUnique {
				isUnique := true
				for _, path := range paths {
					if nodesEqual(path, node) {
						isUnique = false
						break
					}
				}
				if isUnique {
					paths = append(paths, node)
					uniqueCount++
				}
			}
			continue
		}

		deltas := []positionType{
			{0, -1}, // Up
			{1, 0},  // Right
			{0, 1},  // Down
			{-1, 0}, // Left
		}

		for _, delta := range deltas {
			nextPosition := positionType{
				x: position.x + delta.x,
				y: position.y + delta.y,
			}
			nextHeight := tile(grid, nextPosition)

			if nextHeight == currentHeight+1 {
				stack = append(stack, stackNode{start: node.start, end: nextPosition})
			}
		}
	}

	return uniqueCount, totalCount
}
