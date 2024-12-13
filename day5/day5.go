package day5

import (
	"fmt"
	"github.com/peichels/aoc2024/input"
	"log"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	before int
	after  int
}

func GetSolutionPart1() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day5_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	var rules []rule
	var updates []int
	var updatelines [][]int
	for _, line := range inputData {
		if strings.Contains(line, "|") {
			leftPart := strings.Split(line, "|")[0]
			rightPart := strings.Split(line, "|")[1]
			left, _ := strconv.Atoi(leftPart)
			right, _ := strconv.Atoi(rightPart)
			rules = append(rules, rule{left, right})
		}
		if strings.Contains(line, ",") {
			updates = []int{}
			updatesString := strings.Split(line, ",")
			for _, num := range updatesString {
				numInt, _ := strconv.Atoi(num)
				updates = append(updates, numInt)
			}
			updatelines = append(updatelines, updates)
		}
	}
	var updateLinesCorrect []int
	for u, updateline := range updatelines {
		numRulesCorrect := 0
		numRulesNotApply := 0
		for _, rule := range rules {
			afterFound := slices.Index(updateline, rule.after)
			beforeFound := slices.Index(updateline, rule.before)
			if beforeFound == -1 && afterFound == -1 {
				numRulesNotApply++
				continue
			}
			if (beforeFound == -1 && afterFound != -1) || (beforeFound != -1 && afterFound == -1) {
				numRulesNotApply++
				continue
			}
			if beforeFound < afterFound {
				numRulesCorrect++
				continue
			}
		}
		if numRulesCorrect+numRulesNotApply == len(rules) {
			updateLinesCorrect = append(updateLinesCorrect, u)
		}

	}
	for _, updateLineCorrect := range updateLinesCorrect {
		for i := range updatelines[updateLineCorrect] {
			if i >= len(updatelines[updateLineCorrect])/2 {
				total += uint64(updatelines[updateLineCorrect][i])
				break
			}
		}
	}

	return total, nil
}

func GetSolutionPart2() (uint64, error) {
	inputData, err := input.ReadInputFileAsStringList("./input/files/day5_1.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var total uint64 = 0
	var rules []rule
	var updates []int
	var updatelines [][]int
	for _, line := range inputData {
		if strings.Contains(line, "|") {
			leftPart := strings.Split(line, "|")[0]
			rightPart := strings.Split(line, "|")[1]
			left, _ := strconv.Atoi(leftPart)
			right, _ := strconv.Atoi(rightPart)
			rules = append(rules, rule{left, right})
		}
		if strings.Contains(line, ",") {
			updates = []int{}
			updatesString := strings.Split(line, ",")
			for _, num := range updatesString {
				numInt, _ := strconv.Atoi(num)
				updates = append(updates, numInt)
			}
			updatelines = append(updatelines, updates)
		}
	}
	var updateLinesCorrect []int
	var updateLinesInCorrect []int
	for u, updateline := range updatelines {
		numRulesCorrect := 0
		numRulesInCorrect := 0
		numRulesNotApply := 0
	rulesLoop:
		for {
			fixPerformed := false
			for _, rule := range rules {
				afterFound := slices.Index(updateline, rule.after)
				beforeFound := slices.Index(updateline, rule.before)
				if beforeFound == -1 || afterFound == -1 {
					numRulesNotApply++
					continue
				}
				if beforeFound < afterFound {
					numRulesCorrect++
					continue
				}
				if beforeFound > afterFound {
					fmt.Println("Incorrect update:", updateline, ", rule:", rule)
					newSlice := make([]int, len(updateline))
					copy(newSlice, updateline)
					aft := updateline[afterFound]
					bef := updateline[beforeFound]
					newSlice[beforeFound] = aft
					newSlice[afterFound] = bef
					fmt.Println("Fixed update:", newSlice)
					copy(updateline, newSlice)
					numRulesInCorrect++
					fixPerformed = true
					continue
				}
			}
			if !fixPerformed {
				break rulesLoop
			}
		}

		if numRulesCorrect+numRulesNotApply == len(rules) {
			updateLinesCorrect = append(updateLinesCorrect, u)
		}
		if numRulesInCorrect > 0 {
			updateLinesInCorrect = append(updateLinesInCorrect, u)
		}
	}
	for _, updateLineCorrect := range updateLinesInCorrect {
		for i := range updatelines[updateLineCorrect] {
			if i >= len(updatelines[updateLineCorrect])/2 {
				total += uint64(updatelines[updateLineCorrect][i])
				break
			}
		}
	}

	return total, nil
}
