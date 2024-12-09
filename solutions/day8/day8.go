package day8

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

type Position struct {
	x, y int
}

// Run runs the day 8 challenge
func Run() error {
	fmt.Println("Day 8:")

	dayPath := "8/input"
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
	antinodes := make(map[Position]bool)
	totalAntinodes := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' {
				continue
			}

			for x := 0; x < len(grid); x++ {
				for y := 0; y < len(grid[x]); y++ {
					if i == x && j == y || grid[i][j] != grid[x][y] {
						continue
					}

					manhattan := Position{x - i, y - j}
					antinode := Position{x + manhattan.x, y + manhattan.y}

					if isValidPosition(antinode, grid) && !antinodes[antinode] {
						antinodes[antinode] = true
						totalAntinodes++
					}
				}
			}
		}
	}

	return totalAntinodes
}

func part2Puzzle(input string) int {
	grid := parseInput(input)
	antinodes := make(map[Position]bool)
	totalAntinodes := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '.' {
				continue
			}

			for x := 0; x < len(grid); x++ {
				for y := 0; y < len(grid[x]); y++ {
					if i == x && j == y || grid[i][j] != grid[x][y] {
						continue
					}

					manhattan := Position{x - i, y - j}
					antinode := Position{x, y}

					positionsToMark := []Position{{i, j}, {x, y}}

					for {
						antinode = Position{antinode.x + manhattan.x, antinode.y + manhattan.y}

						if !isValidPosition(antinode, grid) {
							break
						}

						positionsToMark = append(positionsToMark, antinode)
					}

					for _, pos := range positionsToMark {
						if !antinodes[pos] {
							antinodes[pos] = true
							totalAntinodes++
						}
					}
				}
			}
		}
	}

	return totalAntinodes
}

func parseInput(input string) [][]rune {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func isValidPosition(pos Position, grid [][]rune) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid)
}
