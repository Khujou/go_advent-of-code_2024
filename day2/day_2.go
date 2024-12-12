package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const INPUT_FILE = "input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeInput2DArr(f *os.File) [][]int {
	inputArr := make([][]int, 0, 2)
	err := error(nil)

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		inputArr = append(inputArr, make([]int, len(line)))
		for j := 0; j < len(line); j++ {
			inputArr[i][j], err = strconv.Atoi(line[j])
			check(err)
		}
		i++
	}
	return inputArr
}

func isReportSafe(report []int) bool {
	isSortedAscending := slices.IsSorted(report)
	isSortedDescending := slices.IsSortedFunc(report, func(a, b int) int { return b - a })
	for i := 1; i < len(report); i++ {
		if !(isSortedAscending || isSortedDescending) {
			return false
		}
		difference := findDifference(report[i], report[i-1])
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}

func isReportSafeWithDampener(report []int) bool {
	if isReportSafe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		copy := slices.Delete(slices.Clone(report), i, i+1)
		//fmt.Println(copy)
		if isReportSafe(copy) {
			return true
		}
	}
	return false
}

func findDifference(a int, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	difference := max(a, b) - min(a, b)
	return difference
}

func main() {
	f, err := os.Open(INPUT_FILE)
	check(err)
	input := makeInput2DArr(f)
	partOneAns := bothParts(input, isReportSafe)
	fmt.Printf("Part One - %d\n", partOneAns)
	partTwoAns := bothParts(input, isReportSafeWithDampener)
	fmt.Printf("Part Two - %d", partTwoAns)
}

func bothParts(input [][]int, callback func([]int) bool) int {
	safeCount := 0
	for i := 0; i < len(input); i++ {
		//fmt.Printf("Report #%d - %d -> \n", i+1, input[i])
		reportSafety := callback(input[i])
		if reportSafety {
			safeCount++
		}
		//fmt.Printf("%t\n", reportSafety)
	}

	return safeCount
}
