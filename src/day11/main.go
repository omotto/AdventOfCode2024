package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2024/pkg/file"
)

func getNumStones(s []string, times int) int {
	strStones := strings.Split(s[0], " ")
	stones := make([]int, len(strStones))
	for i, strStone := range strStones {
		stones[i], _ = strconv.Atoi(strStone)
	}
	for i := 0; i < times; i++ {
		var newStones []int
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if digits := strconv.Itoa(stone); len(digits)%2 == 0 {
				a, _ := strconv.Atoi(digits[:len(digits)/2])
				b, _ := strconv.Atoi(digits[len(digits)/2:])
				newStones = append(newStones, a)
				newStones = append(newStones, b)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	return len(stones)
}

func getNumStones2(s []string, times int) int {
	strStones := strings.Split(s[0], " ")
	stoneMap := make(map[int]int)
	for _, strStone := range strStones {
		value, _ := strconv.Atoi(strStone)
		stoneMap[value] = stoneMap[value] + 1
	}
	for i := 0; i < times; i++ {
		newStoneMap := make(map[int]int)
		for stone, count := range stoneMap {
			var newStones []int
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if digits := strconv.Itoa(stone); len(digits)%2 == 0 {
				a, _ := strconv.Atoi(digits[:len(digits)/2])
				b, _ := strconv.Atoi(digits[len(digits)/2:])
				newStones = append(newStones, a)
				newStones = append(newStones, b)
			} else {
				newStones = append(newStones, stone*2024)
			}
			for _, newStone := range newStones {
				newStoneMap[newStone] = newStoneMap[newStone] + count // by the number of times that stone appears
			}
		}
		stoneMap = newStoneMap
	}
	result := 0
	for _, count := range stoneMap {
		result += count
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day11/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumStones(output, 25))
	fmt.Println(getNumStones2(output, 75))
}
