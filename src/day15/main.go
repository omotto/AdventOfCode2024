package main

import (
	"bytes"
	"fmt"
	"path/filepath"

	"advent2024/pkg/file"
)

func parseInput(s []string) ([]string, int, int, string) {
	var (
		room         []string
		sx, sy       int
		instructions bytes.Buffer
		i            int
	)
	for y, line := range s {
		if len(line) < 3 {
			i = y + 1
			break
		}
		room = append(room, line)
		if sx == 0 && sy == 0 {
			for x := 0; x < len(line); x++ {
				if line[x] == '@' {
					sx, sy = x, y
					break
				}
			}
		}
	}
	for ; i < len(s); i++ {
		instructions.WriteString(s[i])
	}
	return room, sx, sy, instructions.String()
}

func run(room []string, sx, sy int, instruction byte) ([]string, int, int) {
	var x, y int
	switch instruction {
	case '^':
		switch room[sy-1][sx] {
		case '.':
			x, y = sx, sy-1
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy-1] = room[sy-1][:sx] + "@" + room[sy-1][sx+1:]
		case '#':
			x, y = sx, sy
		case 'O':
			x, y = sx, sy
			for i := sy - 1; i > 0; i-- {
				if room[i][x] == '#' {
					break
				}
				if room[i][x] == '.' {
					x, y = sx, sy-1
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy-1] = room[sy-1][:sx] + "@" + room[sy-1][sx+1:]
					room[i] = room[i][:sx] + "O" + room[i][sx+1:]
					break
				}
			}
		}
	case '<':
		switch room[sy][sx-1] {
		case '.':
			x, y = sx-1, sy
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy] = room[sy][:sx-1] + "@" + room[sy][sx:]
		case '#':
			x, y = sx, sy
		case 'O':
			x, y = sx, sy
			for i := sx - 1; i > 0; i-- {
				if room[sy][i] == '#' {
					break
				}
				if room[sy][i] == '.' {
					x, y = sx-1, sy
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy] = room[sy][:sx-1] + "@" + room[sy][sx:]
					room[sy] = room[sy][:i] + "O" + room[sy][i+1:]
					break
				}
			}
		}
	case '>':
		switch room[sy][sx+1] {
		case '.':
			x, y = sx+1, sy
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy] = room[sy][:sx+1] + "@" + room[sy][sx+2:]
		case '#':
			x, y = sx, sy
		case 'O':
			x, y = sx, sy
			for i := sx + 1; i < len(room[sy]); i++ {
				if room[sy][i] == '#' {
					break
				}
				if room[sy][i] == '.' {
					x, y = sx+1, sy
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy] = room[sy][:sx+1] + "@" + room[sy][sx+2:]
					room[sy] = room[sy][:i] + "O" + room[sy][i+1:]
					break
				}
			}
		}
	case 'v':
		switch room[sy+1][sx] {
		case '.':
			x, y = sx, sy+1
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy+1] = room[sy+1][:sx] + "@" + room[sy+1][sx+1:]
		case '#':
			x, y = sx, sy
		case 'O':
			x, y = sx, sy
			for i := sy + 1; i < len(room); i++ {
				if room[i][x] == '#' {
					break
				}
				if room[i][x] == '.' {
					x, y = sx, sy+1
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy+1] = room[sy+1][:sx] + "@" + room[sy+1][sx+1:]
					room[i] = room[i][:sx] + "O" + room[i][sx+1:]
					break
				}
			}
		}
	}
	return room, x, y
}

func countSumBoxesCoords(s []string) int {
	room, sx, sy, instructions := parseInput(s)
	for i := 0; i < len(instructions); i++ {
		room, sx, sy = run(room, sx, sy, instructions[i])
	}
	result := 0
	for y := 0; y < len(room); y++ {
		for x := 0; x < len(room[y]); x++ {
			if room[y][x] == 'O' {
				result += x + y*100
			}
		}
	}
	return result
}

func parseInputX2(s []string) ([]string, int, int, string) {
	var (
		room         []string
		instructions bytes.Buffer
		i            int
	)
	for y, line := range s {
		if len(line) < 3 {
			i = y + 1
			break
		}
		var buffer bytes.Buffer
		for x := 0; x < len(line); x++ {
			switch line[x] {
			case '.':
				buffer.WriteString("..")
			case '#':
				buffer.WriteString("##")
			case '@':
				buffer.WriteString("@.")
			case 'O':
				buffer.WriteString("[]")
			}
		}
		room = append(room, buffer.String())
	}
	for ; i < len(s); i++ {
		instructions.WriteString(s[i])
	}
	for y := 0; y < len(room); y++ {
		for x := 0; x < len(room[y]); x++ {
			if room[y][x] == '@' {
				return room, x, y, instructions.String()
			}
		}
	}
	return room, 0, 0, instructions.String()
}

