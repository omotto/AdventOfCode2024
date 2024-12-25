package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"advent2024/pkg/file"
)

func parseInput(s []string) ([][]int, [][]int) {
	var (
		locks, keys     [][]int
		idxLock, idxKey int
		isLock          bool
	)
	first := true
	for _, line := range s {
		if first {
			if strings.Count(line, "#") == len(line) {
				lock := make([]int, len(line))
				locks = append(locks, lock)
				idxLock++
				isLock = true
			} else {
				key := make([]int, len(line))
				keys = append(keys, key)
				idxKey++
				isLock = false
			}
			first = false
			continue
		}
		if len(line) < 2 {
			first = true
			if !isLock {
				for idx := 0; idx < len(keys[idxKey-1]); idx++ {
					keys[idxKey-1][idx]--
				}
			}
			continue
		}

		for idx, char := range line {
			if char == '#' {
				if isLock {
					locks[idxLock-1][idx]++
				} else {
					keys[idxKey-1][idx]++
				}
			}
		}
	}
	return locks, keys
}

func getNumValidKeys(s []string) int {
	locks, keys := parseInput(s)
	result := 0
	var fits bool
	for _, key := range keys {
		for _, lock := range locks {
			fits = true
			for idx := 0; idx < len(lock); idx++ {
				if key[idx]+lock[idx] > 5 {
					fits = false
					break
				}
			}
			if fits {
				result++
			}
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day25/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumValidKeys(output))
}
