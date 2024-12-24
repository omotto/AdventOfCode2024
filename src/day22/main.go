package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2024/pkg/file"
)

func mix(a, b int) int {
	return a ^ b
}

func prune(a int) int {
	return a % 16777216
}

func multiply(a, b int) int {
	return prune(mix(a, a*b))
}

func divide(a, b int) int {
	return prune(mix(a, a/b))
}

func nextSecretNumber(secretNumber, times int) int {
	for i := 0; i < times; i++ {
		secretNumber = multiply(divide(multiply(secretNumber, 64), 32), 2048)
	}
	return secretNumber
}

func getSumSecretNumber(s []string) int {
	result := 0
	for _, line := range s {
		secretNumber, _ := strconv.Atoi(line)
		result += nextSecretNumber(secretNumber, 2000)
	}
	return result
}

func intSliceToString(values []int) string {
	var s []string
	for _, number := range values {
		s = append(s, strconv.Itoa(number))
	}
	return strings.Join(s, ",")
}

func getNumBananas(s []string) int {
	sequences := make(map[string]int)
	for _, line := range s {
		var sequence []int
		visited := make(map[string]struct{})
		secretNumber, _ := strconv.Atoi(line)
		prevSecretNumber, currSecretNumber := secretNumber, secretNumber
		for i := 0; i < 2000; i++ {
			sequence = append(sequence, currSecretNumber%10-prevSecretNumber%10)
			prevSecretNumber = currSecretNumber
			currSecretNumber = multiply(divide(multiply(prevSecretNumber, 64), 32), 2048)
			if len(sequence) == 4 {
				if _, ok := visited[intSliceToString(sequence)]; !ok {
					visited[intSliceToString(sequence)] = struct{}{}
					if v, ok := sequences[intSliceToString(sequence)]; !ok {
						sequences[intSliceToString(sequence)] = prevSecretNumber % 10
					} else {
						sequences[intSliceToString(sequence)] = v + prevSecretNumber%10
					}
				}
				sequence = sequence[1:]
			}
		}
	}
	highest := 0
	for _, v := range sequences {
		if v > highest {
			highest = v
		}
	}
	return highest
}

func main() {
	absPathName, _ := filepath.Abs("src/day22/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumSecretNumber(output))
	fmt.Println(getNumBananas(output))
}
