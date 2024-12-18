package main

import (
	"fmt"
	"path/filepath"
	"slices"

	"advent2024/pkg/file"
)

const (
	CoordKey    = "%d:%d"
	CoordDirKey = "%d:%d:%d"

	UP    = 1
	DOWN  = 2
	RIGHT = 3
	LEFT  = 4
)

func getMinScoreMaze(room map[string]struct{}, sx, sy, ex, ey int) map[string]int {
	type Tile struct {
		x, y, dir, score int
	}
	visited := map[string]int{
		fmt.Sprintf("%d:%d:%d", sx, sy, RIGHT): 0,
	}
	queue := make([]Tile, 0)
	queue = append(queue, Tile{
		x:     sx,
		y:     sy,
		dir:   RIGHT,
		score: 0,
	})
	for len(queue) > 0 {
		tile := queue[0]  // Get first
		queue = queue[1:] // Remove it
		if v, ok := visited[fmt.Sprintf(CoordDirKey, tile.x, tile.y, tile.dir)]; ok && v < tile.score {
			continue
		}
		newScore := tile.score + 1
		if _, ok := room[fmt.Sprintf(CoordKey, tile.x-1, tile.y)]; !ok && tile.x > 0 {
			if v, ok := visited[fmt.Sprintf(CoordDirKey, tile.x-1, tile.y, LEFT)]; (ok && newScore < v) || !ok {
				visited[fmt.Sprintf(CoordDirKey, tile.x-1, tile.y, LEFT)] = newScore
				queue = append(queue, Tile{
					x:     tile.x - 1,
					y:     tile.y,
					dir:   LEFT,
					score: newScore,
				})
			}
		}
		if _, ok := room[fmt.Sprintf(CoordKey, tile.x+1, tile.y)]; !ok && tile.x < ex {
			if v, ok := visited[fmt.Sprintf(CoordDirKey, tile.x+1, tile.y, RIGHT)]; (ok && newScore < v) || !ok {
				visited[fmt.Sprintf(CoordDirKey, tile.x+1, tile.y, RIGHT)] = newScore
				queue = append(queue, Tile{
					x:     tile.x + 1,
					y:     tile.y,
					dir:   RIGHT,
					score: newScore,
				})
			}
		}
		if _, ok := room[fmt.Sprintf(CoordKey, tile.x, tile.y-1)]; !ok && tile.y > 0 {
			if v, ok := visited[fmt.Sprintf(CoordDirKey, tile.x, tile.y-1, UP)]; (ok && newScore < v) || !ok {
				visited[fmt.Sprintf(CoordDirKey, tile.x, tile.y-1, UP)] = newScore
				queue = append(queue, Tile{
					x:     tile.x,
					y:     tile.y - 1,
					dir:   UP,
					score: newScore,
				})
			}
		}
		if _, ok := room[fmt.Sprintf(CoordKey, tile.x, tile.y+1)]; !ok && tile.y < ey {
			if v, ok := visited[fmt.Sprintf(CoordDirKey, tile.x, tile.y+1, DOWN)]; (ok && newScore < v) || !ok {
				visited[fmt.Sprintf(CoordDirKey, tile.x, tile.y+1, DOWN)] = newScore
				queue = append(queue, Tile{
					x:     tile.x,
					y:     tile.y + 1,
					dir:   DOWN,
					score: newScore,
				})
			}
		}
	}
	return visited
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
	minScores := getMinScoreMaze(room, 0, 0, ex, ey)
	var scores []int
	if v, ok := minScores[fmt.Sprintf(CoordDirKey, ex, ey, UP)]; ok {
		scores = append(scores, v)
	}
	if v, ok := minScores[fmt.Sprintf(CoordDirKey, ex, ey, DOWN)]; ok {
		scores = append(scores, v)
	}
	if v, ok := minScores[fmt.Sprintf(CoordDirKey, ex, ey, LEFT)]; ok {
		scores = append(scores, v)
	}
	if v, ok := minScores[fmt.Sprintf(CoordDirKey, ex, ey, RIGHT)]; ok {
		scores = append(scores, v)
	} /*
		for k, v := range minScores {
			fmt.Printf("%s = %d\r\n", k, v)
		}*/
	return slices.Min(scores)
}

func getCoord(s []string, bytes, ex, ey int) string {
	for idx := bytes; idx < len(s); idx++ {
		room := parseInput(s, idx)
		minScores := getMinScoreMaze(room, 0, 0, ex, ey)
		var scores []int
		if v, ok := minScores[fmt.Sprintf(CoordDirKey, ex, ey, UP)]; ok {
			scores = append(scores, v)
		}
		if v, ok := minScores[fmt.Sprintf(CoordDirKey, ex, ey, DOWN)]; ok {
			scores = append(scores, v)
		}
		if v, ok := minScores[fmt.Sprintf(CoordDirKey, ex, ey, LEFT)]; ok {
			scores = append(scores, v)
		}
		if v, ok := minScores[fmt.Sprintf(CoordDirKey, ex, ey, RIGHT)]; ok {
			scores = append(scores, v)
		}
		if len(scores) == 0 {
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
