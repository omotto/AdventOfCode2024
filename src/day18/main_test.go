package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMinPath(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "After just 12 bytes have corrupted locations in your memory space, the shortest path from the top left corner to the exit would take 22 steps.",
			input: []string{
				"5,4",
				"4,2",
				"4,5",
				"3,0",
				"2,1",
				"6,3",
				"2,4",
				"1,5",
				"0,6",
				"3,3",
				"2,6",
				"5,1",
				"1,2",
				"5,5",
				"2,5",
				"6,5",
				"1,4",
				"0,4",
				"6,4",
				"1,1",
				"6,1",
				"1,0",
				"0,5",
				"1,6",
				"2,0",
			},
			result: 22,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getMinPath(tc.input, 12, 6, 6)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetCoord(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result string
	}{
		{
			desc: "So, in this example, the coordinates of the first byte that prevents the exit from being reachable are 6,1.",
			input: []string{
				"5,4",
				"4,2",
				"4,5",
				"3,0",
				"2,1",
				"6,3",
				"2,4",
				"1,5",
				"0,6",
				"3,3",
				"2,6",
				"5,1",
				"1,2",
				"5,5",
				"2,5",
				"6,5",
				"1,4",
				"0,4",
				"6,4",
				"1,1",
				"6,1",
				"1,0",
				"0,5",
				"1,6",
				"2,0",
			},
			result: "6,1",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getCoord(tc.input, 12, 6, 6)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
