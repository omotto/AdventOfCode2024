package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNextSecretNumber(t *testing.T) {
	tcs := []struct {
		desc   string
		input  int
		times  int
		result int
	}{
		{
			desc:   "15887950.",
			input:  123,
			times:  1,
			result: 15887950,
		},
		{
			desc:   "16495136.",
			input:  123,
			times:  2,
			result: 16495136,
		},
		{
			desc:   "12249484.",
			input:  123,
			times:  8,
			result: 12249484,
		},
		{
			desc:   "7753432.",
			input:  123,
			times:  9,
			result: 7753432,
		},
		{
			desc:   "5908254.",
			input:  123,
			times:  10,
			result: 5908254,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := nextSecretNumber(tc.input, tc.times)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumBananas(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "So, by asking the monkey to sell the first time each buyer's prices go down 2, then up 1, then down 1, then up 3, you would get 23 (7 + 7 + 9) bananas!.",
			input: []string{
				"1",
				"2",
				"3",
				"2024",
			},
			result: 23,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumBananas(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
