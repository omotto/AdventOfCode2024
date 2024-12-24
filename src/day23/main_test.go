package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSumTLANNetworks(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "That narrows the list down to 7 sets of three inter-connected computers.",
			input: []string{
				"kh-tc",
				"qp-kh",
				"de-cg",
				"ka-co",
				"yn-aq",
				"qp-ub",
				"cg-tb",
				"vc-aq",
				"tb-ka",
				"wh-tc",
				"yn-cg",
				"kh-ub",
				"ta-co",
				"de-co",
				"tc-td",
				"tb-wq",
				"wh-td",
				"ta-ka",
				"td-qp",
				"aq-cg",
				"wq-ub",
				"ub-vc",
				"de-ta",
				"wq-aq",
				"wq-vc",
				"wh-yn",
				"ka-de",
				"kh-ta",
				"co-tc",
				"wh-qp",
				"tb-vc",
				"td-yn",
			},
			result: 7,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getSumTLANNetworks(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestPassword(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result string
	}{
		{
			desc: "In this example, the password would be co,de,ka,ta.",
			input: []string{
				"kh-tc",
				"qp-kh",
				"de-cg",
				"ka-co",
				"yn-aq",
				"qp-ub",
				"cg-tb",
				"vc-aq",
				"tb-ka",
				"wh-tc",
				"yn-cg",
				"kh-ub",
				"ta-co",
				"de-co",
				"tc-td",
				"tb-wq",
				"wh-td",
				"ta-ka",
				"td-qp",
				"aq-cg",
				"wq-ub",
				"ub-vc",
				"de-ta",
				"wq-aq",
				"wq-vc",
				"wh-yn",
				"ka-de",
				"kh-ta",
				"co-tc",
				"wh-qp",
				"tb-vc",
				"td-yn",
			},
			result: "co,de,ka,ta",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getPassword(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
