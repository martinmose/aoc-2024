package main

import (
	"aoc_2024/solutions/day7"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// day1.Run()
	// day2.Run()
	// day3.Run()
	// day4.Run()
	// day5.Run()
	// day6.Run()
	day7.Run()
}
