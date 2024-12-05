package day4

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

// Run runs the day 4 challenge
func Run() error {
	fmt.Println("Day 4:")

	dayPath := "4/input"
	body, err := utils.HTTPGet(dayPath)
	if err != nil {
		return err
	}

	grid := strings.Split(strings.TrimSpace(body), "\n")

	part1Result := part1Puzzle(grid)

	fmt.Println("Part 1 result:", part1Result)

	part2Result := part2Puzzle(grid)

	fmt.Println("Part 2 result:", part2Result)

	return nil
}

func part1Puzzle(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	word := "XMAS"
	wordLen := len(word)
	total := 0

	checkDirection := func(r, c, dr, dc int) bool {
		for i := 0; i < wordLen; i++ {
			nr, nc := r+(i*dr), c+(i*dc)
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != word[i] {
				return false
			}
		}
		return true
	}

	directions := [][2]int{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Diagonal down-right
		{-1, -1}, // Diagonal up-left
		{-1, 1},  // Diagonal up-right
		{1, -1},  // Diagonal down-left
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if checkDirection(r, c, dir[0], dir[1]) {
					total++
				}
			}
		}
	}

	return total
}

func part2Puzzle(grid []string) int {
	width := len(grid[0])
	height := len(grid)
	count := 0

	getPos := func(x, y int) byte {
		if x < 0 || x >= width || y < 0 || y >= height {
			return 0
		}
		return grid[y][x]
	}

	// Traverse the grid
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'A' {
				lines := 0
				for dx := -1; dx <= 1; dx += 2 {
					for dy := -1; dy <= 1; dy += 2 {
						if getPos(x+dx, y+dy) == 'M' && getPos(x-dx, y-dy) == 'S' {
							lines++
						}
					}
				}
				if lines == 2 {
					count++
				}
			}
		}
	}

	return count
}
