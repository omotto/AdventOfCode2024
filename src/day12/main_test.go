package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetPrices(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "In the first example, region A has price 4 * 10 = 40, region B has price 4 * 8 = 32, region C has price 4 * 10 = 40, region D has price 1 * 4 = 4, and region E has price 3 * 8 = 24. So, the total price for the first example is 140.",
			input: []string{
				"AAAA",
				"BBCD",
				"BBCC",
				"EEEC",
			},
			result: 140,
		},
		{
			desc: "In the second example, the region with all of the O plants has price 21 * 36 = 756, and each of the four smaller X regions has price 1 * 4 = 4, for a total price of 772 (756 + 4 + 4 + 4 + 4).",
			input: []string{
				"OOOOO",
				"OXOXO",
				"OOOOO",
				"OXOXO",
				"OOOOO",
			},
			result: 772,
		},
		{
			desc: "So, it has a total price of 1930",
			input: []string{
				"RRRRIICCFF",
				"RRRRIICCCF",
				"VVRRRCCFFF",
				"VVRCCCJFFF",
				"VVVVCJJCFE",
				"VVIVCCJJEE",
				"VVIIICJJEE",
				"MIIIIIJJEE",
				"MIIISIJEEE",
				"MMMISSJEEE",
			},
			result: 1930,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getPrices(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}

func TestGetPrices2(t *testing.T) {
	tcs := []struct {
		desc   string
		input  []string
		result int
	}{
		{
			desc: "Using the new method of calculating the per-region price by multiplying the region's area by its number of sides, regions A through E have prices 16, 16, 32, 4, and 12, respectively, for a total price of 80.",
			input: []string{
				"AAAA",
				"BBCD",
				"BBCC",
				"EEEC",
			},
			result: 80,
		},
		{
			desc: "The second example above (full of type X and O plants) would have a total price of 436.",
			input: []string{
				"OOOOO",
				"OXOXO",
				"OOOOO",
				"OXOXO",
				"OOOOO",
			},
			result: 436,
		},
		{
			desc: "Adding these together produces its new total price of 1206.",
			input: []string{
				"RRRRIICCFF",
				"RRRRIICCCF",
				"VVRRRCCFFF",
				"VVRCCCJFFF",
				"VVVVCJJCFE",
				"VVIVCCJJEE",
				"VVIIICJJEE",
				"MIIIIIJJEE",
				"MIIISIJEEE",
				"MMMISSJEEE",
			},
			result: 1206,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			got := getPrices2(tc.input)
			if diff := cmp.Diff(tc.result, got); diff != "" {
				t.Errorf("sum has diff %s", diff)
			}
		})
	}
}
