package main

import (
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

func splitInput(dat []string) ([]int, []int) {

	sideLen := len(dat) / 2

	left := make([]int, sideLen)
	right := make([]int, sideLen)

	err := error(nil)

	for i := 0; i < len(dat); i++ {
		idx := i / 2
		if i%2 == 0 {
			left[idx], err = strconv.Atoi(dat[i])
		} else {
			right[idx], err = strconv.Atoi(dat[i])
		}
		check(err)
	}

	return left, right
}

/*
	3   4	- Part One : 11
	4   3	- Part Two : 31
	2   5
	1   3
	3   9
	3   3
*/

func main() {
	dat, err := os.ReadFile(INPUT_FILE)
	check(err)
	inputArr := strings.Fields(string(dat))
	partOneAns := partOne(inputArr)
	fmt.Printf("Part One - %d\n", partOneAns)

	partTwoAns := partTwo(inputArr)
	fmt.Printf("Part Two = %d", partTwoAns)

}

func partOne(dat []string) int {

	sideLen := len(dat) / 2

	left, right := splitInput(dat)

	slices.Sort(left)
	slices.Sort(right)

	res := 0

	for i := 0; i < sideLen; i++ {
		leftInt := left[i]
		rightInt := right[i]
		if leftInt < 0 {
			leftInt *= -1
		}
		if rightInt < 0 {
			rightInt *= -1
		}
		res += max(leftInt, rightInt) - min(leftInt, rightInt)
	}

	return res
}

func partTwo(dat []string) int {
	sideLen := len(dat) / 2

	left, right := splitInput(dat)

	res := 0

	for i := 0; i < sideLen; i++ {
		leftInt := left[i]
		rightCount := 0
		for i := 0; i < sideLen; i++ {
			if right[i] == leftInt {
				rightCount++
			}
		}

		res += leftInt * rightCount
	}

	return res
}
