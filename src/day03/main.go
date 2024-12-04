package main

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"advent2024/pkg/file"
)

func getMult(mults []string) int {
	result := 0
	for _, mult := range mults {
		parts := strings.Split(mult, ",")
		a, _ := strconv.Atoi(parts[0][4:])
		b, _ := strconv.Atoi(parts[1][:len(parts[1])-1])
		result += a * b
	}
	return result
}

func getSumMults(s []string) int {
	result := 0
	regex, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	for _, line := range s {
		mults := regex.FindAllString(line, -1)
		result += getMult(mults)
	}
	return result
}

func getMult2(mults []string, multIdx, doIdx, dontIdx [][]int) int {
	result := 0
	for i, mult := range mults {
		// Ger mul position
		strIdx := multIdx[i][0]
		// Get nearest left side don't position
		dontPos := -1
		for _, dontStrIdx := range dontIdx {
			if dontStrIdx[0] < strIdx {
				dontPos = dontStrIdx[0]
			} else {
				break
			}
		}
		// Get nearest do left side do position
		doPos := -1
		for _, doStrIdx := range doIdx {
			if doStrIdx[0] < strIdx {
				doPos = doStrIdx[0]
			} else {
				break
			}
		}
		// Check if mul can be applied or not
		if dontPos > -1 && dontPos > doPos {
			continue
		}
		parts := strings.Split(mult, ",")
		a, _ := strconv.Atoi(parts[0][4:])
		b, _ := strconv.Atoi(parts[1][:len(parts[1])-1])
		result += a * b
	}
	return result
}

func getSumMults2(s []string) int {
	var buffer bytes.Buffer
	multRegex, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	doRegex, _ := regexp.Compile("do\\(\\)")
	dontRegex, _ := regexp.Compile("don\\'t\\(\\)")
	// Put all in one single line
	for _, line := range s {
		_, _ = buffer.WriteString(line)
	}
	line := buffer.String()
	multIdx := multRegex.FindAllStringIndex(line, -1)
	mults := multRegex.FindAllString(line, -1)
	doIdx := doRegex.FindAllStringIndex(line, -1)
	dontIdx := dontRegex.FindAllStringIndex(line, -1)
	return getMult2(mults, multIdx, doIdx, dontIdx)
}

func main() {
	absPathName, _ := filepath.Abs("src/day03/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumMults(output))
	fmt.Println(getSumMults2(output))
}
