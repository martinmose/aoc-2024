package main

import (
	"aoc_2024/solutions/day3"
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
	day3.Run()
}
