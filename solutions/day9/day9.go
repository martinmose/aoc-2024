package day9

import (
	"aoc_2024/utils"
	"fmt"
	"strconv"
	"strings"
)

type file struct {
	id   int
	size int
}

// Run runs the day 9 challenge
func Run() error {
	fmt.Println("Day 9:")

	dayPath := "9/input"
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
	blocks := parseDiskMapPart1(input)
	compactedBlocks := compactDiskPart1(blocks)
	return calculateChecksum(compactedBlocks)
}

func part2Puzzle(input string) int {
	files := parseDiskMapPart2(input)
	compactedFiles := compactDiskPart2(files)
	return calculateChecksumPart2(compactedFiles)
}

func parseDiskMapPart1(input string) []int {
	input = strings.TrimSpace(input)
	var blocks []int
	fileID := 0

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(string(input[i]))
		if i%2 == 0 {
			for j := 0; j < value; j++ {
				blocks = append(blocks, fileID)
			}
			fileID++
		} else {
			for j := 0; j < value; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	return blocks
}

func compactDiskPart1(blocks []int) []int {
	readIndex := len(blocks) - 1

	for i := 0; i < len(blocks); i++ {
		if readIndex <= i {
			break
		}
		if blocks[i] > -1 {
			continue
		}
		for readIndex >= 0 {
			if blocks[readIndex] > -1 {
				blocks[i] = blocks[readIndex]
				blocks[readIndex] = -1
				readIndex--
				break
			}
			readIndex--
		}
	}

	return blocks
}

func calculateChecksum(blocks []int) int {
	checksum := 0
	for position, fileID := range blocks {
		if fileID != -1 {
			checksum += position * fileID
		}
	}
	return checksum
}

func parseDiskMapPart2(input string) []file {
	input = strings.TrimSpace(input)
	var files []file
	fileID := 0

	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(string(input[i]))
		if value == 0 {
			continue
		}
		if i%2 == 0 {
			files = append(files, file{id: fileID, size: value})
			fileID++
		} else {
			files = append(files, file{id: -1, size: value})
		}
	}

	return files
}

func compactDiskPart2(files []file) []file {
	for i := 0; i < len(files); i++ {
		if files[i].id > -1 {
			continue
		}
		for j := len(files) - 1; j > i; j-- {
			if files[j].id == -1 {
				continue
			}
			sizeDiff := files[i].size - files[j].size
			if sizeDiff >= 0 {
				files[i] = files[j]
				files[j].id = -1
				if sizeDiff > 0 {
					files = append(files[:i+1], append([]file{{id: -1, size: sizeDiff}}, files[i+1:]...)...)
				}
				break
			}
		}
	}
	return files
}

func calculateChecksumPart2(files []file) int {
	checksum := 0
	position := 0

	for _, file := range files {
		for i := 0; i < file.size; i++ {
			if file.id > -1 {
				checksum += position * file.id
			}
			position++
		}
	}

	return checksum
}
