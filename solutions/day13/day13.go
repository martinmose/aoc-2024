package day13

import (
	"aoc_2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type game struct {
	ButtonA []int64
	ButtonB []int64
	Prize   []int64
}

// Run runs the day 13 challenge
func Run() error {
	fmt.Println("Day 13:")

	dayPath := "13/input"
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

func part1Puzzle(input string) int64 {
	games := parseInput(input, false)
	var result int64

	for _, game := range games {
		a, b := solveEquation(game)
		if a > 0 && b > 0 && math.Abs(a-math.Round(a)) < 1e-6 && math.Abs(b-math.Round(b)) < 1e-6 {
			result += int64(math.Round(a))*3 + int64(math.Round(b))
		}
	}

	return result
}

func solveEquation(game game) (float64, float64) {
	denominator := float64(game.ButtonA[1])*float64(game.ButtonB[0]) - float64(game.ButtonA[0])*float64(game.ButtonB[1])
	if math.Abs(denominator) < 1e-6 {
		return -1, -1
	}

	aNumerator := float64(game.Prize[1])*float64(game.ButtonB[0]) - float64(game.Prize[0])*float64(game.ButtonB[1])
	a := aNumerator / denominator

	bNumerator := float64(game.Prize[0]) - a*float64(game.ButtonA[0])
	b := bNumerator / float64(game.ButtonB[0])

	return a, b
}

func part2Puzzle(input string) int64 {
	games := parseInput(input, true)
	var totalTokens int64

	for _, game := range games {
		D := game.ButtonA[0]*game.ButtonB[1] - game.ButtonB[0]*game.ButtonA[1]
		if D == 0 {
			continue
		}

		aNumerator := game.Prize[0]*game.ButtonB[1] - game.ButtonB[0]*game.Prize[1]
		bNumerator := game.ButtonA[0]*game.Prize[1] - game.Prize[0]*game.ButtonA[1]

		if aNumerator%D != 0 || bNumerator%D != 0 {
			continue
		}

		a := aNumerator / D
		b := bNumerator / D

		if a > 0 && b > 0 {
			totalTokens += a*3 + b
		}
	}

	return totalTokens
}

func parseInput(input string, addOffset bool) []game {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var games []game

	for i := 0; i < len(lines); i += 4 {
		aParts := strings.Split(strings.TrimPrefix(lines[i], "Button A: X+"), ", Y+")
		bParts := strings.Split(strings.TrimPrefix(lines[i+1], "Button B: X+"), ", Y+")
		prizeParts := strings.Split(strings.TrimPrefix(lines[i+2], "Prize: X="), ", Y=")

		ButtonA := []int64{parseInt64(aParts[0]), parseInt64(aParts[1])}
		ButtonB := []int64{parseInt64(bParts[0]), parseInt64(bParts[1])}
		Prize := []int64{parseInt64(prizeParts[0]), parseInt64(prizeParts[1])}

		if addOffset {
			Prize[0] += 1e13
			Prize[1] += 1e13
		}

		games = append(games, game{ButtonA: ButtonA, ButtonB: ButtonB, Prize: Prize})
	}

	return games
}

func parseInt64(s string) int64 {
	val, _ := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	return val
}
