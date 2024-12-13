package day1

import (
	"fmt"
	"github.com/peichels/aoc2024/input"
	"math"
	"slices"
	"strconv"
	"strings"
)

type location struct {
	x int
	y int
}

func GetSolutionPart1() (int, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day1_1.txt")
	var listOne, listTwo []int
	for _, row := range inputData {
		splitted := strings.SplitN(row, " ", 3)
		leftNum, _ := strconv.Atoi(splitted[0])
		rightNum, _ := strconv.Atoi(strings.TrimSpace(splitted[len(splitted)-1]))
		listOne = append(listOne, leftNum)
		listTwo = append(listTwo, rightNum)
	}
	slices.Sort(listOne)
	slices.Sort(listTwo)
	totalDistance := 0
	distance := 0
	for i, left := range listOne {
		distance = int(math.Abs(float64(left - listTwo[i])))
		totalDistance += distance
	}
	if err != nil {
		return 0, err
	}
	fmt.Println(totalDistance)
	return totalDistance, nil
}

func GetSolutionPart2() (int, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day1_1.txt")
	var listOne, listTwo []int
	for _, row := range inputData {
		splitted := strings.SplitN(row, " ", 3)
		leftNum, _ := strconv.Atoi(splitted[0])
		rightNum, _ := strconv.Atoi(strings.TrimSpace(splitted[len(splitted)-1]))
		listOne = append(listOne, leftNum)
		listTwo = append(listTwo, rightNum)
	}
	slices.Sort(listOne)
	slices.Sort(listTwo)
	totalDistance := 0
	occurredN := 0
	for _, left := range listOne {
		occurredN = 0
		for _, right := range listTwo {
			if left == right {
				occurredN++
			}
		}
		totalDistance += occurredN * left
	}
	if err != nil {
		return 0, err
	}
	fmt.Println(totalDistance)
	return totalDistance, nil
}
