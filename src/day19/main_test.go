package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumDesigns(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "In this example, 6 of the eight designs are possible with the available towel patterns.",
			input: []string{
				"r, wr, b, g, bwu, rb, gb, br",
				"",
				"brwrr",
				"bggr",
				"gbbr",
				"rrbgbr",
				"ubwu",
				"bwurrg",
				"brgr",
				"bbrgwb",
			},
			result: 6,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumDesigns(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumPossibleDesigns(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "Adding up all of the ways the towels in this example could be arranged into the desired designs yields 16 (2 + 1 + 4 + 6 + 1 + 2).",
			input: []string{
				"r, wr, b, g, bwu, rb, gb, br",
				"",
				"brwrr",
				"bggr",
				"gbbr",
				"rrbgbr",
				"ubwu",
				"bwurrg",
				"brgr",
				"bbrgwb",
			},
			result: 16,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumPossibleDesigns(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
