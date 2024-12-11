package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetNumStones(t *testing.T) {
	tcs := []struct {
		desc   string
		input  string
		result int
	}{
		{
			desc:   "In this example, after blinking six times, you would have 22 stones. After blinking 25 times, you would have 55312 stones!",
			input:  "125 17",
			result: 55312,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumStones([]string{tc.input}, 25)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumStones2(t *testing.T) {
	tcs := []struct {
		desc   string
		input  string
		result int
	}{
		{
			desc:   "In this example, after blinking six times, you would have 22 stones. After blinking 25 times, you would have 55312 stones!",
			input:  "125 17",
			result: 55312,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumStones2([]string{tc.input}, 25)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetNumStones3(t *testing.T) {
	tcs := []struct {
		desc   string
		input  string
		result int
	}{
		{
			desc:   "In this example, after blinking six times, you would have 22 stones. After blinking 25 times, you would have 55312 stones!",
			input:  "125 17",
			result: 55312,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getNumStones3([]string{tc.input}, 25)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
