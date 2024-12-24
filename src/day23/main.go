package main

import (
	"advent2024/pkg/file"
	"fmt"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

func getConnections(s []string) map[string][]string {
	connections := map[string][]string{}
	for _, line := range s {
		parts := strings.Split(line, "-")
		connections[parts[0]] = append(connections[parts[0]], parts[1])
		connections[parts[1]] = append(connections[parts[1]], parts[0])
	}
	return connections
}

func getSumTLANNetworks(s []string) int {
	connections := getConnections(s)
	checked := map[string]struct{}{}
	for c1, v := range connections {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				c2 := v[i]
				c3 := v[j]
				idx1 := slices.IndexFunc(connections[c2], func(s string) bool { return s == c3 })
				idx2 := slices.IndexFunc(connections[c3], func(s string) bool { return s == c2 })
				if idx1 != -1 || idx2 != -1 {
					if c1[0] == 't' || c2[0] == 't' || c3[0] == 't' {
						str := []string{c1, c2, c3}
						sort.Slice(str, func(i, j int) bool {
							return str[i] < str[j]
						})
						checked[strings.Join(str, "-")] = struct{}{}
					}
				}
			}
		}
	}
	return len(checked)
}

func getPassword(s []string) string {
	connections := getConnections(s)
	largest := []string{}
	for lan, lans := range connections {
		cn := []string{lan}
		for _, v := range lans {
			add := true
			for _, n := range cn {
				if idx := slices.IndexFunc(connections[v], func(s string) bool { return s == n }); idx == -1 {
					add = false
					break
				}
			}
			if add {
				cn = append(cn, v)
			}
		}
		if len(cn) > len(largest) {
			largest = cn
		}
	}
	sort.Slice(largest, func(i, j int) bool {
		return largest[i] < largest[j]
	})
	return strings.Join(largest, ",")
}

func main() {
	absPathName, _ := filepath.Abs("src/day23/input.txt")
	output, _ := file.ReadInput(absPathName)

	fmt.Println(getSumTLANNetworks(output))
	fmt.Println(getPassword(output))
}
