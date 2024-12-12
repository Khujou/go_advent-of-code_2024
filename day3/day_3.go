package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const INPUT_FILE = "input.txt"

const GET_MUL = `mul\(([0-9]+),([0-9]+)\)` // Find's a substring that matches 'mul(#,#)'
const GET_DO = `do\(\)`
const GET_DONT = `don't\(\)`

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func multiplyTwoStrings(a string, b string) int {
	aNum, err := strconv.Atoi(a)
	check(err)
	bNum, err := strconv.Atoi(b)
	check(err)
	return aNum * bNum
}

func main() {
	dat, err := os.ReadFile(INPUT_FILE)
	check(err)
	input := string(dat)
	partOneAns := partOne(input)
	fmt.Printf("Part One - %d\n", partOneAns)

	partTwoAns := partTwo(input)
	fmt.Printf("Part Two - %d", partTwoAns)
}

func partTwo(input string) int {

	find := regexp.MustCompile(fmt.Sprintf("(%s|%s|%s)", GET_MUL, GET_DO, GET_DONT))

	//fmt.Println(find.String())

	validSeqs := find.FindAllString(input, -1)

	//fmt.Println(validSeqs)

	find = regexp.MustCompile(`([0-9]+)`)
	conditional := regexp.MustCompile(fmt.Sprintf("(%s|%s)", GET_DO, GET_DONT))
	ans := 0
	do := true

	for i := 0; i < len(validSeqs); i++ {
		checkStr := conditional.FindString(validSeqs[i])
		if checkStr == `do()` {
			do = true
		} else if checkStr == `don't()` {
			do = false
		} else if do {
			nums := find.FindAllString(validSeqs[i], -1)
			ans += multiplyTwoStrings(nums[0], nums[1])
		}
	}

	return ans
}

func partOne(input string) int {

	find := regexp.MustCompile(GET_MUL)
	validSeqs := find.FindAllString(input, -1)

	find = regexp.MustCompile(`([0-9]+)`)
	ans := 0

	for i := 0; i < len(validSeqs); i++ {
		nums := find.FindAllString(validSeqs[i], -1)
		ans += multiplyTwoStrings(nums[0], nums[1])
	}

	return ans

}
