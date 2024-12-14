package main

import (
	"fmt"
	"path/filepath"

	"advent2024/pkg/file"
)

func main() {
	absPathName, _ := filepath.Abs("src/day15/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(countRobots(output, 101, 103, 100))
}
