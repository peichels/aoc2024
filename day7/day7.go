package day7

import (
	"fmt"
	"github.com/peichels/aoc2024/input"
	"log"
	"math"
	"strconv"
	"strings"
)

type line struct {
	outcome uint64
	numbers []int32
}

func GetSolutionPart1() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/test_day7_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	var lines []line
	var unreachableLines []line
	for _, lineString := range inputData {
		outcome, _ := strconv.Atoi(strings.Split(lineString, ":")[0])
		numbers := strings.Split(lineString, " ")[1:]
		var numbersInt []int32
		for _, number := range numbers {
			numberInt, _ := strconv.Atoi(number)
			numbersInt = append(numbersInt, int32(numberInt))
		}
		timesTotal := 1
		addTotal := 0
		for _, number := range numbersInt {
			timesTotal *= int(number)
			addTotal += int(number)
		}
		if timesTotal < outcome {
			fmt.Println("Can not reach outcome for Line: ", lineString, ": ", timesTotal)
			unreachableLines = append(unreachableLines, line{uint64(outcome), numbersInt})
			continue
		}

		if addTotal > outcome {
			fmt.Println("Sum is larger than outcome for Line: ", lineString, ": ", addTotal)
			unreachableLines = append(unreachableLines, line{uint64(outcome), numbersInt})
			continue
		}
		lines = append(lines, line{uint64(outcome), numbersInt})

	}

	for _, line := range lines {
		permutations := math.Pow(2, float64(len(line.numbers)-1))
		fmt.Println(line.numbers)
		fmt.Println("Permutations: ", permutations)
		if line.outcome == 292 {
			fmt.Println("292")
		}
		for i := 0; i < int(permutations); i++ {
			// i = 0 len(line.numbers)-1 = 3 000
			// i = 1 len(line.numbers)-1 = 3 001
			// i = 2 len(line.numbers)-1 = 3 010
			// i = 3 len(line.numbers)-1 = 3 011
			// i = 4 len(line.numbers)-1 = 3 100
			// i = 5 len(line.numbers)-1 = 3 101
			// i = 6 len(line.numbers)-1 = 3 110
			// i = 7 len(line.numbers)-1 = 3 111
			var localTotal uint64 = 0
			for j, number := range line.numbers {
				if j == 0 {
					localTotal = uint64(number)
					continue
				}
				if i&(1<<j) != 0 {
					localTotal += uint64(number)
				} else {
					localTotal *= uint64(number)
				}
			}
			if localTotal == line.outcome {
				total += localTotal
				break
			} else if localTotal > line.outcome {
				continue
			}
		}
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
