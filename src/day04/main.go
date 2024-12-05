package main

import (
	"fmt"
	"path/filepath"

	"advent2024/pkg/file"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
	UP_LEFT
	UP_RIGHT
	DOWN_LEFT
	DOWN_RIGHT
)

func find(s []string, x, y, d int, word string, charPos int) int {
	result := 0
	if s[x][y] == word[charPos] {
		if charPos == len(word)-1 {
			result = 1
		} else {
			switch d {
			case UP:
				if y > 0 {
					result = find(s, x, y-1, d, word, charPos+1)
				}
			case DOWN:
				if y < len(s)-1 {
					result = find(s, x, y+1, d, word, charPos+1)
				}
			case LEFT:
				if x > 0 {
					result = find(s, x-1, y, d, word, charPos+1)
				}
			case RIGHT:
				if x < len(s[y])-1 {
					result = find(s, x+1, y, d, word, charPos+1)
				}
			case UP_LEFT:
				if y > 0 && x > 0 {
					result = find(s, x-1, y-1, d, word, charPos+1)
				}
			case UP_RIGHT:
				if y > 0 && x < len(s[y])-1 {
					result = find(s, x+1, y-1, d, word, charPos+1)
				}
			case DOWN_LEFT:
				if y < len(s)-1 && x > 0 {
					result = find(s, x-1, y+1, d, word, charPos+1)
				}
			case DOWN_RIGHT:
				if y < len(s)-1 && x < len(s[y])-1 {
					result = find(s, x+1, y+1, d, word, charPos+1)
				}
			}
		}
	}
	return result
}

func findXMAS(s []string) int {
	result := 0
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			result += find(s, x, y, UP, "XMAS", 0)
			result += find(s, x, y, DOWN, "XMAS", 0)
			result += find(s, x, y, LEFT, "XMAS", 0)
			result += find(s, x, y, RIGHT, "XMAS", 0)
			result += find(s, x, y, UP_LEFT, "XMAS", 0)
			result += find(s, x, y, UP_RIGHT, "XMAS", 0)
			result += find(s, x, y, DOWN_LEFT, "XMAS", 0)
			result += find(s, x, y, DOWN_RIGHT, "XMAS", 0)
		}
	}
	return result
}

func findX_MAS(s []string) int {
	result := 0
	for y := 1; y < len(s)-1; y++ {
		for x := 1; x < len(s[y])-1; x++ {
			if string(s[y][x]) == "A" {
				if string(s[y-1][x-1]) == "M" && string(s[y-1][x+1]) == "M" && string(s[y+1][x-1]) == "S" && string(s[y+1][x+1]) == "S" {
					result++
				} else if string(s[y-1][x-1]) == "S" && string(s[y-1][x+1]) == "S" && string(s[y+1][x-1]) == "M" && string(s[y+1][x+1]) == "M" {
					result++
				} else if string(s[y-1][x-1]) == "S" && string(s[y-1][x+1]) == "M" && string(s[y+1][x-1]) == "S" && string(s[y+1][x+1]) == "M" {
					result++
				} else if string(s[y-1][x-1]) == "M" && string(s[y-1][x+1]) == "S" && string(s[y+1][x-1]) == "M" && string(s[y+1][x+1]) == "S" {
					result++
				}
			}
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day04/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(findXMAS(output))
	fmt.Println(findX_MAS(output))
}
