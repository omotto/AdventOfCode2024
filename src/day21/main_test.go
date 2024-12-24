package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumCheats(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		times  int
		result int
	}{
		{
			desc: "In the above example, complexity of the five codes can be found by calculating 68 * 29, 60 * 980, 68 * 179, 64 * 456, and 64 * 379. Adding these together produces 126384.",
			input: []string{
				"029A",
				"980A",
				"179A",
				"456A",
				"379A",
			},
			times:  2,
			result: 126384,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumComplexities(tc.input, tc.times)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
