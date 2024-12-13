package day12

import (
	"github.com/peichels/aoc2024/input"
	"log"
)

type coord struct {
	y int
	x int
}

var directions = []coord{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

func GetSolutionPart1() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day12_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	plots := make([][]rune, len(inputData))
	for i, line := range inputData {
		plots[i] = []rune(line)
	}
	visited := make(map[coord]bool)
	totalPrice := 0

	for y := 0; y < len(plots); y++ {
		for x := 0; x < len(plots[0]); x++ {
			if !visited[coord{y, x}] {
				area, perimeter := findRegion(coord{y, x}, plots, visited)
				totalPrice += area * perimeter
			}
		}
	}

	return uint64(totalPrice), nil
}

func findRegion(start coord, plots [][]rune, visited map[coord]bool) (int, int) {
	plotList := []coord{start}
	plantType := plots[start.y][start.x]
	area := 0
	perimeter := 0

	for len(plotList) > 0 {
		current := plotList[len(plotList)-1]
		plotList = plotList[:len(plotList)-1]

		if visited[current] {
			continue
		}
		visited[current] = true
		area++

		for _, dir := range directions {
			neighbor := coord{current.y + dir.y, current.x + dir.x}
			if inLimits(neighbor, len(plots), len(plots[0])) {
				if plots[neighbor.y][neighbor.x] == plantType {
					if !visited[neighbor] {
						plotList = append(plotList, neighbor)
					}
				} else {
					// other planttype, so fence
					perimeter++
				}
			} else {
				// on the border, so fence
				perimeter++
			}
		}
	}

	return area, perimeter
}

func inLimits(c coord, height, width int) bool {
	return c.y >= 0 && c.y < height && c.x >= 0 && c.x < width
}
func GetSolutionPart2() (uint64, error) {
	_, err := input.ReadInputFileAsStringList("./input/files/day12_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0

	return total, nil
}
