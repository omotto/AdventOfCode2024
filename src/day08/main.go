package main

import (
	"fmt"
	"path/filepath"

	"advent2024/pkg/file"
)

func checkAntinodes(sx, sy int, s []string, locations map[string]struct{}, allowMultiple bool) int {
	result := 0
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if !(sx == x && sy == y) && s[y][x] == s[sy][sx] {
				if allowMultiple {
					c := 0
					for {
						dy := (sy - y) * c
						dx := (sx - x) * c
						if sy+dy >= len(s) || sy+dy < 0 || sx+dx < 0 || sx+dx >= len(s[0]) {
							break
						}
						if s[sy+dy][sx+dx] == '.' || s[sy+dy][sx+dx] == s[sy][sx] {
							if _, ok := locations[fmt.Sprintf("%d:%d", sx+dx, sy+dy)]; !ok {
								locations[fmt.Sprintf("%d:%d", sx+dx, sy+dy)] = struct{}{}
								result++
							}
						}
						c++
					}
					c = 0
					for {
						dy := (y - sy) * c
						dx := (x - sx) * c
						if y+dy >= len(s) || y+dy < 0 || x+dx < 0 || x+dx >= len(s[0]) {
							break
						}
						if s[y+dy][x+dx] == '.' || s[sy+dy][sx+dx] == s[sy][sx] {
							if _, ok := locations[fmt.Sprintf("%d:%d", x+dx, y+dy)]; !ok {
								locations[fmt.Sprintf("%d:%d", x+dx, y+dy)] = struct{}{}
								result++
							}
						}
						c++
					}
				} else {
					dy := sy - y
					dx := sx - x
					if sy+dy < len(s) && sy+dy >= 0 && sx+dx >= 0 && sx+dx < len(s[0]) && (s[sy+dy][sx+dx] == '.' || s[sy+dy][sx+dx] == s[sy][sx]) {
						if _, ok := locations[fmt.Sprintf("%d:%d", sx+dx, sy+dy)]; !ok {
							locations[fmt.Sprintf("%d:%d", sx+dx, sy+dy)] = struct{}{}
							result++
						}
					}
					dy = y - sy
					dx = x - sx
					if y+dy < len(s) && y+dy >= 0 && x+dx >= 0 && x+dx < len(s[0]) && (s[y+dy][x+dx] == '.' || s[sy+dy][sx+dx] == s[sy][sx]) {
						if _, ok := locations[fmt.Sprintf("%d:%d", x+dx, y+dy)]; !ok {
							locations[fmt.Sprintf("%d:%d", x+dx, y+dy)] = struct{}{}
							result++
						}
					}
				}
			}
		}
	}
	return result
}

func getSumSignalLocations(s []string, allowMultiple bool) int {
	result := 0
	locations := map[string]struct{}{}
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] != '.' {
				result += checkAntinodes(x, y, s, locations, allowMultiple)
			}
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day08/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumSignalLocations(output, false))
	fmt.Println(getSumSignalLocations(output, true))
}
