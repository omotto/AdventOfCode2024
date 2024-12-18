package main

import (
	"fmt"
	"path/filepath"

	"advent2024/pkg/file"
)

const (
	CoordKey = "%d:%d"
)

func getMinScoreMaze(room map[string]struct{}, sx, sy, ex, ey int) int {
	type Tile struct {
		x, y, score int
	}
	visited := map[string]int{
		fmt.Sprintf(CoordKey, sx, sy): 0,
	}
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
			return tile.score
		}
		newScore := tile.score + 1
		if _, ok := room[fmt.Sprintf(CoordKey, tile.x-1, tile.y)]; !ok && tile.x > 0 {
			if _, ok := visited[fmt.Sprintf(CoordKey, tile.x-1, tile.y)]; !ok {
				visited[fmt.Sprintf(CoordKey, tile.x-1, tile.y)] = newScore
				queue = append(queue, Tile{
					x:     tile.x - 1,
					y:     tile.y,
					score: newScore,
				})
			}
		}
		if _, ok := room[fmt.Sprintf(CoordKey, tile.x+1, tile.y)]; !ok && tile.x < ex {
			if _, ok := visited[fmt.Sprintf(CoordKey, tile.x+1, tile.y)]; !ok {
				visited[fmt.Sprintf(CoordKey, tile.x+1, tile.y)] = newScore
				queue = append(queue, Tile{
					x:     tile.x + 1,
					y:     tile.y,
					score: newScore,
				})
			}
		}
		if _, ok := room[fmt.Sprintf(CoordKey, tile.x, tile.y-1)]; !ok && tile.y > 0 {
			if _, ok := visited[fmt.Sprintf(CoordKey, tile.x, tile.y-1)]; !ok {
				visited[fmt.Sprintf(CoordKey, tile.x, tile.y-1)] = newScore
				queue = append(queue, Tile{
					x:     tile.x,
					y:     tile.y - 1,
					score: newScore,
				})
			}
		}
		if _, ok := room[fmt.Sprintf(CoordKey, tile.x, tile.y+1)]; !ok && tile.y < ey {
			if _, ok := visited[fmt.Sprintf(CoordKey, tile.x, tile.y+1)]; !ok {
				visited[fmt.Sprintf(CoordKey, tile.x, tile.y+1)] = newScore
				queue = append(queue, Tile{
					x:     tile.x,
					y:     tile.y + 1,
					score: newScore,
				})
			}
		}
	}
	return -1
}

func parseInput(s []string, bytes int) map[string]struct{} {
	var x, y int
	room := make(map[string]struct{})
	for idx, line := range s {
		if idx == bytes {
			break
		}
		_, _ = fmt.Sscanf(line, "%d,%d", &x, &y)
		room[fmt.Sprintf(CoordKey, x, y)] = struct{}{}
	}
	return room
}

func getMinPath(s []string, bytes, ex, ey int) int {
	room := parseInput(s, bytes)
	score := getMinScoreMaze(room, 0, 0, ex, ey)
	return score
}

func getCoord(s []string, bytes, ex, ey int) string {
	for idx := bytes; idx < len(s); idx++ {
		room := parseInput(s, idx)
		score := getMinScoreMaze(room, 0, 0, ex, ey)
		if score == -1 {
			return s[idx-1]
		}
	}
	return "not found"
}

func main() {
	absPathName, _ := filepath.Abs("src/day18/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getMinPath(output, 1024, 70, 70))
	fmt.Println(getCoord(output, 1024, 70, 70))
}
