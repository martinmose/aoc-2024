package day21

import (
	"aoc_2024/utils"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

var numMap = map[string]coord{
	"A": {2, 0},
	"0": {1, 0},
	"1": {0, 1},
	"2": {1, 1},
	"3": {2, 1},
	"4": {0, 2},
	"5": {1, 2},
	"6": {2, 2},
	"7": {0, 3},
	"8": {1, 3},
	"9": {2, 3},
}

var dirMap = map[string]coord{
	"A": {2, 1},
	"^": {1, 1},
	"<": {0, 0},
	"v": {1, 0},
	">": {2, 0},
}

// Run runs the day 21 challenge
func Run() error {
	dayPath := "21/input"
	body, err := utils.HTTPGet(dayPath)
	if err != nil {
		return err
	}

	part1Result := part1Puzzle(body)
	println("Part 1 result:", part1Result)

	part2Result := part2Puzzle(body)
	println("Part 2 result:", part2Result)

	return nil
}

func part1Puzzle(input string) int {
	robots := 2
	lines := splitInput(input)
	return getSequence(lines, numMap, dirMap, robots)
}

func part2Puzzle(input string) int {
	robots := 25
	lines := splitInput(input)
	return getSequence(lines, numMap, dirMap, robots)
}

func splitInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func getSequence(lines []string, numMap, dirMap map[string]coord, robotCount int) int {
	total := 0
	cache := make(map[string][]int)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		chars := strings.Split(line, "")
		moves := getNumPadSequence(chars, "A", numMap)
		length := countSequences(moves, robotCount, 1, cache, dirMap)
		total += atoiNoErr(strings.TrimLeft(line[:3], "0")) * length
	}

	return total
}

func getNumPadSequence(input []string, start string, numMap map[string]coord) []string {
	curr := numMap[start]
	seq := []string{}

	for _, char := range input {
		dest := numMap[char]
		dx, dy := dest.x-curr.x, dest.y-curr.y

		horiz, vert := []string{}, []string{}

		for i := 0; i < abs(dx); i++ {
			if dx >= 0 {
				horiz = append(horiz, ">")
			} else {
				horiz = append(horiz, "<")
			}
		}

		for i := 0; i < abs(dy); i++ {
			if dy >= 0 {
				vert = append(vert, "^")
			} else {
				vert = append(vert, "v")
			}
		}

		if curr.y == 0 && dest.x == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if curr.x == 0 && dest.y == 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, "A")
	}
	return seq
}

func countSequences(input []string, maxRobots, robot int, cache map[string][]int, dirMap map[string]coord) int {
	key := strings.Join(input, "")
	if val, ok := cache[key]; ok && robot <= len(val) && val[robot-1] != 0 {
		return val[robot-1]
	}

	if _, ok := cache[key]; !ok {
		cache[key] = make([]int, maxRobots)
	}

	seq := getDirPadSequence(input, "A", dirMap)
	if robot == maxRobots {
		return len(seq)
	}

	steps := splitSequence(seq)
	count := 0
	for _, step := range steps {
		c := countSequences(step, maxRobots, robot+1, cache, dirMap)
		count += c
	}

	cache[key][robot-1] = count
	return count
}

func getDirPadSequence(input []string, start string, dirMap map[string]coord) []string {
	curr := dirMap[start]
	seq := []string{}

	for _, char := range input {
		dest := dirMap[char]
		dx, dy := dest.x-curr.x, dest.y-curr.y

		horiz, vert := []string{}, []string{}

		for i := 0; i < abs(dx); i++ {
			if dx >= 0 {
				horiz = append(horiz, ">")
			} else {
				horiz = append(horiz, "<")
			}
		}

		for i := 0; i < abs(dy); i++ {
			if dy >= 0 {
				vert = append(vert, "^")
			} else {
				vert = append(vert, "v")
			}
		}

		if curr.x == 0 && dest.y == 1 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if curr.y == 1 && dest.x == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, "A")
	}
	return seq
}

func splitSequence(input []string) [][]string {
	var result [][]string
	var current []string

	for _, char := range input {
		current = append(current, char)
		if char == "A" {
			result = append(result, current)
			current = []string{}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func atoiNoErr(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
