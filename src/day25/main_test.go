package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumValidKeys(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "So, in this example, the number of unique lock/key pairs that fit together without overlapping in any column is 3.",
			input: []string{
				"#####",
				".####",
				".####",
				".####",
				".#.#.",
				".#...",
				".....",
				"",
				"#####",
				"##.##",
				".#.##",
				"...##",
				"...#.",
				"...#.",
				".....",
				"",
				".....",
				"#....",
				"#....",
				"#...#",
				"#.#.#",
				"#.###",
				"#####",
				"",
				".....",
				".....",
				"#.#..",
				"###..",
				"###.#",
				"###.#",
				"#####",
				"",
				".....",
				".....",
				".....",
				"#....",
				"#.#..",
				"#.#.#",
				"#####",
				"",
			},
			result: 3,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumValidKeys(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
