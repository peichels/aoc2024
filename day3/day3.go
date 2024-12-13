package day3

import (
	"fmt"
	"github.com/peichels/aoc2024/input"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func GetSolutionPart1() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day3_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	checkString := ""
	for x := 0; x < len([]rune(inputData[0]))-11; x++ {
		if string(inputData[0][x])+string(inputData[0][x+1])+string(inputData[0][x+2]) == "mul" {
			checkString = string([]rune(inputData[0])[x : x+12])
			match, err := regexp.MatchString("mul\\([0-9]+,[0-9]+\\)", checkString)
			if err != nil {
				fmt.Println("Error:", err)
			}

			if match {
				total += uint64(calculateMul(checkString))
			}
		}

	}
	return total, nil
}

func calculateMul(checkString string) uint32 {
	fmt.Println("String: " + checkString)
	left := strings.Split(strings.Split(checkString, ",")[0], "(")[1]
	right := strings.Split(strings.Split(checkString, ",")[1], ")")[0]
	leftInt, leftErr := strconv.Atoi(left)
	rightInt, rightErr := strconv.Atoi(right)
	if leftErr != nil || rightErr != nil {
		fmt.Println(left + " " + right)
		return 0
	}
	return uint32(leftInt * rightInt)
}

func GetSolutionPart2() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day3_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	checkString := ""
	include := true
	for x := 0; x < len([]rune(inputData[0]))-11; x++ {
		if string(inputData[0][x])+string(inputData[0][x+1])+string(inputData[0][x+2])+string(inputData[0][x+3])+string(inputData[0][x+4]) == "don't" {
			include = false
			continue
		} else if string(inputData[0][x])+string(inputData[0][x+1]) == "do" {
			include = true
			continue
		}
		if string(inputData[0][x])+string(inputData[0][x+1])+string(inputData[0][x+2]) == "mul" && include {
			checkString = string([]rune(inputData[0])[x : x+12])
			//match, err := regexp.MatchString("mul[(0-9,0-9)]+", checkString)
			match, err := regexp.MatchString("mul\\([0-9]+,[0-9]+\\)", checkString)
			if err != nil {
				fmt.Println("Error:", err)
			}

			//if match && strings.Contains(checkString, ",") && strings.Contains(checkString, ")") {
			if match {
				total += uint64(calculateMul(checkString))
			}
		}

	}
	return total, nil
}
