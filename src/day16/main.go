package main

import (
	"advent2024/pkg/file"
	"fmt"
	"path/filepath"
	"slices"
)

const (
	UP    = 1
	DOWN  = 2
	RIGHT = 3
	LEFT  = 4
)

func getStartEnd(s []string) (int, int, int, int) {
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

func getMinScoreMaze(room []string, sx, sy int) map[string]int {
	type Tile struct {
		x, y, dir, score int
	}
	visited := map[string]int{
		fmt.Sprintf("%d:%d:%d", sx, sy, RIGHT): 0,
		/*		fmt.Sprintf("%d:%d:%d", sx, sy, UP):    0,
				fmt.Sprintf("%d:%d:%d", sx, sy, DOWN):  0,
				fmt.Sprintf("%d:%d:%d", sx, sy, LEFT):  0,*/
	}
	queue := make([]Tile, 0)
	queue = append(queue, []Tile{
		{
			x:     sx,
			y:     sy,
			dir:   RIGHT,
			score: 0,
		}, /*, {
			x:     sx,
			y:     sy,
			dir:   LEFT,
			score: 0,
		},
		{
			x:     sx,
			y:     sy,
			dir:   UP,
			score: 0,
		},
		{
			x:     sx,
			y:     sy,
			dir:   DOWN,
			score: 0,
		},*/
	}...)
	for len(queue) > 0 {
		tile := queue[0]  // Get first
		queue = queue[1:] // Remove it
		if v, ok := visited[fmt.Sprintf("%d:%d:%d", tile.x, tile.y, tile.dir)]; ok && v < tile.score {
			continue
		}
		// Possible movements: Same direction or 90º
		var newX, newY int
		switch tile.dir {
		case UP:
			newX, newY = tile.x, tile.y-1
		case DOWN:
			newX, newY = tile.x, tile.y+1
		case LEFT:
			newX, newY = tile.x-1, tile.y
		case RIGHT:
			newX, newY = tile.x+1, tile.y
		}
		// Check same direction
		var newScore int
		if room[newY][newX] != '#' {
			newScore = tile.score + 1
			if v, ok := visited[fmt.Sprintf("%d:%d:%d", newX, newY, tile.dir)]; (ok && newScore < v) || !ok {
				visited[fmt.Sprintf("%d:%d:%d", newX, newY, tile.dir)] = newScore
				queue = append(queue, Tile{
					x:     newX,
					y:     newY,
					dir:   tile.dir,
					score: newScore,
				})
			}
		}
		// Check 90º
		switch tile.dir {
		case UP, DOWN:
			newScore = tile.score + 1000
			if v, ok := visited[fmt.Sprintf("%d:%d:%d", tile.x, tile.y, LEFT)]; (ok && newScore < v) || !ok {
				visited[fmt.Sprintf("%d:%d:%d", tile.x, tile.y, LEFT)] = newScore
				queue = append(queue, Tile{
					x:     tile.x,
					y:     tile.y,
					dir:   LEFT,
					score: newScore,
				})
			}
			if v, ok := visited[fmt.Sprintf("%d:%d:%d", tile.x, tile.y, RIGHT)]; (ok && newScore < v) || !ok {
				visited[fmt.Sprintf("%d:%d:%d", tile.x, tile.y, RIGHT)] = newScore
				queue = append(queue, Tile{
					x:     tile.x,
					y:     tile.y,
					dir:   RIGHT,
					score: newScore,
				})
			}
		case LEFT, RIGHT:
			newScore = tile.score + 1000
			if v, ok := visited[fmt.Sprintf("%d:%d:%d", tile.x, tile.y, UP)]; (ok && newScore < v) || !ok {
				visited[fmt.Sprintf("%d:%d:%d", tile.x, tile.y, UP)] = newScore
				queue = append(queue, Tile{
					x:     tile.x,
					y:     tile.y,
					dir:   UP,
					score: newScore,
				})
			}
			if v, ok := visited[fmt.Sprintf("%d:%d:%d", tile.x, tile.y, DOWN)]; (ok && newScore < v) || !ok {
				visited[fmt.Sprintf("%d:%d:%d", tile.x, tile.y, DOWN)] = newScore
				queue = append(queue, Tile{
					x:     tile.x,
					y:     tile.y,
					dir:   DOWN,
					score: newScore,
				})
			}
		}
	}
	return visited
}

func getLowerScoreMaze(s []string) int {
	sx, sy, ex, ey := getStartEnd(s)
	visitedScoredTiles := getMinScoreMaze(s, sx, sy)
	var endScoreValues []int
	if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", ex, ey, UP)]; ok {
		endScoreValues = append(endScoreValues, v)
	}
	if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", ex, ey, DOWN)]; ok {
		endScoreValues = append(endScoreValues, v)
	}
	if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", ex, ey, LEFT)]; ok {
		endScoreValues = append(endScoreValues, v)
	}
	if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", ex, ey, RIGHT)]; ok {
		endScoreValues = append(endScoreValues, v)
	}
	return slices.Min(endScoreValues)
}

