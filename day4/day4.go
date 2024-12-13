package day4

import (
	"github.com/peichels/aoc2024/input"
	"log"
	"strings"
)

func GetSolutionPart1() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day4_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	findWord := "XMAS"
	findWordReversed := "SAMX"
	// horizontal check
	for _, line := range inputData {
		total += uint64(strings.Count(line, findWord))
		total += uint64(strings.Count(line, findWordReversed))
	}
	// vertical check
	for x := 0; x < len(inputData); x++ {
		vertLineString := ""
		for y := 0; y < len(inputData); y++ {
			vertLineString += string(inputData[y][x])
		}
		total += uint64(strings.Count(vertLineString, findWord))
		total += uint64(strings.Count(vertLineString, findWordReversed))
	}
	// diagonal check forward slash 1st half
	for x := 3; x < len(inputData[0]); x++ {
		forthDiagString := ""
		for y := 0; y <= x; y++ {
			forthDiagString += string(inputData[y][x-y])
		}
		total += uint64(strings.Count(forthDiagString, findWord))
		total += uint64(strings.Count(forthDiagString, findWordReversed))
	}
	// diagonal check forward slash 2nd half
	for x := len(inputData[0]) - 4; x > 0; x-- {
		forthDiagString := ""
		for y := len(inputData) - 1; y >= x; y-- {
			forthDiagString += string(inputData[y][x+(len(inputData))-y-1])
		}
		total += uint64(strings.Count(forthDiagString, findWord))
		total += uint64(strings.Count(forthDiagString, findWordReversed))
	}
	// diagonal check backslash 1st half
	for x := len(inputData[0]) - 4; x > 0; x-- {
		backDiagString := ""
		for y := 0; y+x < len(inputData); y++ {
			backDiagString += string(inputData[y][x+y])
		}
		total += uint64(strings.Count(backDiagString, findWord))
		total += uint64(strings.Count(backDiagString, findWordReversed))
	}
	// diagonal check backslash 2nd half
	for x := 3; x < len(inputData[0]); x++ {
		backDiagString := ""
		for y := len(inputData) - 1; y >= len(inputData)-1-x; y-- {
			backDiagString += string(inputData[y][y-(len(inputData))+x+1])
		}
		total += uint64(strings.Count(backDiagString, findWord))
		total += uint64(strings.Count(backDiagString, findWordReversed))
	}

	return total, nil
}

func GetSolutionPart2() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day4_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	for x := 1; x < len(inputData[0])-1; x++ {
		for y := 1; y < len(inputData)-1; y++ {
			if string(inputData[y][x]) == "A" {
				if (string(inputData[y-1][x-1]) == "M" && string(inputData[y+1][x+1]) == "S") || (string(inputData[y-1][x-1]) == "S" && string(inputData[y+1][x+1]) == "M") {
					if (string(inputData[y+1][x-1]) == "M" && string(inputData[y-1][x+1]) == "S") || (string(inputData[y+1][x-1]) == "S" && string(inputData[y-1][x+1]) == "M") {
						total++
					}
				}
			}
		}
	}

	return total, nil
}
