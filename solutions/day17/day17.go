package day17

import (
	"aoc_2024/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Run() error {
	fmt.Println("Day 17:")

	dayPath := "17/input"
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
	a, b, c, program := parseRegisters(input)
	output := runProgram(a, b, c, program)
	return joinInts(output, ",")
}

func part2Puzzle(input string) string {
	a, b, c, program := parseRegisters(input)
	a = 0

	for pos := len(program) - 1; pos >= 0; pos-- {
		a <<= 3 // Shift left by 3 bits
		for !slices.Equal(runProgram(a, b, c, program), program[pos:]) {
			a++
		}
	}

	return strconv.Itoa(a)
}

func parseRegisters(input string) (int, int, int, []int) {
	var a, b, c int
	var program []int

	for _, line := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(line, "Register A:"):
			a = parseInt(strings.TrimSpace(line[12:]))
		case strings.HasPrefix(line, "Register B:"):
			b = parseInt(strings.TrimSpace(line[12:]))
		case strings.HasPrefix(line, "Register C:"):
			c = parseInt(strings.TrimSpace(line[12:]))
		case strings.HasPrefix(line, "Program:"):
			program = extractInts(strings.TrimSpace(line[9:]))
		}
	}

	return a, b, c, program
}

func runProgram(a, b, c int, program []int) []int {
	var out []int

	for ip := 0; ip < len(program); ip += 2 {
		opcode, operand := program[ip], program[ip+1]

		// Process operand
		value := operand
		switch operand {
		case 4:
			value = a
		case 5:
			value = b
		case 6:
			value = c
		}

		// Execute instruction
		switch opcode {
		case 0: // ADV - Divide A by 2^value
			a >>= value
		case 1: // BXL - XOR B with literal
			b ^= operand
		case 2: // BST - Set B to value mod 8
			b = value % 8
		case 3: // JNZ - Jump if A is not zero
			if a != 0 {
				ip = operand - 2
			}
		case 4: // BXC - XOR B with C
			b ^= c
		case 5: // OUT - Output value mod 8
			out = append(out, value%8)
		case 6: // BDV - Divide A by 2^value, store in B
			b = a >> value
		case 7: // CDV - Divide A by 2^value, store in C
			c = a >> value
		}
	}

	return out
}

func extractInts(s string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(s, -1)
	ints := make([]int, len(matches))
	for i, match := range matches {
		ints[i], _ = strconv.Atoi(match)
	}
	return ints
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func joinInts(arr []int, sep string) string {
	strNums := make([]string, len(arr))
	for i, num := range arr {
		strNums[i] = strconv.Itoa(num)
	}
	return strings.Join(strNums, sep)
}
