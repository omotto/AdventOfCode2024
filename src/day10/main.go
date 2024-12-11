package main

import (
	"fmt"
	"path/filepath"
	"sync"

	"advent2024/pkg/file"
)

func getNumTrailheads(s []string, x, y int, peaks map[string]struct{}, all bool) int {
	result := 0
	if s[y][x] == '9' {
		if !all {
			if _, ok := peaks[fmt.Sprintf("%d:%d", x, y)]; ok {
				return 0
			}
			peaks[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
		}
		return 1
	}
	if y > 0 && s[y-1][x] == s[y][x]+1 {
		result += getNumTrailheads(s, x, y-1, peaks, all)
	}
	if y < len(s)-1 && s[y+1][x] == s[y][x]+1 {
		result += getNumTrailheads(s, x, y+1, peaks, all)
	}
	if x > 0 && s[y][x-1] == s[y][x]+1 {
		result += getNumTrailheads(s, x-1, y, peaks, all)
	}
	if x < len(s[y])-1 && s[y][x+1] == s[y][x]+1 {
		result += getNumTrailheads(s, x+1, y, peaks, all)
	}
	return result
}

func getSumTrailheads(s []string, all bool) int {
	result := 0
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == '0' {
				peaks := make(map[string]struct{})
				result += getNumTrailheads(s, x, y, peaks, all)
			}
		}
	}
	return result
}

func getSumTrailheads2(s []string, all bool) int {
	var bases [][2]int
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == '0' {
				bases = append(bases, [2]int{x, y})
			}
		}
	}
	var wg sync.WaitGroup
	ch := make(chan int, len(bases))
	for _, base := range bases {
		wg.Add(1)
		go func(x, y int) {
			defer wg.Done()
			peaks := make(map[string]struct{})
			ch <- getNumTrailheads(s, x, y, peaks, all)
		}(base[0], base[1])
	}
	wg.Wait()
	close(ch)
	result := 0
	for value := range ch {
		result += value
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day10/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumTrailheads(output, false))
	fmt.Println(getSumTrailheads(output, true))

	fmt.Println(getSumTrailheads2(output, false))
	fmt.Println(getSumTrailheads2(output, true))
}
