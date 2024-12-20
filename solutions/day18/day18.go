package day18

import (
	"aoc_2024/utils"
	"container/heap"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

type PriorityItem struct {
	Coord    Coord
	Priority int
	Index    int
}

type PriorityQueue []*PriorityItem

var directions = []Coord{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

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
	start, end := Coord{0, 0}, Coord{70, 70}
	return strconv.Itoa(findShortestPath(grid, start, end))
}

func part2Puzzle(input string) string {
	coords := parseInput(input)
	grid := make(map[Coord]bool)
	start, end := Coord{0, 0}, Coord{70, 70}

	for _, coord := range coords {
		grid[coord] = true
		if !isPathExists(grid, start, end) {
			return fmt.Sprintf("%d,%d", coord.X, coord.Y)
		}
	}

	return "No blocking coordinate found"
}

func parseInput(input string) []Coord {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var coords []Coord
	for _, line := range lines {
		values := extractInts(line)
		coords = append(coords, Coord{X: values[0], Y: values[1]})
	}
	return coords
}

func buildGrid(coords []Coord, limit int) map[Coord]bool {
	grid := make(map[Coord]bool)
	for i, coord := range coords {
		if i >= limit {
			break
		}
		grid[coord] = true
	}
	return grid
}

func findShortestPath(grid map[Coord]bool, start, end Coord) int {
	distances := make(map[Coord]int)
	distances[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PriorityItem{Coord: start, Priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PriorityItem)
		if current.Coord == end {
			return current.Priority
		}

		if current.Priority > distances[current.Coord] {
			continue
		}

		for _, dir := range directions {
			next := Coord{X: current.Coord.X + dir.X, Y: current.Coord.Y + dir.Y}
			if isValidMove(next, grid, end) {
				newDist := current.Priority + 1
				if existingDist, exists := distances[next]; !exists || newDist < existingDist {
					distances[next] = newDist
					heap.Push(pq, &PriorityItem{Coord: next, Priority: newDist})
				}
			}
		}
	}

	return -1
}

func isPathExists(grid map[Coord]bool, start, end Coord) bool {
	visited := make(map[Coord]bool)
	queue := []Coord{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			return true
		}

		for _, dir := range directions {
			next := Coord{X: current.X + dir.X, Y: current.Y + dir.Y}
			if isValidMove(next, grid, end) && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return false
}

func isValidMove(pos Coord, grid map[Coord]bool, end Coord) bool {
	return pos.X >= 0 && pos.Y >= 0 && pos.X <= end.X && pos.Y <= end.Y && !grid[pos]
}

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*PriorityItem)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
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
