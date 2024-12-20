package day16

import (
	"aoc_2024/utils"
	"fmt"
	"sort"
	"strings"
)

type Point struct {
	X, Y int
}

type Direction struct {
	Dx, Dy int
}

// Holds all the information about the maze
type Maze struct {
	Grid  [][]string
	Start Point
	End   Point
}

type QueueItem struct {
	Pos   Point
	Dir   int
	Score int
	Path  []Point
}

var directions = []Direction{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
}

const (
	Wall     = "#"
	Start    = "S"
	End      = "E"
	TurnCost = 1000
	MoveCost = 1
	StartDir = 1
)

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

func parseMaze(input string) Maze {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]string, len(lines))
	var start, end Point

	for y, line := range lines {
		grid[y] = strings.Split(line, "")
		for x, ch := range grid[y] {
			switch ch {
			case Start:
				start = Point{X: x, Y: y}
			case End:
				end = Point{X: x, Y: y}
			}
		}
	}

	return Maze{Grid: grid, Start: start, End: end}
}

func findLowestScore(m Maze) int {
	queue := []QueueItem{{Pos: m.Start, Dir: StartDir, Score: 0}}
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
			queue = append(queue, QueueItem{
				Pos:   nextPos,
				Dir:   current.Dir,
				Score: current.Score + MoveCost,
			})
		}

		// Try turning
		for _, newDir := range []int{(current.Dir + 1) % 4, (current.Dir + 3) % 4} {
			queue = append(queue, QueueItem{
				Pos:   current.Pos,
				Dir:   newDir,
				Score: current.Score + TurnCost,
			})
		}
	}

	return -1
}

func findAllOptimalPaths(m Maze, targetScore int) [][]Point {
	queue := []QueueItem{{Pos: m.Start, Dir: StartDir, Score: 0, Path: []Point{m.Start}}}
	visited := make(map[string]int)
	var paths [][]Point

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
			newPath := make([]Point, len(current.Path))
			copy(newPath, current.Path)
			queue = append(queue, QueueItem{
				Pos:   nextPos,
				Dir:   current.Dir,
				Score: current.Score + MoveCost,
				Path:  append(newPath, nextPos),
			})
		}

		// Try turning
		for _, newDir := range []int{(current.Dir + 1) % 4, (current.Dir + 3) % 4} {
			queue = append(queue, QueueItem{
				Pos:   current.Pos,
				Dir:   newDir,
				Score: current.Score + TurnCost,
				Path:  current.Path,
			})
		}
	}

	return paths
}

func countUniqueTiles(paths [][]Point) int {
	unique := make(map[string]bool)
	for _, path := range paths {
		for _, p := range path {
			unique[p.key(0)] = true
		}
	}
	return len(unique)
}

func isValid(m Maze, p Point) bool {
	return p.Y >= 0 && p.Y < len(m.Grid) &&
		p.X >= 0 && p.X < len(m.Grid[0]) &&
		m.Grid[p.Y][p.X] != Wall
}

func isEnd(m Maze, p Point) bool {
	return p == m.End
}

func (p Point) add(d Direction) Point {
	return Point{X: p.X + d.Dx, Y: p.Y + d.Dy}
}

func (p Point) key(dir int) string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, dir)
}
