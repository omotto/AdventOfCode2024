package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetTokens(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "So, the most prizes you could possibly win is two; the minimum tokens you would have to spend to win all (two) prizes is 480.",
			input: []string{
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=8400, Y=5400",
				"",
				"Button A: X+26, Y+66",
				"Button B: X+67, Y+21",
				"Prize: X=12748, Y=12176",
				"",
				"Button A: X+17, Y+86",
				"Button B: X+84, Y+37",
				"Prize: X=7870, Y=6450",
				"",
				"Button A: X+69, Y+23",
				"Button B: X+27, Y+71",
				"Prize: X=18641, Y=10279",
			},
			result: 480,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getTokens(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetTokens2(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "So, the most prizes you could possibly win is two; the minimum tokens you would have to spend to win all (two) prizes is 480.",
			input: []string{
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=8400, Y=5400",
				"",
				"Button A: X+26, Y+66",
				"Button B: X+67, Y+21",
				"Prize: X=12748, Y=12176",
				"",
				"Button A: X+17, Y+86",
				"Button B: X+84, Y+37",
				"Prize: X=7870, Y=6450",
				"",
				"Button A: X+69, Y+23",
				"Button B: X+27, Y+71",
				"Prize: X=18641, Y=10279",
			},
			result: 480,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getTokens2(tc.input, 0)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
