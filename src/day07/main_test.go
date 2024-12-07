package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetValidRows1(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "the sum of the test values for the three equations listed above is 3749",
			input: []string{
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			result: 3749,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getValidRows(tc.input, 1)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetValidRows2(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "Adding up all six test values (the three that could be made before using only + and * plus the new three that can now be made by also using ||) produces the new total calibration result of 11387.",
			input: []string{
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			result: 11387,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getValidRows(tc.input, 2)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