func check(room []string, sx, sy, x, y, dir, minScore int, visitedScoredTiles map[string]int, tiles map[string]struct{}) {
	if !(x == sx && y == sy) {
		// Start condition (from the END)
		if dir == 0 {
			tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
			var endScoreValues []int
			if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, UP)]; ok {
				endScoreValues = append(endScoreValues, v)
			}
			if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, DOWN)]; ok {
				endScoreValues = append(endScoreValues, v)
			}
			if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, LEFT)]; ok {
				endScoreValues = append(endScoreValues, v)
			}
			if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, RIGHT)]; ok {
				endScoreValues = append(endScoreValues, v)
			}
			lowerScore := slices.Min(endScoreValues)
			if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, UP)]; ok && v == lowerScore {
				check(room, sx, sy, x, y+1, UP, lowerScore, visitedScoredTiles, tiles)
			}
			if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, DOWN)]; ok && v == lowerScore {
				check(room, sx, sy, x, y-1, DOWN, lowerScore, visitedScoredTiles, tiles)
			}
			if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, LEFT)]; ok && v == lowerScore {
				check(room, sx, sy, x+1, y, LEFT, lowerScore, visitedScoredTiles, tiles)
			}
			if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, RIGHT)]; ok && v == lowerScore {
				check(room, sx, sy, x-1, y, RIGHT, lowerScore, visitedScoredTiles, tiles)
			}
		} else {
			switch dir {
			case UP:
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, UP)]; ok && v == minScore-1 && room[y+1][x] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x, y+1, UP, minScore-1, visitedScoredTiles, tiles)
				}
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, LEFT)]; ok && v == minScore-1001 && room[y][x+1] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x+1, y, LEFT, minScore-1001, visitedScoredTiles, tiles)
				}
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, RIGHT)]; ok && v == minScore-1001 && room[y][x-1] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x-1, y, RIGHT, minScore-1001, visitedScoredTiles, tiles)
				}
			case DOWN:
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, DOWN)]; ok && v == minScore-1 && room[y-1][x] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x, y-1, DOWN, minScore-1, visitedScoredTiles, tiles)
				}
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, LEFT)]; ok && v == minScore-1001 && room[y][x+1] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x+1, y, LEFT, minScore-1001, visitedScoredTiles, tiles)
				}
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, RIGHT)]; ok && v == minScore-1001 && room[y][x-1] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x-1, y, RIGHT, minScore-1001, visitedScoredTiles, tiles)
				}
			case LEFT:
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, LEFT)]; ok && v == minScore-1 && room[y][x+1] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x+1, y, LEFT, minScore-1, visitedScoredTiles, tiles)
				}
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, UP)]; ok && v == minScore-1001 && room[y+1][x] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x, y+1, UP, minScore-1001, visitedScoredTiles, tiles)
				}
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, DOWN)]; ok && v == minScore-1001 && room[y-1][x] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x, y-1, DOWN, minScore-1001, visitedScoredTiles, tiles)
				}
			case RIGHT:
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, RIGHT)]; ok && v == minScore-1 && room[y][x-1] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x-1, y, RIGHT, minScore-1, visitedScoredTiles, tiles)
				}
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, UP)]; ok && v == minScore-1001 && room[y+1][x] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x, y+1, UP, minScore-1001, visitedScoredTiles, tiles)
				}
				if v, ok := visitedScoredTiles[fmt.Sprintf("%d:%d:%d", x, y, DOWN)]; ok && v == minScore-1001 && room[y-1][x] != '#' {
					tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
					check(room, sx, sy, x, y-1, DOWN, minScore-1001, visitedScoredTiles, tiles)
				}
			}
		}
	} else {
		tiles[fmt.Sprintf("%d:%d", x, y)] = struct{}{}
	}
}

func getNumTilesLowerScoreMaze(s []string) int {
	sx, sy, ex, ey := getStartEnd(s)
	visitedScoredTiles := getMinScoreMaze(s, sx, sy)
	tiles := map[string]struct{}{}
	check(s, sx, sy, ex, ey, 0, 0, visitedScoredTiles, tiles)
	return len(tiles)
}

func main() {
	absPathName, _ := filepath.Abs("src/day16/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getLowerScoreMaze(output))
	fmt.Println(getNumTilesLowerScoreMaze(output))
}
