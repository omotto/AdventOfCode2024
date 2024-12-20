package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumCheats(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		edge   int
		result int
	}{
		{
			desc: "There is one cheat that saves 64 picoseconds.",
			input: []string{
				"###############",
				"#...#...#.....#",
				"#.#.#.#.#.###.#",
				"#S#...#.#.#...#",
				"#######.#.#.###",
				"#######.#.#...#",
				"#######.#.###.#",
				"###..E#...#...#",
				"###.#######.###",
				"#...###...#...#",
				"#.#####.#.###.#",
				"#.#...#.#.#...#",
				"#.#.#.#.#.#.###",
				"#...#...#...###",
				"###############",
			},
			edge:   64,
			result: 1,
		},
		{
			desc: "There is two cheats that saves at least 40 picoseconds.",
			input: []string{
				"###############",
				"#...#...#.....#",
				"#.#.#.#.#.###.#",
				"#S#...#.#.#...#",
				"#######.#.#.###",
				"#######.#.#...#",
				"#######.#.###.#",
				"###..E#...#...#",
				"###.#######.###",
				"#...###...#...#",
				"#.#####.#.###.#",
				"#.#...#.#.#...#",
				"#.#.#.#.#.#.###",
				"#...#...#...###",
				"###############",
			},
			edge:   40,
			result: 2,
		},
		{
			desc: "There are 8 cheats that save at least 12 picoseconds.",
			input: []string{
				"###############",
				"#...#...#.....#",
				"#.#.#.#.#.###.#",
				"#S#...#.#.#...#",
				"#######.#.#.###",
				"#######.#.#...#",
				"#######.#.###.#",
				"###..E#...#...#",
				"###.#######.###",
				"#...###...#...#",
				"#.#####.#.###.#",
				"#.#...#.#.#...#",
				"#.#.#.#.#.#.###",
				"#...#...#...###",
				"###############",
			},
			edge:   12,
			result: 8,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumCheats(tc.input, tc.edge)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumCheats2(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		edge   int
		result int
	}{
		{
			desc: "There are 3 cheats that save 76 picoseconds.",
			input: []string{
				"###############",
				"#...#...#.....#",
				"#.#.#.#.#.###.#",
				"#S#...#.#.#...#",
				"#######.#.#.###",
				"#######.#.#...#",
				"#######.#.###.#",
				"###..E#...#...#",
				"###.#######.###",
				"#...###...#...#",
				"#.#####.#.###.#",
				"#.#...#.#.#...#",
				"#.#.#.#.#.#.###",
				"#...#...#...###",
				"###############",
			},
			edge:   76,
			result: 3,
		},
		{
			desc: "There are 29 cheats that save at least 72 picoseconds.",
			input: []string{
				"###############",
				"#...#...#.....#",
				"#.#.#.#.#.###.#",
				"#S#...#.#.#...#",
				"#######.#.#.###",
				"#######.#.#...#",
				"#######.#.###.#",
				"###..E#...#...#",
				"###.#######.###",
				"#...###...#...#",
				"#.#####.#.###.#",
				"#.#...#.#.#...#",
				"#.#.#.#.#.#.###",
				"#...#...#...###",
				"###############",
			},
			edge:   72,
			result: 29,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumCheats2(tc.input, tc.edge)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
