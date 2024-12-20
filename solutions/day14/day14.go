package day14

import (
	"aoc_2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type position struct {
	X, Y int
}

type velocity struct {
	X, Y int
}

const (
	xMax = 101
	yMax = 103
)

// Run runs the day 14 challenge
func Run() error {
	fmt.Println("Day 14:")

	dayPath := "14/input"
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
	positions, velocities := parseInput(input)
	newPositions := calculateNewPositions(positions, velocities, 100)
	quadrantCounts := make([]int, 5)

	for _, pos := range newPositions {
		quadrant := getQuadrant(pos)
		if quadrant > 0 {
			quadrantCounts[quadrant]++
		}
	}

	safetyFactor := 1
	for i := 1; i <= 4; i++ {
		if quadrantCounts[i] == 0 {
			return 0
		}
		safetyFactor *= quadrantCounts[i]
	}

	return safetyFactor
}

func part2Puzzle(input string) int {
	positions, velocities := parseInput(input)
	seconds := 0
	var deviationCalibration float64

	for {
		average := calculateAverage(positions)
		deviation := calculateDeviation(positions, average)

		if deviationCalibration == 0 {
			deviationCalibration = float64(deviation[0] + deviation[1])
		} else if deviationCalibration/float64(deviation[0]+deviation[1]) > 1.5 {
			break
		}

		updatePositions(positions, velocities, 1)
		seconds++
	}

	return seconds
}

func parseInput(input string) ([]position, []velocity) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var positions []position
	var velocities []velocity

	for _, line := range lines {
		segments := strings.Fields(
			strings.NewReplacer("=", " ", ",", " ").Replace(line),
		)
		posX, _ := strconv.Atoi(segments[1])
		posY, _ := strconv.Atoi(segments[2])
		velX, _ := strconv.Atoi(segments[4])
		velY, _ := strconv.Atoi(segments[5])

		positions = append(positions, position{X: posX, Y: posY})
		velocities = append(velocities, velocity{X: velX, Y: velY})
	}

	return positions, velocities
}

func calculateNewPositions(positions []position, velocities []velocity, steps int) []position {
	var newPositions []position
	for i := range positions {
		newX := (positions[i].X + velocities[i].X*steps) % xMax
		if newX < 0 {
			newX += xMax
		}
		newY := (positions[i].Y + velocities[i].Y*steps) % yMax
		if newY < 0 {
			newY += yMax
		}
		newPositions = append(newPositions, position{X: newX, Y: newY})
	}
	return newPositions
}

func getQuadrant(pos position) int {
	if pos.X == xMax/2 || pos.Y == yMax/2 {
		return 0
	}

	if pos.X < xMax/2 {
		if pos.Y < yMax/2 {
			return 1
		}
		return 2
	}

	if pos.Y < yMax/2 {
		return 3
	}

	return 4
}

func calculateAverage(positions []position) [2]int {
	var avg [2]int
	for _, pos := range positions {
		avg[0] += pos.X
		avg[1] += pos.Y
	}
	avg[0] /= len(positions)
	avg[1] /= len(positions)
	return avg
}

func calculateDeviation(positions []position, average [2]int) [2]int {
	var deviation [2]int
	for _, pos := range positions {
		deviation[0] += int(math.Abs(float64(average[0] - pos.X)))
		deviation[1] += int(math.Abs(float64(average[1] - pos.Y)))
	}
	deviation[0] /= len(positions)
	deviation[1] /= len(positions)
	return deviation
}

func updatePositions(positions []position, velocities []velocity, steps int) {
	for i := range positions {
		positions[i].X = (positions[i].X + velocities[i].X*steps) % xMax
		if positions[i].X < 0 {
			positions[i].X += xMax
		}
		positions[i].Y = (positions[i].Y + velocities[i].Y*steps) % yMax
		if positions[i].Y < 0 {
			positions[i].Y += yMax
		}
	}
}
