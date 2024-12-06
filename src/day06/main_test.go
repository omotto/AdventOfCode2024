package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSumGuardPositions(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "In this example, the guard will visit 41 distinct positions on your map.",
			input: []string{
				"....#.....",
				".........#",
				"..........",
				"..#.......",
				".......#..",
				"..........",
				".#..^.....",
				"........#.",
				"#.........",
				"......#...",
			},
			result: 41,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got, _ := getGuardPositions(tc.input)
			if diff := cmp.Diff(tc.result, len(got)); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetSumGuardStuckPositions(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "In this example, there are 6 different positions you could choose.",
			input: []string{
				"....#.....",
				".........#",
				"..........",
				"..#.......",
				".......#..",
				"..........",
				".#..^.....",
				"........#.",
				"#.........",
				"......#...",
			},
			result: 6,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			path, _ := getGuardPositions(tc.input)
			got := getSumGuardStuckPositions(tc.input, path)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
