package main

import (
	"bytes"
	"fmt"
	"path/filepath"

	"advent2024/pkg/file"
)

func countRobots(s []string, maxX, maxY, times int) int {
	var (
		pX, pY, vX, vY int
		a, b, c, d     int
	)
	for _, line := range s {
		_, _ = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pX, &pY, &vX, &vY)
		newPx := pX + vX*times
		newPy := pY + vY*times
		newPx = newPx % maxX
		if newPx < 0 {
			newPx = maxX + newPx
		}
		newPy = newPy % maxY
		if newPy < 0 {
			newPy = maxY + newPy
		}
		if newPx < (maxX-1)/2 {
			if newPy < (maxY-1)/2 {
				a++
			} else if newPy > (maxY-1)/2 {
				b++
			}
		} else if newPx > (maxX-1)/2 {
			if newPy < (maxY-1)/2 {
				c++
			} else if newPy > (maxY-1)/2 {
				d++
			}
		}
	}
	return a * b * c * d
}

type Robot struct {
	pX, pY, vX, vY int
}

func printRobots(points map[string]struct{}, maxX, maxY int) {
	for y := 0; y < maxY; y++ {
		var buffer bytes.Buffer
		for x := 0; x < maxX; x++ {
			if _, ok := points[fmt.Sprintf("%d:%d", x, y)]; ok {
				buffer.WriteString("#")
			} else {
				buffer.WriteString(".")
			}
		}
		buffer.WriteString("\r\n")
		fmt.Println(buffer.String())
	}
}

func countSeconds(s []string, maxX, maxY int) int {
	robots := make([]Robot, len(s))
	for idx, line := range s {
		_, _ = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robots[idx].pX, &robots[idx].pY, &robots[idx].vX, &robots[idx].vY)
	}
	seconds := 0
	for {
		// Move
		for idx, robot := range robots {
			newPx := robot.pX + robot.vX
			newPy := robot.pY + robot.vY
			newPx = newPx % maxX
			if newPx < 0 {
				newPx = maxX + newPx
			}
			newPy = newPy % maxY
			if newPy < 0 {
				newPy = maxY + newPy
			}
			robots[idx].pX = newPx
			robots[idx].pY = newPy
		}
		seconds++
		// Check
		points := map[string]struct{}{}
		for _, robot := range robots {
			points[fmt.Sprintf("%d:%d", robot.pX, robot.pY)] = struct{}{}
		}
		// NO robots overlapping
		if len(points) == len(robots) {
			printRobots(points, maxX, maxY)
			break
		}
	}
	return seconds
}

func countSeconds2(s []string, maxX, maxY int) int {
	robots := make([]Robot, len(s))
	for idx, line := range s {
		_, _ = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robots[idx].pX, &robots[idx].pY, &robots[idx].vX, &robots[idx].vY)
	}
	seconds := 0
	maxWeight := 0
	for i := 0; i < maxX*maxY; i++ { // sequence is repeated every
		// Move
		for idx, robot := range robots {
			newPx := robot.pX + robot.vX
			newPy := robot.pY + robot.vY
			newPx = newPx % maxX
			if newPx < 0 {
				newPx = maxX + newPx
			}
			newPy = newPy % maxY
			if newPy < 0 {
				newPy = maxY + newPy
			}
			robots[idx].pX = newPx
			robots[idx].pY = newPy
		}
		// Check
		points := map[string]struct{}{}
		for _, robot := range robots {
			points[fmt.Sprintf("%d:%d", robot.pX, robot.pY)] = struct{}{}
		}
		weight := 0
		for _, robot := range robots {
			if _, ok := points[fmt.Sprintf("%d:%d", robot.pX+1, robot.pY)]; ok {
				weight++
			}
			if _, ok := points[fmt.Sprintf("%d:%d", robot.pX-1, robot.pY)]; ok {
				weight++
			}
			if _, ok := points[fmt.Sprintf("%d:%d", robot.pX, robot.pY+1)]; ok {
				weight++
			}
			if _, ok := points[fmt.Sprintf("%d:%d", robot.pX, robot.pY-1)]; ok {
				weight++
			}
			if _, ok := points[fmt.Sprintf("%d:%d", robot.pX-1, robot.pY-1)]; ok {
				weight++
			}
			if _, ok := points[fmt.Sprintf("%d:%d", robot.pX+1, robot.pY+1)]; ok {
				weight++
			}
			if _, ok := points[fmt.Sprintf("%d:%d", robot.pX+1, robot.pY-1)]; ok {
				weight++
			}
			if _, ok := points[fmt.Sprintf("%d:%d", robot.pX-1, robot.pY+1)]; ok {
				weight++
			}
		}
		if weight > maxWeight {
			maxWeight = weight
			seconds = i + 1
			printRobots(points, maxX, maxY)
		}
	}
	return seconds
}

func main() {
	absPathName, _ := filepath.Abs("src/day14/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(countRobots(output, 101, 103, 100))
	//fmt.Println(countSeconds(output, 101, 103))
	fmt.Println(countSeconds2(output, 101, 103))
}
