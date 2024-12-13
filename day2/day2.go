package day2

import (
	"github.com/peichels/aoc2024/input"
	"log"
	"slices"
	"strconv"
	"strings"
)

type location struct {
	x int
	y int
}

func GetSolutionPart1() (int, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day2_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	totalReportsSafe := 0
	for _, report := range inputData {
		levelsString := strings.Split(report, " ")
		var levels []int
		for _, levelString := range levelsString {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			levels = append(levels, level)
		}
		if isReportSafe(levels) {
			totalReportsSafe++
		}
	}
	return totalReportsSafe, nil
}

func isReportSafe(levels []int) bool {
	levelIncrease := false
	reportIsSafe := false
	if levels[1] > levels[0] {
		levelIncrease = true
	}
levelLoop:
	for i, lvl := range levels {
		if i >= 1 {
			if lvl > levels[i-1] {
				if !levelIncrease {
					reportIsSafe = false
					//fmt.Printf("Unsafe %v at: %d\n", levels, levels[i])
					break levelLoop
				}
				if lvl-levels[i-1] <= 3 && lvl-levels[i-1] >= 1 {
					reportIsSafe = true
				} else {
					reportIsSafe = false
					//fmt.Printf("Unsafe %v at: %d\n", levels, levels[i])
					break levelLoop
				}
			}

			if lvl < levels[i-1] {
				if levelIncrease {
					reportIsSafe = false
					//fmt.Printf("Unsafe %v at: %d\n", levels, levels[i])
					break levelLoop
				}
				if lvl-levels[i-1] >= -3 && lvl-levels[i-1] <= -1 {
					reportIsSafe = true
				} else {
					reportIsSafe = false
					//fmt.Printf("Unsafe %v at: %d\n", levels, levels[i])
					break levelLoop
				}
			}
			if lvl == levels[i-1] {
				reportIsSafe = false
				//fmt.Printf("Unsafe %v at: %d\n", levels, levels[i])
				break levelLoop
			}

		}
	}
	return reportIsSafe
}

func GetSolutionPart2() (int, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day2_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	totalReportsSafe := 0
	for lineNo, report := range inputData {
		levelsString := strings.Split(report, " ")
		var levels []int
		for _, levelString := range levelsString {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			levels = append(levels, level)
		}

		if isReportSafe(levels) {
			totalReportsSafe++
		} else {

			for i := range levels {
				newSlice := make([]int, len(levels))
				copy(newSlice, levels)

				slices.Delete(newSlice, i, i+1)
				newSlice = newSlice[:len(newSlice)-1]

				if isReportSafe(newSlice) {
					log.Printf("line: %d, original: %v, fixed: %v\n", lineNo, levels, newSlice)
					totalReportsSafe++
					break
				} else {
					continue
				}
			}

		}

	}

	return totalReportsSafe, nil
}
