package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"advent2024/pkg/file"
)

func decompressCode(s string) ([]int, int, int, [][3]int) {
	table := make([][3]int, len(s)/2+1)
	filledLength := 0
	emptyLength := 0
	for i := 0; i < len(s); i += 2 {
		table[i/2][0], _ = strconv.Atoi(string(s[i]))
		if i == len(s)-1 {
			table[i/2][1] = 0
		} else {
			table[i/2][1], _ = strconv.Atoi(string(s[i+1]))
		}
		table[i/2][2] = filledLength + emptyLength
		filledLength += table[i/2][0]
		emptyLength += table[i/2][1]
	}
	var result []int
	for i := 0; i < len(table); i++ {
		for j := 0; j < table[i][0]; j++ {
			result = append(result, i)
		}
		for j := 0; j < table[i][1]; j++ {
			result = append(result, -1)
		}
	}
	return result, filledLength, emptyLength, table
}

func getFilledCode(code []int, filledLength int) []int {
	endIdx := len(code) - 1
	beginIdx := 0
	var finalCode []int
	for {
		if code[beginIdx] != -1 {
			finalCode = append(finalCode, code[beginIdx])
		} else {
			for i := endIdx; i > 0; i-- {
				if code[i] != -1 {
					endIdx = i
					break
				}
			}
			finalCode = append(finalCode, code[endIdx])
			endIdx--
		}
		beginIdx++
		if beginIdx == filledLength {
			break
		}
	}
	return finalCode
}

func getChecksum(s []string) int {
	code, filled, _, _ := decompressCode(s[0])
	resultCode := getFilledCode(code, filled)
	result := 0
	for idx, value := range resultCode {
		result += idx * value
	}
	return result
}

func remove(code []int, pos, times int) []int {
	for j := 0; j < times; j++ {
		code[pos+j] = -1
	}
	return code
}

func add(code []int, pos, valueAdd, times int) []int {
	for j := 0; j < times; j++ {
		code[pos+j] = valueAdd
	}
	return code
}

func getFilledCodeBlocks(code []int, table [][3]int) []int {
	for i := len(table) - 1; i >= 0; i-- {
		timesToFill := table[i][0]
		lastFoundPos := table[i][2]
		k := 0
		for j := 0; j < lastFoundPos; j++ {
			if code[j] == -1 {
				k++
			} else {
				k = 0
			}
			if k == timesToFill {
				code = remove(code, lastFoundPos, timesToFill)
				code = add(code, j-timesToFill+1, i, timesToFill)
				break
			}
		}
	}
	return code
}

func getChecksumBlocks(s []string) int {
	code, _, _, table := decompressCode(s[0])
	resultCode := getFilledCodeBlocks(code, table)
	result := 0
	for idx, value := range resultCode {
		if value != -1 {
			result += idx * value
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day09/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getChecksum(output))
	fmt.Println(getChecksumBlocks(output))
}
