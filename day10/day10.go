package day10

import (
	"fmt"
	"github.com/peichels/aoc2024/input"
	"log"
	"slices"
	"strconv"
)

type coord struct {
	y int
	x int
}

func isValidCoord(pos coord, rows, cols int) bool {
	return pos.y >= 0 && pos.y < rows && pos.x >= 0 && pos.x < cols
}

func GetSolutionPart1() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/test_day1-_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	var firstDiskLayout []string
	for i, karakter := range inputData[0] {
		if i%2 == 0 {
			// fileblocks
			number, _ := strconv.Atoi(string(karakter))
			for x := 0; x < number; x++ {
				firstDiskLayout = append(firstDiskLayout, strconv.Itoa((i / 2)))
			}
		} else {
			// freespace
			number, _ := strconv.Atoi(string(karakter))
			for x := 0; x < number; x++ {
				firstDiskLayout = append(firstDiskLayout, ".")
			}
		}
	}
	startPoint := slices.Index(firstDiskLayout, ".")

	endPoint := len(firstDiskLayout) - 1
	diskLayout := firstDiskLayout
	lastDiskLayout := defrag(startPoint, endPoint, diskLayout)

	for i, karakter := range lastDiskLayout {
		if karakter == "." {
			//49433*5246
			break
		}
		number, _ := strconv.Atoi(karakter)
		total += uint64(number) * uint64(i)
	}
	fmt.Println(lastDiskLayout)
	return total + uint64(49433*5246), nil
}

func defrag(start int, end int, layout []string) []string {

	if start >= end {
		return layout
	}

	if layout[start] != "." {
		for {
			start++
			if layout[start] == "." {
				break
			}
		}
	}
	if layout[end] != "." {
		layout[start] = layout[end]
		layout[end] = "."
	} else {
		for {
			end--
			if layout[end] != "." {
				break
			}
		}
		layout[start] = layout[end]
		layout[end] = "."
	}
	//fmt.Println(layout)
	defrag(start+1, end-1, layout)
	return layout
}

func GetSolutionPart2() (uint64, error) {
	_, err := input.ReadInputFileAsStringList("./input/files/day5_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0

	return total, nil
}
