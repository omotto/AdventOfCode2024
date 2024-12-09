package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDecompressCode(t *testing.T) {
	tcs := []struct {
		desc   string
		input  string
		result []int
	}{
		{
			desc:   "The first example above, 2333133121414131402, represents these individual blocks 00...111...2...333.44.5555.6666.777.888899",
			input:  "2333133121414131402",
			result: []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got, _, _, _ := decompressCode(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetFilledCode(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []int
		length int
		result []int
	}{
		{
			desc:   "The first example requires a few more steps 0099811188827773336446555566..............",
			input:  []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
			length: 28,
			result: []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getFilledCode(tc.input, tc.length)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetChecksum(t *testing.T) {
	tcs := []struct {
		desc   string
		input  string
		result int
	}{
		{
			desc:   "the checksum is the sum of these, 1928.",
			input:  "2333133121414131402",
			result: 1928,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getChecksum([]string{tc.input})
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetChecksumBlocks(t *testing.T) {
	tcs := []struct {
		desc   string
		input  string
		result int
	}{
		{
			desc:   "the checksum is the sum of these, 2858.",
			input:  "2333133121414131402",
			result: 2858,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getChecksumBlocks([]string{tc.input})
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