func check(room []string, y, x1, x2, dir int, objects *[][2]int) bool {
	*objects = append(*objects, [2]int{x1, y})
	if room[y+dir][x1] == '#' || room[y+dir][x2] == '#' {
		return false
	}
	if room[y+dir][x1] == '.' && room[y+dir][x2] == '.' {
		return true
	}
	var v []bool
	if room[y+dir][x1] == '[' {
		v = append(v, check(room, y+dir, x1, x2, dir, objects))
	}
	if room[y+dir][x1] == ']' {
		v = append(v, check(room, y+dir, x1-1, x2-1, dir, objects))
	}
	if room[y+dir][x2] == '[' {
		v = append(v, check(room, y+dir, x1+1, x2+1, dir, objects))
	}
	for _, vv := range v {
		if vv == false {
			return false
		}
	}
	return true
}

func move(room []string, dir int, objects [][2]int) []string {
	// Remove from room
	for _, object := range objects {
		x := object[0]
		y := object[1]
		room[y] = room[y][:x] + ".." + room[y][x+2:]
	}
	// Add to the room in the new positions
	for _, object := range objects {
		x := object[0]
		y := object[1]
		room[y+dir] = room[y+dir][:x] + "[]" + room[y+dir][x+2:]
	}
	return room
}

func runX2(room []string, sx, sy int, instruction byte) ([]string, int, int) {
	var x, y int
	switch instruction {
	case '^':
		switch room[sy-1][sx] {
		case '.':
			x, y = sx, sy-1
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy-1] = room[sy-1][:sx] + "@" + room[sy-1][sx+1:]
		case '#':
			x, y = sx, sy
		case '[':
			x, y = sx, sy
			var objects [][2]int
			if check(room, sy-1, x, x+1, -1, &objects) {
				room = move(room, -1, objects)
				room[y] = room[y][:sx] + "." + room[y][sx+1:]
				y--
				room[y] = room[y][:sx] + "@" + room[y][sx+1:]
			}
		case ']':
			x, y = sx, sy
			var objects [][2]int
			if check(room, sy-1, x-1, x, -1, &objects) {
				room = move(room, -1, objects)
				room[y] = room[y][:sx] + "." + room[y][sx+1:]
				y--
				room[y] = room[y][:sx] + "@" + room[y][sx+1:]
			}
		}
	case '<':
		switch room[sy][sx-1] {
		case '.':
			x, y = sx-1, sy
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy] = room[sy][:sx-1] + "@" + room[sy][sx:]
		case '#':
			x, y = sx, sy
		case ']':
			x, y = sx, sy
			for i := sx - 1; i > 0; i-- {
				if room[sy][i] == '#' {
					break
				}
				if room[sy][i] == '.' {
					x, y = sx-1, sy
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy] = room[sy][:sx-1] + "@" + room[sy][sx:]
					for j := i; j < sx-1; j += 2 {
						room[sy] = room[sy][:j] + "[" + room[sy][j+1:]
						room[sy] = room[sy][:j+1] + "]" + room[sy][j+2:]
					}
					break
				}
			}
		}
	case '>':
		switch room[sy][sx+1] {
		case '.':
			x, y = sx+1, sy
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy] = room[sy][:sx+1] + "@" + room[sy][sx+2:]
		case '#':
			x, y = sx, sy
		case '[':
			x, y = sx, sy
			for i := sx + 1; i < len(room[sy]); i++ {
				if room[sy][i] == '#' {
					break
				}
				if room[sy][i] == '.' {
					x, y = sx+1, sy
					room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
					room[sy] = room[sy][:sx+1] + "@" + room[sy][sx+2:]
					for j := i; j > sx+1; j -= 2 {
						room[sy] = room[sy][:j] + "]" + room[sy][j+1:]
						room[sy] = room[sy][:j-1] + "[" + room[sy][j:]
					}
					break
				}
			}
		}
	case 'v':
		switch room[sy+1][sx] {
		case '.':
			x, y = sx, sy+1
			room[sy] = room[sy][:sx] + "." + room[sy][sx+1:]
			room[sy+1] = room[sy+1][:sx] + "@" + room[sy+1][sx+1:]
		case '#':
			x, y = sx, sy
		case '[':
			x, y = sx, sy
			var objects [][2]int
			if check(room, sy+1, x, x+1, +1, &objects) {
				room = move(room, +1, objects)
				room[y] = room[y][:sx] + "." + room[y][sx+1:]
				y++
				room[y] = room[y][:sx] + "@" + room[y][sx+1:]
			}
		case ']':
			x, y = sx, sy
			var objects [][2]int
			if check(room, sy+1, x-1, x, +1, &objects) {
				room = move(room, +1, objects)
				room[y] = room[y][:sx] + "." + room[y][sx+1:]
				y++
				room[y] = room[y][:sx] + "@" + room[y][sx+1:]
			}
		}
	}
	return room, x, y
}

func countSumBoxesCoordsX2(s []string) int {
	room, sx, sy, instructions := parseInputX2(s)
	for i := 0; i < len(instructions); i++ {
		room, sx, sy = runX2(room, sx, sy, instructions[i])
	}
	result := 0
	for y := 0; y < len(room); y++ {
		for x := 0; x < len(room[y]); x++ {
			if room[y][x] == '[' {
				result += x + y*100
			}
		}
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day15/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(countSumBoxesCoords(output))
	fmt.Println(countSumBoxesCoordsX2(output))
}
