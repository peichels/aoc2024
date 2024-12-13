package day11

import (
	"github.com/peichels/aoc2024/input"
	"log"
	"math"
	"strconv"
	"strings"
)

func GetSolutionPart1() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day11_1.txt")
	arrangement := strings.Split(inputData[0], " ")
	stoneList := []int{}
	for _, item := range arrangement {
		number, _ := strconv.Atoi(item)
		stoneList = append(stoneList, number)
	}
	total := reArrangeInt(stoneList, 75)
	return uint64(total), err
}

func powTen(pow int) int {
	n := 1
	for range pow {
		n *= 10
	}
	return n
}

func reArrangeInt(arrangement []int, times int) int {
	newArrangementMap := make(map[int]int)
	for _, value := range arrangement {
		newArrangementMap[value]++
	}
	for i := 0; i < times; i++ {
		newStoneMap := make(map[int]int)
		for key, count := range newArrangementMap {
			if key == 0 {
				newStoneMap[1] += count
			} else if digits := numDigits(key); digits%2 == 0 {
				filter := powTen(digits / 2)
				left, right := key/filter, key%filter
				newStoneMap[left] += count
				newStoneMap[right] += count
			} else {
				newStoneMap[key*2024] += count
			}
		}
		newArrangementMap = newStoneMap
	}
	result := 0
	for _, count := range newArrangementMap {
		result += count
	}
	return result
}

func numDigits(item int) int {
	return (int(math.Log10(float64(item)))) + 1
}

func reArrange(arrangement []string) []string {
	newArrangement := make([]string, 0, len(arrangement)*2) // Pre-allocate space
	for _, item := range arrangement {
		if item == "0" {
			newArrangement = append(newArrangement, "1")
		} else if len(item)%2 == 0 {
			left, _ := strconv.Atoi(item[:len(item)/2])
			right, _ := strconv.Atoi(item[len(item)/2:])
			newArrangement = append(newArrangement, strconv.Itoa(left), strconv.Itoa(right))
		} else {
			number, _ := strconv.Atoi(item)
			newArrangement = append(newArrangement, strconv.Itoa(number*2024))
		}
	}
	return newArrangement
}

func GetSolutionPart2() (uint64, error) {
	_, err := input.ReadInputFileAsStringList("./input/files/day5_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0

	return total, nil
}
