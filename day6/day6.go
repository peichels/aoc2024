package day6

import (
	"fmt"
	"github.com/peichels/aoc2024/input"
	"log"
)

type coord struct {
	y int
	x int
}

func GetSolutionPart1() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/test_day6_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	var move = map[string]coord{}
	move["U"] = coord{-1, 0}
	move["D"] = coord{1, 0}
	move["L"] = coord{0, -1}
	move["R"] = coord{0, 1}
	var posVisited []coord
	var currentDirection string = "U"
	currentPos := findStart(inputData)
	for {
		if !visited(posVisited, currentPos) {
			posVisited = append(posVisited, currentPos)
		}
		newPos := coord{currentPos.y + move[currentDirection].y, currentPos.x + move[currentDirection].x}
		if string(inputData[newPos.y][newPos.x]) == "#" {
			currentDirection = turnRight(currentDirection)
			continue
		}
		currentPos = newPos
		if currentPos.y == len(inputData)-1 || currentPos.y == 0 || currentPos.x == len(inputData[0])-1 || currentPos.x == 0 {
			if !visited(posVisited, currentPos) {
				posVisited = append(posVisited, currentPos)
			}
			total = uint64(len(posVisited))
			break
		}
	}

	return total, nil
}

func visited(visited []coord, pos coord) bool {
	for _, checkCoord := range visited {
		if checkCoord.y == pos.y && checkCoord.x == pos.x {
			return true
		}
	}
	return false
}

func findStart(data []string) coord {
	for lineNo, line := range data {
		for x, symbol := range line {
			if string(symbol) == "^" {
				return coord{lineNo, x}
			}
		}
	}
	return coord{}
}

func turnRight(currentDirection string) string {
	switch currentDirection {
	case "U":
		return "R"
	case "R":
		return "D"
	case "D":
		return "L"
	case "L":
		return "U"
	}
	return "X"
}

func GetSolutionPart2() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/test_day6_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	startPos := findStart(inputData)
	possiblePositions := findPossibleObstructionPositions(inputData, startPos)
	fmt.Println("Number of possible positions:", len(possiblePositions))
	return uint64(len(possiblePositions)), err
}

func findPossibleObstructionPositions(grid []string, startPos coord) []coord {
	var possiblePositions []coord
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '.' && (y != startPos.y || x != startPos.x) {
				if causesLoop(grid, startPos, coord{y, x}) {
					possiblePositions = append(possiblePositions, coord{y, x})
				}
			}
		}
	}
	return possiblePositions
}

func causesLoop(grid []string, startPos, obstruction coord) bool {
	var move = map[string]coord{}
	move["U"] = coord{-1, 0}
	move["D"] = coord{1, 0}
	move["L"] = coord{0, -1}
	move["R"] = coord{0, 1}
	grid[obstruction.y] = grid[obstruction.y][:obstruction.x] + "#" + grid[obstruction.y][obstruction.x+1:]
	defer func() {
		grid[obstruction.y] = grid[obstruction.y][:obstruction.x] + "." + grid[obstruction.y][obstruction.x+1:]
	}()

	currentPos := startPos
	currentDirection := "U"
	visited := map[coord]bool{}

	for {
		if visited[currentPos] {
			return true
		}
		visited[currentPos] = true

		newPos := coord{currentPos.y + move[currentDirection].y, currentPos.x + move[currentDirection].x}
		if grid[newPos.y][newPos.x] == '#' {
			currentDirection = turnRight(currentDirection)
			continue
		}
		currentPos = newPos
		if currentPos.y == len(grid)-1 || currentPos.y == 0 || currentPos.x == len(grid[0])-1 || currentPos.x == 0 {
			return false
		}
	}
}
