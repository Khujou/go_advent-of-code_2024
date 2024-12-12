package main

import (
	"bufio"
	"fmt"
	"os"
)

const INPUT_FILE = "input.txt"
const XMAS = `XMAS`

func check(err error) {
	if err != nil {
		panic(err)
	}

}

func makeInput2DArray(f *os.File) [][]rune {
	input := make([][]rune, 0, 0)

	scanner := bufio.NewScanner(f)
	for row := 0; scanner.Scan(); row++ {
		line := []rune(scanner.Text())
		input = append(input, make([]rune, len(line)))
		for col := 0; col < len(line); col++ {
			input[row][col] = line[col]
		}
	}

	return input
}

func isOutOfBounds(pos []int, size []int) bool {
	if pos[0] < 0 || pos[0] >= size[0] {
		return true
	} else if pos[1] < 0 || pos[1] >= size[1] {
		return true
	}
	return false
}

func countXMAS(input [][]rune, pos []int, ans *int) {
	strs := []string{}
	for i := -1; i <= 1; i++ {
		col := 0
		for j := -1; j <= 1; j++ {
			row := 0
			str := []rune{}
			for k := 0; k < 4; k++ {
				row = k*j + pos[0]
				col = k*i + pos[1]
				if isOutOfBounds([]int{row, col}, []int{len(input), len(input[0])}) {
					continue
				}
				str = append(str, input[row][col])
			}
			strstr := string(str)
			if strstr == "XMAS" {
				//fmt.Printf("XMAS @ %d, going to %d, %d\n", pos, row, col)
				strs = append(strs, strstr)
			}
		}
	}
	*ans += len(strs)
}

func countX_MAS(input [][]rune, pos []int, ans *int) {
	strs := []string{}
	for i := -1; i <= 1; i += 2 {
		col := 0
		for j := -1; j <= 1; j += 2 {
			row := 0
			str := []rune{}
			for k := 1; k >= -1; k-- {
				row = k*j + pos[0]
				col = k*i + pos[1]
				if isOutOfBounds([]int{row, col}, []int{len(input), len(input[0])}) {
					continue
				}
				str = append(str, input[row][col])
			}
			strstr := string(str)
			if strstr == "MAS" {
				//fmt.Printf("X_MAS @ %d, going to %d, %d\n", pos, row, col)
				strs = append(strs, strstr)
			}
		}
	}
	if len(strs) == 2 {
		*ans++
	}
}

func main() {
	f, err := os.Open(INPUT_FILE)
	check(err)
	input := makeInput2DArray(f)
	partOneAns := bothParts(input, 0, 'X', countXMAS)
	fmt.Printf("Part One - %d\n", partOneAns)
	partTwoAns := bothParts(input, 1, 'A', countX_MAS)
	fmt.Printf("Part Two - %d", partTwoAns)
}

func bothParts(input [][]rune, offset int, findRune rune, callback func([][]rune, []int, *int)) int {
	ans := 0

	for row := 0 + offset; row < len(input)-offset; row++ {
		line := input[row]
		for col := 0 + offset; col < len(line)-offset; col++ {
			if line[col] == findRune {
				callback(input, []int{row, col}, &ans)
			}
		}
	}

	return ans
}
