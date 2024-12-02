[![Checks](https://github.com/martinmose/aoc-2024/actions/workflows/checks.yml/badge.svg?branch=main)](https://github.com/martinmose/aoc-2024/api/actions/workflows/checks.yml)

# Advent of Code 2024

This is my repository for solving the puzzles from [Advent of Code 2024](https://adventofcode.com/2024).

This year, I decided to give Go a spin to get more familiar with the language. It might feel like I've gone a bit overboard with the structure, almost making it too "corporate," but I wanted it to feel like a real project.

## Up and Running

1. Add a `.env` file at the root of the project.
2. Add the following line to the `.env` file:

   `AOC_SESSION=XXX`

   Replace `XXX` with the session cookie from the Advent of Code website.

## Run the program

Adjust inside `main.go` what day it should run.

Get the results with: `go run main.go`

## Code Formatting

To ensure consistent code quality and style, here are the tools I'm using:

### `go fmt`

Formats the code according to Go's standards.  
Run it with: `go fmt ./...`

### `go vet`

Performs static analysis to catch common mistakes.  
Run it with: `go vet ./...`

### `golint`

Provides linting to enforce Go code style.  
Install it with: `go install golang.org/x/lint/golint@latest`  
Run it with: `golint ./...`

### `staticcheck`

Offers advanced linting and catches performance and correctness issues.  
Install it with: `go install honnef.co/go/tools/cmd/staticcheck@latest`  
Run it with: `staticcheck ./...`

## Tests

Run all tests with: `go test ./...`
