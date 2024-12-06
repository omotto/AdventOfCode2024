package main

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"advent2024/pkg/file"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

func getStart(s []string) (int, int, int) {
	for y, line := range s {
		for x := 0; x < len(line); x++ {
			if s[y][x] == '^' {
				return x, y, UP
			} else if s[y][x] == 'v' {
				return x, y, DOWN
			} else if s[y][x] == '>' {
				return x, y, RIGHT
			} else if s[y][x] == '<' {
				return x, y, LEFT
			}
		}
	}
	return -1, -1, -1
}

func getGuardPositions(s []string) (map[string]struct{}, bool) {
	x, y, direction := getStart(s)
	path := make(map[string]struct{})
	path[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
	overlapped := make(map[string]struct{})
	overlapped[fmt.Sprintf("%d:%d:%d", x, y, direction)] = struct{}{}
	exit := false
	for !exit {
		switch direction {
		case UP:
			if y > 0 {
				if s[y-1][x] != '#' {
					y--
					if _, ok := overlapped[fmt.Sprintf("%d:%d:%d", x, y, direction)]; ok {
						return nil, true
					}
					overlapped[fmt.Sprintf("%d:%d:%d", x, y, direction)] = struct{}{}
					if _, ok := path[fmt.Sprintf("%d:%d", x, y)]; !ok {
						path[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					}
				} else {
					direction = RIGHT
				}
			} else {
				exit = true
			}
		case DOWN:
			if y < len(s)-1 {
				if s[y+1][x] != '#' {
					y++
					if _, ok := overlapped[fmt.Sprintf("%d:%d:%d", x, y, direction)]; ok {
						return nil, true
					}
					overlapped[fmt.Sprintf("%d:%d:%d", x, y, direction)] = struct{}{}
					if _, ok := path[fmt.Sprintf("%d:%d", x, y)]; !ok {
						path[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					}
				} else {
					direction = LEFT
				}
			} else {
				exit = true
			}
		case RIGHT:
			if x < len(s[y])-1 {
				if s[y][x+1] != '#' {
					x++
					if _, ok := overlapped[fmt.Sprintf("%d:%d:%d", x, y, direction)]; ok {
						return nil, true
					}
					overlapped[fmt.Sprintf("%d:%d:%d", x, y, direction)] = struct{}{}
					if _, ok := path[fmt.Sprintf("%d:%d", x, y)]; !ok {
						path[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					}
				} else {
					direction = DOWN
				}
			} else {
				exit = true
			}
		case LEFT:
			if x > 0 {
				if s[y][x-1] != '#' {
					x--
					if _, ok := overlapped[fmt.Sprintf("%d:%d:%d", x, y, direction)]; ok {
						return nil, true
					}
					overlapped[fmt.Sprintf("%d:%d:%d", x, y, direction)] = struct{}{}
					if _, ok := path[fmt.Sprintf("%d:%d", x, y)]; !ok {
						path[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					}
				} else {
					direction = UP
				}
			} else {
				exit = true
			}
		}
	}
	return path, false
}

func getSumGuardStuckPositions(s []string, path map[string]struct{}) int {
	result := 0
	for k, _ := range path {
		parts := strings.Split(k, ":")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		if s[y][x] == '.' {
			ss := slices.Clone(s)
			ss[y] = ss[y][:x] + string('#') + ss[y][x+1:]
			_, loop := getGuardPositions(ss)
			if loop {
				result++
			}
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day06/input.txt")
	output, _ := file.ReadInput(absPathName)

	path, _ := getGuardPositions(output)
	fmt.Println(len(path))
	fmt.Println(getSumGuardStuckPositions(output, path))
}
