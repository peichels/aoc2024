package main

import (
	"fmt"
	"github.com/peichels/aoc2024/day7"
	"time"
)

func main() {
	startTime := time.Now()
	solution1, _ := day7.GetSolutionPart1()
	fmt.Println(solution1)
	endTime := time.Now()
	fmt.Println("Duration: ", endTime.Sub(startTime))

	//startTime = time.Now()
	//solution2, _ := day12.GetSolutionPart2()
	//fmt.Println(solution2)
	//endTime = time.Now()
	//fmt.Println("Duration: ", endTime.Sub(startTime))
}
