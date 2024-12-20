package day18

import (
	"aoc_2024/utils"
	"container/heap"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type coord struct {
	X, Y int
}

type priorityItem struct {
	Coord    coord
	Priority int
	Index    int
}

type priorityQueue []*priorityItem

var directions = []coord{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

// Run runs the day 18 challenge
func Run() error {
	fmt.Println("Day 18:")

	dayPath := "18/input"
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

func part1Puzzle(input string) string {
	coords := parseInput(input)
	grid := buildGrid(coords, 1024)
	start, end := coord{0, 0}, coord{70, 70}
	return strconv.Itoa(findShortestPath(grid, start, end))
}

func part2Puzzle(input string) string {
	coords := parseInput(input)
	grid := make(map[coord]bool)
	start, end := coord{0, 0}, coord{70, 70}

	for _, coord := range coords {
		grid[coord] = true
		if !isPathExists(grid, start, end) {
			return fmt.Sprintf("%d,%d", coord.X, coord.Y)
		}
	}

	return "No blocking coordinate found"
}

func parseInput(input string) []coord {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var coords []coord
	for _, line := range lines {
		values := extractInts(line)
		coords = append(coords, coord{X: values[0], Y: values[1]})
	}
	return coords
}

func buildGrid(coords []coord, limit int) map[coord]bool {
	grid := make(map[coord]bool)
	for i, coord := range coords {
		if i >= limit {
			break
		}
		grid[coord] = true
	}
	return grid
}

func findShortestPath(grid map[coord]bool, start, end coord) int {
	distances := make(map[coord]int)
	distances[start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &priorityItem{Coord: start, Priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*priorityItem)
		if current.Coord == end {
			return current.Priority
		}

		if current.Priority > distances[current.Coord] {
			continue
		}

		for _, dir := range directions {
			next := coord{X: current.Coord.X + dir.X, Y: current.Coord.Y + dir.Y}
			if isValidMove(next, grid, end) {
				newDist := current.Priority + 1
				if existingDist, exists := distances[next]; !exists || newDist < existingDist {
					distances[next] = newDist
					heap.Push(pq, &priorityItem{Coord: next, Priority: newDist})
				}
			}
		}
	}

	return -1
}

func isPathExists(grid map[coord]bool, start, end coord) bool {
	visited := make(map[coord]bool)
	queue := []coord{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return true
		}

		for _, dir := range directions {
			next := coord{X: current.X + dir.X, Y: current.Y + dir.Y}
			if isValidMove(next, grid, end) && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return false
}

func isValidMove(pos coord, grid map[coord]bool, end coord) bool {
	return pos.X >= 0 && pos.Y >= 0 && pos.X <= end.X && pos.Y <= end.Y && !grid[pos]
}

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*priorityItem)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

// Helper functions to parse integers from a string
func extractInts(s string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(s, -1)
	var result []int
	for _, match := range matches {
		value, _ := strconv.Atoi(match)
		result = append(result, value)
	}
	return result
}
