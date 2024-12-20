package day16

import (
	"aoc_2024/utils"
	"fmt"
	"sort"
	"strings"
)

type point struct {
	X, Y int
}

type direction struct {
	Dx, Dy int
}

type maze struct {
	Grid  [][]string
	Start point
	End   point
}

type queueItem struct {
	Pos   point
	Dir   int
	Score int
	Path  []point
}

var directions = []direction{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
}

const (
	wall       = "#"
	startConst = "S"
	endConst   = "E"
	turnCost   = 1000
	moveCost   = 1
	startDir   = 1
)

// Run runs the day 16 challenge
func Run() error {
	fmt.Println("Day 16:")

	dayPath := "16/input"
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
	maze := parseMaze(input)
	return findLowestScore(maze)
}

func part2Puzzle(input string) int {
	maze := parseMaze(input)
	lowestScore := findLowestScore(maze)
	paths := findAllOptimalPaths(maze, lowestScore)
	return countUniqueTiles(paths)
}

func parseMaze(input string) maze {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]string, len(lines))
	var start, end point

	for y, line := range lines {
		grid[y] = strings.Split(line, "")
		for x, ch := range grid[y] {
			switch ch {
			case startConst:
				start = point{X: x, Y: y}
			case endConst:
				end = point{X: x, Y: y}
			}
		}
	}

	return maze{Grid: grid, Start: start, End: end}
}

func findLowestScore(m maze) int {
	queue := []queueItem{{Pos: m.Start, Dir: startDir, Score: 0}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Score < queue[j].Score
		})

		current := queue[0]
		queue = queue[1:]

		if isEnd(m, current.Pos) {
			return current.Score
		}

		key := current.Pos.key(current.Dir)
		if visited[key] {
			continue
		}
		visited[key] = true

		// Try moving forward
		nextPos := current.Pos.add(directions[current.Dir])
		if isValid(m, nextPos) {
			queue = append(queue, queueItem{
				Pos:   nextPos,
				Dir:   current.Dir,
				Score: current.Score + moveCost,
			})
		}

		// Try turning
		for _, newDir := range []int{(current.Dir + 1) % 4, (current.Dir + 3) % 4} {
			queue = append(queue, queueItem{
				Pos:   current.Pos,
				Dir:   newDir,
				Score: current.Score + turnCost,
			})
		}
	}

	return -1
}

func findAllOptimalPaths(m maze, targetScore int) [][]point {
	queue := []queueItem{{Pos: m.Start, Dir: startDir, Score: 0, Path: []point{m.Start}}}
	visited := make(map[string]int)
	var paths [][]point

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Score > targetScore {
			continue
		}

		key := current.Pos.key(current.Dir)
		if score, exists := visited[key]; exists && score < current.Score {
			continue
		}
		visited[key] = current.Score

		if isEnd(m, current.Pos) && current.Score == targetScore {
			paths = append(paths, current.Path)
			continue
		}

		// Try moving forward
		nextPos := current.Pos.add(directions[current.Dir])
		if isValid(m, nextPos) {
			newPath := make([]point, len(current.Path))
			copy(newPath, current.Path)
			queue = append(queue, queueItem{
				Pos:   nextPos,
				Dir:   current.Dir,
				Score: current.Score + moveCost,
				Path:  append(newPath, nextPos),
			})
		}

		// Try turning
		for _, newDir := range []int{(current.Dir + 1) % 4, (current.Dir + 3) % 4} {
			queue = append(queue, queueItem{
				Pos:   current.Pos,
				Dir:   newDir,
				Score: current.Score + turnCost,
				Path:  current.Path,
			})
		}
	}

	return paths
}

func countUniqueTiles(paths [][]point) int {
	unique := make(map[string]bool)
	for _, path := range paths {
		for _, p := range path {
			unique[p.key(0)] = true
		}
	}
	return len(unique)
}

func isValid(m maze, p point) bool {
	return p.Y >= 0 && p.Y < len(m.Grid) &&
		p.X >= 0 && p.X < len(m.Grid[0]) &&
		m.Grid[p.Y][p.X] != wall
}

func isEnd(m maze, p point) bool {
	return p == m.End
}

func (p point) add(d direction) point {
	return point{X: p.X + d.Dx, Y: p.Y + d.Dy}
}

func (p point) key(dir int) string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, dir)
}
