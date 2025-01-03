package main

import (
	"fmt"
	"math"
	"path/filepath"

	"advent2024/pkg/file"
)

type coords struct {
	x, y int
}

const (
	V = 1
	H = 0
)

func parseInput(s []string) (int, int, int, int) {
	var sx, sy, ex, ey int
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == 'S' {
				sx, sy = x, y
			} else if s[y][x] == 'E' {
				ex, ey = x, y
			}
		}
	}
	return sx, sy, ex, ey
}

func getScoreMaze(room []string, sx, sy, ex, ey int) map[coords]int {
	type Tile struct {
		x, y, score int
	}
	directions := [4][2]int{
		{-1, 0}, {+1, 0}, {0, -1}, {0, +1},
	}
	visited := map[coords]int{coords{sx, sy}: 0}
	queue := make([]Tile, 0)
	queue = append(queue, Tile{
		x:     sx,
		y:     sy,
		score: 0,
	})
	for len(queue) > 0 {
		tile := queue[0]  // Get first
		queue = queue[1:] // Remove it
		if tile.x == ex && tile.y == ey {
			return visited
		}
		newScore := tile.score + 1
		for _, direction := range directions {
			newX := tile.x + direction[0]
			newY := tile.y + direction[1]
			if _, ok := visited[coords{newX, newY}]; !ok && room[newY][newX] != '#' {
				visited[coords{newX, newY}] = newScore
				queue = append(queue, Tile{
					x:     newX,
					y:     newY,
					score: newScore,
				})
			}
		}
	}
	return nil
}

func getCheats(s []string) [][3]int {
	var cheats [][3]int
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == '#' {
				if y > 0 && y < len(s)-1 && s[y-1][x] != '#' && s[y+1][x] != '#' {
					cheats = append(cheats, [3]int{x, y, V})
				} else if x > 0 && x < len(s[y])-1 && s[y][x-1] != '#' && s[y][x+1] != '#' {
					cheats = append(cheats, [3]int{x, y, H})
				}
			}
		}
	}
	return cheats
}

func getNumCheats(s []string, edge int) int {
	sx, sy, ex, ey := parseInput(s)
	scoreMap := getScoreMaze(s, sx, sy, ex, ey)
	cheats := getCheats(s)
	result := 0
	for _, cheat := range cheats {
		var v1, v2 int
		x, y, d := cheat[0], cheat[1], cheat[2]
		if d == H {
			v1 = scoreMap[coords{x - 1, y}]
			v2 = scoreMap[coords{x + 1, y}]
		} else {
			v1 = scoreMap[coords{x, y - 1}]
			v2 = scoreMap[coords{x, y + 1}]
		}
		if int(math.Abs(float64(v2-v1))) > edge {
			result++
		}
	}
	return result
}

func getDistance(sx, sy, ex, ey int) int {
	return int(math.Abs(float64(sx-ex)) + math.Abs(float64(sy-ey)))
}

func getCheats2(scoreMap map[coords]int) [][2]coords {
	var newDestinations [][2]coords
	for sxy, _ := range scoreMap {
		for dx := -20; dx < 21; dx++ { // From LEFT to RIGHT
			dxMax := 20 - int(math.Abs(float64(dx)))
			for dy := -dxMax; dy < dxMax+1; dy++ { // From UP to Down
				if _, ok := scoreMap[coords{sxy.x + dx, sxy.y + dy}]; ok {
					newDestinations = append(newDestinations, [2]coords{{sxy.x, sxy.y}, {sxy.x + dx, sxy.y + dy}})
				}
			}
		}
	}
	return newDestinations
}

func getNumCheats2(s []string, edge int) int {
	sx, sy, ex, ey := parseInput(s)
	scoreMap := getScoreMaze(s, sx, sy, ex, ey)
	cheats := getCheats2(scoreMap)
	result := 0
	for _, cheat := range cheats {
		if scoreMap[coords{cheat[1].x, cheat[1].y}]-scoreMap[coords{cheat[0].x, cheat[0].y}]-getDistance(cheat[0].x, cheat[0].y, cheat[1].x, cheat[1].y) >= edge {
			result++
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day20/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getNumCheats(output, 100))
	fmt.Println(getNumCheats2(output, 100))
}
