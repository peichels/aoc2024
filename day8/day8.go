package day8

import (
	"fmt"
	"github.com/peichels/aoc2024/input"
	"log"
	"math"
	"strings"
)

type coord struct {
	y int
	x int
}

func isValidCoord(pos coord, rows, cols int) bool {
	return pos.y >= 0 && pos.y < rows && pos.x >= 0 && pos.x < cols
}

func GetSolutionPart1() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day8_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	var antinodes []coord
	var matrix [][]string
	for _, lineString := range inputData {
		var line []string
		for _, char := range lineString {
			line = append(line, string(char))
		}
		matrix = append(matrix, line)
	}
	newMatrix := make([][]string, len(matrix))
	for i := range newMatrix {
		newMatrix[i] = make([]string, len(matrix[0]))
		for j := range newMatrix[i] {
			newMatrix[i][j] = "."
		}
	}
	for y, line := range matrix {
		for x, char := range line {
			if char != "." {
				antennaLocation := coord{y, x}
				antennaChar := char
				fmt.Printf("Found %s antenna at %v\n", antennaChar, antennaLocation)
				//searchLoop:
				for searchY := y + 1; searchY < len(matrix); searchY++ {
					for searchX := 0; searchX < len(matrix[searchY]); searchX++ {
						if matrix[searchY][searchX] == antennaChar {
							otherAntennaLocation := coord{searchY, searchX}
							fmt.Printf("Found %s another antenna at %v\n", antennaChar, otherAntennaLocation)
							distanceY := int(math.Sqrt(math.Pow(float64(otherAntennaLocation.y-antennaLocation.y), 2)))
							distanceX := int(math.Sqrt(math.Pow(float64(otherAntennaLocation.x-antennaLocation.x), 2)))
							if antennaLocation.x > otherAntennaLocation.x {
								distanceX = -distanceX
							}
							antinodeLocation1 := coord{antennaLocation.y - distanceY, antennaLocation.x - distanceX}
							antinodeLocation2 := coord{otherAntennaLocation.y + distanceY, otherAntennaLocation.x + distanceX}
							if isValidCoord(antinodeLocation1, len(matrix), len(matrix[0])) {
								antinodes = append(antinodes, antinodeLocation1)
								newMatrix[antinodeLocation1.y][antinodeLocation1.x] = "#"
							}
							if isValidCoord(antinodeLocation2, len(matrix), len(matrix[0])) {
								antinodes = append(antinodes, antinodeLocation2)
								newMatrix[antinodeLocation2.y][antinodeLocation2.x] = "#"
							}
							//break searchLoop
						}
					}
				}
				antinodes = append(antinodes, coord{y, x})
			}
		}
	}

	for _, line := range newMatrix {
		// convert line to string and count #
		noHashes := strings.Count(fmt.Sprintf(strings.Join(line, "")), "#")
		total += uint64(noHashes)

	}
	return total, nil
}

func GetSolutionPart2() (uint64, error) {
	_, err := input.ReadInputFileAsStringList("./input/files/day5_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0

	return total, nil
}
