package main

import (
	"aoc_2024/solutions/day24"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	fmt.Println("Advent of Code 2024")

	days := []struct {
		name string
		run  func() error
	}{
		// {"Day 1", day1.Run},
		// {"Day 2", day2.Run},
		// {"Day 3", day3.Run},
		// {"Day 4", day4.Run},
		// {"Day 5", day5.Run},
		// {"Day 6", day6.Run},
		// {"Day 7", day7.Run},
		// {"Day 8", day8.Run},
		// {"Day 9", day9.Run},
		// {"Day 10", day10.Run},
		// {"Day 11", day11.Run},
		// {"Day 12", day12.Run},
		// {"Day 13", day13.Run},
		// {"Day 14", day14.Run},
		// {"Day 15", day15.Run},
		// {"Day 16", day16.Run},
		// {"Day 17", day17.Run},
		// {"Day 18", day18.Run},
		// {"Day 19", day19.Run},
		// {"Day 20", day20.Run},
		// {"Day 21", day21.Run},
		// {"Day 22", day22.Run},
		// {"Day 23", day23.Run},
		{"Day 24", day24.Run},
		// {"Day 25", day25.Run},
	}

	for _, day := range days {
		fmt.Printf("Running %s...\n", day.name)
		if err := day.run(); err != nil {
			fmt.Printf("Error in %s: %v\n", day.name, err)
		} else {
			fmt.Printf("%s completed successfully.\n", day.name)
		}
	}
}
