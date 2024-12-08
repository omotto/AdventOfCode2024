package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSumSignalLocations(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "there are 14 total unique locations that contain an antinode within the bounds of the map",
			input: []string{
				"............",
				"........0...",
				".....0......",
				".......0....",
				"....0.......",
				"......A.....",
				"............",
				"............",
				"........A...",
				".........A..",
				"............",
				"............",
			},
			result: 14,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumSignalLocations(tc.input, false)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetSumSignalLocations2(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "The original example now has 34 anti-nodes",
			input: []string{
				"............",
				"........0...",
				".....0......",
				".......0....",
				"....0.......",
				"......A.....",
				"............",
				"............",
				"........A...",
				".........A..",
				"............",
				"............",
			},
			result: 34,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumSignalLocations(tc.input, true)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
