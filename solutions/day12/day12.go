package day12

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

type position struct {
	x, y int
}

var width, height int

// Run runs the day 12 challenge
func Run() error {
	fmt.Println("Day 12:")

	dayPath := "12/input"
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
	stack := make([]position, 0, 1000)
	sum := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] >= 'A' && grid[y][x] <= 'Z' {
				stack = append(stack, position{x, y})
			}

			squares := 0
			fences := 0

			for len(stack) > 0 {
				pos := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				c := tile(grid, pos)
				if c >= 'a' && c <= 'z' {
					continue
				}
				squares++

				lowerC := c + ('a' - 'A')
				directions := []position{
					{0, -1}, // Up
					{1, 0},  // Right
					{0, 1},  // Down
					{-1, 0}, // Left
				}

				for _, d := range directions {
					nextPos := add(pos, d)
					nextC := tile(grid, nextPos)

					if nextC == c {
						stack = append(stack, nextPos)
					} else if nextC != lowerC {
						fences++
					}
				}

				grid[pos.y][pos.x] = lowerC
			}

			sum += squares * fences
		}
	}

	return sum
}

func part2Puzzle(input string) int {
	grid := parseInput(input)
	stack := make([]position, 0, 1000)
	sum := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] >= 'A' && grid[y][x] <= 'Z' {
				stack = append(stack, position{x, y})
			}

			squares := 0
			fences := 0

			for len(stack) > 0 {
				pos := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				c := tile(grid, pos)
				if c >= 'a' && c <= 'z' {
					continue
				}
				squares++

				lowerC := c + ('a' - 'A')
				directions := []position{
					{0, -1},  // Up
					{1, -1},  // Up-Right
					{1, 0},   // Right
					{1, 1},   // Down-Right
					{0, 1},   // Down
					{-1, 1},  // Down-Left
					{-1, 0},  // Left
					{-1, -1}, // Up-Left
				}

				for i, d := range directions {
					if i%2 != 0 {
						continue
					}

					nextPos := add(pos, d)
					nextC := tile(grid, nextPos)

					if nextC == c {
						stack = append(stack, nextPos)
					}
				}

				for r := 0; r < 4; r++ {
					b := [3]bool{}
					for p := 0; p < 3; p++ {
						nextPos := add(pos, directions[(p+(r*2))%8])
						nextC := tile(grid, nextPos)
						b[p] = (nextC == c || nextC == lowerC)
					}

					if !b[0] && !b[1] && !b[2] {
						fences++
					}
					if b[0] && !b[1] && b[2] {
						fences++
					}
					if !b[0] && b[1] && !b[2] {
						fences++
					}
				}

				grid[pos.y][pos.x] = lowerC
			}

			sum += squares * fences
		}
	}

	return sum
}

func parseInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	height = len(lines)
	width = len(lines[0])
	grid := make([][]rune, height)
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func add(a, b position) position {
	return position{x: a.x + b.x, y: a.y + b.y}
}

func tile(grid [][]rune, pos position) rune {
	if pos.x < 0 || pos.x >= width || pos.y < 0 || pos.y >= height {
		return '.'
	}
	return grid[pos.y][pos.x]
}
