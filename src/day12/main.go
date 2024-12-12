package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"advent2024/pkg/file"
)

const key = "%d:%d"

type Region struct {
	area      int
	perimeter int
	positions map[string]struct{}
}

func getRegion(s []string, x, y int, checked map[string]struct{}, region *Region) {
	region.area++
	region.perimeter += 4
	region.positions[fmt.Sprintf(key, x, y)] = struct{}{}
	if x > 0 && s[y][x-1] == s[y][x] {
		if _, ok := checked[fmt.Sprintf(key, x-1, y)]; !ok {
			checked[fmt.Sprintf(key, x-1, y)] = struct{}{}
			getRegion(s, x-1, y, checked, region)
		}
		region.perimeter--
	}
	if x < len(s[0])-1 && s[y][x+1] == s[y][x] {
		if _, ok := checked[fmt.Sprintf(key, x+1, y)]; !ok {
			checked[fmt.Sprintf(key, x+1, y)] = struct{}{}
			getRegion(s, x+1, y, checked, region)
		}
		region.perimeter--

	}
	if y > 0 && s[y-1][x] == s[y][x] {
		if _, ok := checked[fmt.Sprintf(key, x, y-1)]; !ok {
			checked[fmt.Sprintf(key, x, y-1)] = struct{}{}
			getRegion(s, x, y-1, checked, region)
		}
		region.perimeter--
	}
	if y < len(s)-1 && s[y+1][x] == s[y][x] {
		if _, ok := checked[fmt.Sprintf(key, x, y+1)]; !ok {
			checked[fmt.Sprintf(key, x, y+1)] = struct{}{}
			getRegion(s, x, y+1, checked, region)
		}
		region.perimeter--
	}
}

func calculateRegions(s []string) []Region {
	var regions []Region
	checked := make(map[string]struct{})
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if _, ok := checked[fmt.Sprintf(key, x, y)]; !ok {
				region := Region{
					positions: make(map[string]struct{}),
				}
				checked[fmt.Sprintf(key, x, y)] = struct{}{}
				getRegion(s, x, y, checked, &region)
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func getPrices(s []string) int {
	regions := calculateRegions(s)
	result := 0
	for _, region := range regions {
		result += region.area * region.perimeter
	}
	return result
}

func getSides(region map[string]struct{}) int {
	sides := 0
	for k, _ := range region {
		parts := strings.Split(k, ":")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		_, left := region[fmt.Sprintf(key, x-1, y)]
		_, right := region[fmt.Sprintf(key, x+1, y)]
		_, up := region[fmt.Sprintf(key, x, y-1)]
		_, down := region[fmt.Sprintf(key, x, y+1)]
		_, upperLeft := region[fmt.Sprintf(key, x-1, y-1)]
		_, upperRight := region[fmt.Sprintf(key, x+1, y-1)]
		_, downLeft := region[fmt.Sprintf(key, x-1, y+1)]
		_, downRight := region[fmt.Sprintf(key, x+1, y+1)]
		if !left && !up {
			sides++
		}
		if !right && !up {
			sides++
		}
		if !left && !down {
			sides++
		}
		if !right && !down {
			sides++
		}
		if !upperRight && up && right {
			sides++
		}
		if !upperLeft && up && left {
			sides++
		}
		if !downLeft && down && left {
			sides++
		}
		if !downRight && down && right {
			sides++
		}
	}
	return sides
}

func getPrices2(s []string) int {
	regions := calculateRegions(s)
	result := 0
	for _, region := range regions {
		result += region.area * getSides(region.positions)
	}
	return result
}

func main() {
	absPathName, _ := filepath.Abs("src/day12/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getPrices(output))
	fmt.Println(getPrices2(output))
}
