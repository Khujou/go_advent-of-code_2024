package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

const INPUT_FILE = "input.txt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Make 2d array of each page order rule, Make string array of each page update
func parseInput(f *os.File) ([][]int, []string) {
	pageOrderRules := [][]int{}
	pageUpdates := []string{}
	regexStr := regexp.MustCompile(`([0-9]+)\|([0-9]+)`)

	scanner := bufio.NewScanner(f)
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		if regexStr.MatchString(line) {
			pageOrderRule := []int{}
			for _, str := range strings.Split(line, "|") {
				num, err := strconv.Atoi(str)
				check(err)
				pageOrderRule = append(pageOrderRule, num)
			}
			pageOrderRules = append(pageOrderRules, pageOrderRule)
		} else if line == "" {
			continue
		} else {
			pageUpdates = append(pageUpdates, line)
		}
	}

	return pageOrderRules, pageUpdates
}

func makeRegExps(pageOrderRules [][]int) []*regexp.Regexp {
	regexStrs := []*regexp.Regexp{}
	for _, pageOrderRule := range pageOrderRules {
		regexStr := regexp.MustCompile(fmt.Sprintf("%d(.+)%d", pageOrderRule[0], pageOrderRule[1]))
		regexStrs = append(regexStrs, regexStr)
	}

	//fmt.Println(regexStrs)

	return regexStrs
}

func addIfUpdateIsGood(pageOrderingRegExps []*regexp.Regexp, pageUpdate string) (bool, int) {
	pageUpdateArray := strings.Split(pageUpdate, ",")
	for _, pageOrderRuleRegExp := range pageOrderingRegExps {
		//fmt.Println(pageOrderRuleRegExp)
		pages := strings.Split(pageOrderRuleRegExp.String(), "(.+)")
		if !(slices.Contains(pageUpdateArray, pages[0]) && slices.Contains(pageUpdateArray, pages[1])) {
			continue
		}
		if !(pageOrderRuleRegExp.MatchString(pageUpdate)) {
			return false, 0
		}
	}
	//fmt.Printf("All Good!\n\n")
	middleNum, err := strconv.Atoi(pageUpdateArray[len(pageUpdateArray)/2])
	check(err)
	//fmt.Println(middleNum)
	return true, middleNum
}

func addAfterFixingBadUpdate(pageOrderingRegExps []*regexp.Regexp, pageUpdate string) (bool, int) {
	updateIsGood, _ := addIfUpdateIsGood(pageOrderingRegExps, pageUpdate)
	if updateIsGood {
		return false, 0
	}

	pageUpdateArray := strings.Split(pageUpdate, ",")
	for i := 0; i < len(pageOrderingRegExps); i++ {
		pageOrderRuleRegExp := pageOrderingRegExps[i]
		pages := strings.Split(pageOrderRuleRegExp.String(), "(.+)")
		if !(slices.Contains(pageUpdateArray, pages[0]) && slices.Contains(pageUpdateArray, pages[1])) {
			continue
		}
		if !(pageOrderRuleRegExp.MatchString(pageUpdate)) {
			//swap
			tmpIdx := slices.Index(pageUpdateArray, pages[0])
			tmp2Idx := slices.Index(pageUpdateArray, pages[1])
			tmp := pageUpdateArray[tmpIdx]
			pageUpdateArray[tmpIdx] = pageUpdateArray[tmp2Idx]
			pageUpdateArray[tmp2Idx] = tmp
			pageUpdate = strings.Join(pageUpdateArray, ",")

			updateIsGood, _ := addIfUpdateIsGood(pageOrderingRegExps, pageUpdate)
			if updateIsGood {
				break
			}
			i = 0
		}
	}

	middleNum, err := strconv.Atoi(pageUpdateArray[len(pageUpdateArray)/2])
	check(err)
	return true, middleNum
}

func main() {
	f, err := os.Open(INPUT_FILE)
	check(err)
	pageOrderRules, pageUpdates := parseInput(f)
	pageOrderingRegExps := makeRegExps(pageOrderRules)

	beforePartOne := time.Now()
	partOneAns := bothParts(pageOrderingRegExps, pageUpdates, addIfUpdateIsGood)
	fmt.Printf("Part One - %d\nCompleted in %d milliseconds\n", partOneAns, time.Since(beforePartOne).Milliseconds())

	beforePartTwo := time.Now()
	partTwoAns := bothParts(pageOrderingRegExps, pageUpdates, addAfterFixingBadUpdate)

	fmt.Printf("Part Two - %d\nCompleted in %d milliseconds", partTwoAns, time.Since(beforePartTwo).Milliseconds())

}

func bothParts(pageOrderingRegExps []*regexp.Regexp, pageUpdates []string, callback func([]*regexp.Regexp, string) (bool, int)) int {
	ans := 0

	for _, pageUpdate := range pageUpdates {
		//fmt.Printf("Page Update -> %s\n", pageUpdate)
		good, res := callback(pageOrderingRegExps, pageUpdate)
		if good {
			ans += res
		}
	}

	return ans
}
