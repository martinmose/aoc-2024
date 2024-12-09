package day6

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

type Position struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

var directions = []Direction{
	{0, -1}, // Up (^)
	{1, 0},  // Right (>)
	{0, 1},  // Down (v)
	{-1, 0}, // Left (<)
}

// Run runs the day 6 challenge
func Run() error {
	fmt.Println("Day 6:")

	dayPath := "6/input"
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
	grid, startPos, startDir := parseGrid(input)
	visited := make(map[Position]bool)

	currentPos := startPos
	currentDir := startDir

	for {
		visited[currentPos] = true

		nextPos := Position{
			x: currentPos.x + currentDir.dx,
			y: currentPos.y + currentDir.dy,
		}

		if outOfBounds(nextPos, grid) {
			break
		}

		if grid[nextPos.y][nextPos.x] == '#' {
			currentDir = rotateDirection(currentDir)
		} else {
			currentPos = nextPos
		}
	}

	return len(visited)
}

func part2Puzzle(input string) int {
	grid, startPos, startDir := parseGrid(input)
	validPositions := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != '.' || (x == startPos.x && y == startPos.y) {
				continue
			}

			gridCopy := make([][]rune, len(grid))
			for i := range grid {
				gridCopy[i] = make([]rune, len(grid[i]))
				copy(gridCopy[i], grid[i])
			}

			gridCopy[y][x] = '#'

			if isLooping(startPos, startDir, gridCopy) {
				validPositions++
			}
		}
	}
	return validPositions
}

func isLooping(startPos Position, startDir Direction, grid [][]rune) bool {
	currentPos := startPos
	currentDir := startDir
	visitCount := make(map[Position]int)

	for steps := 0; steps < len(grid)*len(grid[0])*4; steps++ {
		visitCount[currentPos]++
		if visitCount[currentPos] > 4 {
			return true
		}

		nextPos := Position{
			x: currentPos.x + currentDir.dx,
			y: currentPos.y + currentDir.dy,
		}

		if outOfBounds(nextPos, grid) {
			return false
		}

		if grid[nextPos.y][nextPos.x] == '#' {
			currentDir = rotateDirection(currentDir)
		} else {
			currentPos = nextPos
		}
	}
	return false
}

func parseGrid(input string) ([][]rune, Position, Direction) {
	grid := [][]rune{}
	var startPos Position
	var startDir Direction

	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		row := []rune(line)
		grid = append(grid, row)
		for x, cell := range row {
			switch cell {
			case '^':
				startPos = Position{x, y}
				startDir = directions[0]
				row[x] = '.'
			case '>':
				startPos = Position{x, y}
				startDir = directions[1]
				row[x] = '.'
			case 'v':
				startPos = Position{x, y}
				startDir = directions[2]
				row[x] = '.'
			case '<':
				startPos = Position{x, y}
				startDir = directions[3]
				row[x] = '.'
			}
		}
	}

	return grid, startPos, startDir
}

func outOfBounds(pos Position, grid [][]rune) bool {
	return pos.x < 0 || pos.x >= len(grid[0]) || pos.y < 0 || pos.y >= len(grid)
}

func rotateDirection(dir Direction) Direction {
	return Direction{
		dx: -dir.dy,
		dy: dir.dx,
	}
}
