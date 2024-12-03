package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetMult(t *testing.T) {
	tcs := []struct {
		desc   string
		mults  []string
		result int
	}{
		{
			desc:   "Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5)",
			mults:  []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
			result: 161,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getMult(tc.mults)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetSumMults(t *testing.T) {
	tcs := []struct {
		desc   string
		mults  string
		result int
	}{
		{
			desc:   "Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5)",
			mults:  "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			result: 161,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumMults([]string{tc.mults})
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetSumMults2(t *testing.T) {
	tcs := []struct {
		desc   string
		mults  string
		result int
	}{
		{
			desc:   "This time, the sum of the results is 48 (2*4 + 8*5)",
			mults:  "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			result: 48,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumMults2([]string{tc.mults})
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
