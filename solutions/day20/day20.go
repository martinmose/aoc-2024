package day20

import (
	"aoc_2024/utils"
	"fmt"
	"strings"
)

type position struct {
	Row, Col int
}

type grid struct {
	Cells      [][]byte
	Start, End position
	Rows       int
	Cols       int
}

// Run runs the day 20 challenge
func Run() error {
	fmt.Println("Day 20:")

	dayPath := "20/input"
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
	grid := parseGrid(input)
	return solve(grid, 2)
}

func part2Puzzle(input string) int {
	grid := parseGrid(input)
	return solve(grid, 20)
}

func solve(grid grid, maxSkipDist int) int {
	distances := bfs(grid)

	timeSkips := make(map[int]int)

	for pos, dist := range distances {
		for dr := -maxSkipDist; dr <= maxSkipDist; dr++ {
			for dc := -maxSkipDist; dc <= maxSkipDist; dc++ {
				skipPos := position{Row: pos.Row + dr, Col: pos.Col + dc}

				if !isInBounds(skipPos, grid) {
					continue
				}

				skipDist := abs(dr) + abs(dc)
				if skipDist == 0 || skipDist > maxSkipDist {
					continue
				}

				if endDist, reachable := distances[skipPos]; reachable {
					timeSaved := endDist - dist - skipDist
					if timeSaved > 0 {
						timeSkips[timeSaved]++
					}
				}
			}
		}
	}

	res := 0
	for saved, count := range timeSkips {
		if saved >= 100 {
			res += count
		}
	}

	return res
}

func parseGrid(input string) grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)
	cols := len(lines[0])
	cells := make([][]byte, rows)

	var start, end position

	for r, line := range lines {
		cells[r] = []byte(line)
		for c, char := range line {
			if char == 'S' {
				start = position{Row: r, Col: c}
			} else if char == 'E' {
				end = position{Row: r, Col: c}
			}
		}
	}

	return grid{Cells: cells, Start: start, End: end, Rows: rows, Cols: cols}
}

func bfs(grid grid) map[position]int {
	distances := make(map[position]int)
	queue := []position{grid.Start}
	distances[grid.Start] = 0

	moves := []position{
		{Row: -1, Col: 0},
		{Row: 1, Col: 0},
		{Row: 0, Col: -1},
		{Row: 0, Col: 1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, move := range moves {
			neighbor := position{Row: current.Row + move.Row, Col: current.Col + move.Col}
			if !isInBounds(neighbor, grid) || grid.Cells[neighbor.Row][neighbor.Col] == '#' {
				continue
			}

			if _, visited := distances[neighbor]; !visited {
				distances[neighbor] = distances[current] + 1
				queue = append(queue, neighbor)
			}
		}
	}

	return distances
}

func isInBounds(pos position, grid grid) bool {
	return pos.Row >= 0 && pos.Row < grid.Rows && pos.Col >= 0 && pos.Col < grid.Cols
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
