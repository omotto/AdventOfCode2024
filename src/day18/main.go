package main

import (
	"fmt"
	"path/filepath"

	"advent2024/pkg/file"
)

const Key = "%d:%d"

func getMinScoreMaze(room map[string]struct{}, sx, sy, ex, ey int) int {
	type Tile struct {
		x, y, score int
	}
	directions := [4][2]int{
		{-1, 0}, {+1, 0}, {0, -1}, {0, +1},
	}
	visited := map[string]int{
		fmt.Sprintf(Key, sx, sy): 0,
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
		for _, direction := range directions {
			newX := tile.x + direction[0]
			newY := tile.y + direction[1]
			if _, ok := room[fmt.Sprintf(Key, newX, newY)]; !ok && newX >= 0 && newY >= 0 && newX <= ex && newY <= ey {
				if _, ok := visited[fmt.Sprintf(Key, newX, newY)]; !ok {
					visited[fmt.Sprintf(Key, newX, newY)] = newScore
					queue = append(queue, Tile{
						x:     newX,
						y:     newY,
						score: newScore,
					})
				}
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
		room[fmt.Sprintf(Key, x, y)] = struct{}{}
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

func getCoordRecursive(s []string, start, end, ex, ey int) string {
	result := "not found"
	if end == start+1 {
		result = s[end-1]
	} else {
		room := parseInput(s, (end-start)/2+start)
		score := getMinScoreMaze(room, 0, 0, ex, ey)
		if score == -1 {
			result = getCoordRecursive(s, start, (end-start)/2+start, ex, ey)
		} else {
			result = getCoordRecursive(s, (end-start)/2+start, end, ex, ey)
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day18/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getMinPath(output, 1024, 70, 70))
	//fmt.Println(getCoord(output, 1024, 70, 70))
	fmt.Println(getCoordRecursive(output, 1024, len(output), 70, 70))

}
