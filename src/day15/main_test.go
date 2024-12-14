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
			desc: "In this example, the quadrants contain 1, 3, 4, and 1 robot. Multiplying these together gives a total safety factor of 12.",
			input: []string{
				"p=0,4 v=3,-3",
				"p=6,3 v=-1,-3",
				"p=10,3 v=-1,2",
				"p=2,0 v=2,-1",
				"p=0,0 v=1,3",
				"p=3,0 v=-2,-2",
				"p=7,6 v=-1,-3",
				"p=3,0 v=-1,-2",
				"p=9,3 v=2,3",
				"p=7,3 v=-1,2",
				"p=2,4 v=2,-3",
				"p=9,5 v=-3,-3",
			},
			result: 12,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := countRobots(tc.input, 11, 7, 100)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
