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
	inputData, err := input.ReadInputFileAsStringList("./input/files/day7_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	var result uint64 = 0
	var lines []line
	//var unreachableLines []line
	for _, lineString := range inputData {
		outcome, _ := strconv.Atoi(strings.Split(lineString, ":")[0])
		numbers := strings.Split(lineString, " ")[1:]
		var numbersInt []int32
		for _, number := range numbers {
			numberInt, _ := strconv.Atoi(number)
			numbersInt = append(numbersInt, int32(numberInt))
		}
		//timesTotal := 1
		//addTotal := 0
		//for _, number := range numbersInt {
		//	timesTotal *= int(number)
		//	addTotal += int(number)
		//}
		//if timesTotal < outcome {
		//	fmt.Println("Can not reach outcome for Line: ", lineString, ": ", timesTotal)
		//	unreachableLines = append(unreachableLines, line{uint64(outcome), numbersInt})
		//	continue
		//}
		//
		//if addTotal > outcome {
		//	fmt.Println("Sum is larger than outcome for Line: ", lineString, ": ", addTotal)
		//	unreachableLines = append(unreachableLines, line{uint64(outcome), numbersInt})
		//	continue
		//}
		lines = append(lines, line{uint64(outcome), numbersInt})

	}

	for _, line := range lines {
		permutations := math.Pow(2, float64(len(line.numbers)-1))
		fmt.Println(line.numbers)
		fmt.Println("Permutations: ", permutations)

		// iterate over all permutations
		for perm := range int(permutations) {

			for index, number := range line.numbers {
				if index == 0 {
					total = uint64(number)
				}
				if index == len(line.numbers)-1 {
					break
				}
				if perm&(1<<index) != 0 {
					total *= uint64(line.numbers[index+1])
					fmt.Printf("%d * %d,", number, line.numbers[index+1])
				} else {
					total += uint64(line.numbers[index+1])
					fmt.Printf("%d + %d,", number, line.numbers[index+1])
				}

			}
			if total == line.outcome {
				fmt.Println(" == ", total)
				result += total
				break
			}
			fmt.Println(" = ", total)

		}
		// 0: x[0] 0 x[1]
		// 1: x[0] 1 x[1]

		// 00: x[0] 0 x[1] 0 x[2]
		// 01: x[0] 0 x[1] 1 x[2]
		// 10: x[0] 1 x[1] 0 x[2]
		// 11: x[0] 1 x[1] 1 x[2]

		// 000: x[0] 0 x[1] 0 x[2] 0 x[3]
		// 001: x[0] 0 x[1] 0 x[2] 1 x[3]
		// 010: x[0] 0 x[1] 1 x[2] 0 x[3]
		// 011: x[0] 0 x[1] 1 x[2] 1 x[3]
		// 100: x[0] 1 x[1] 0 x[2] 0 x[3]
		// 101: x[0] 1 x[1] 0 x[2] 1 x[3]
		// 110: x[0] 1 x[1] 1 x[2] 0 x[3]

	}

	return result, nil
}

func GetSolutionPart2() (uint64, error) {
	_, err := input.ReadInputFileAsStringList("./input/files/day5_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0

	return total, nil
}
