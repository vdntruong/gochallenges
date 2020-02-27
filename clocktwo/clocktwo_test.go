package clocktwo_test

import (
	"testing"

	"github.com/vdntruong/chl/clocktwo"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc  string
		input []int
		out   int
	}{
		{
			desc:  "Negative value",
			input: []int{1, 2, 3, -4},
			out:   0,
		},
		{
			desc:  "0 number for 4 digital (mid-night)",
			input: []int{0, 0, 0, 0},
			out:   1,
		},
		{
			desc:  "End of day (EOD)",
			input: []int{5, 9, 3, 2},
			out:   1,
		},
		{
			desc:  "0 number for 2 digital, orther min value for orther 2 digital (On-times, mid-night times, day times)",
			input: []int{0, 1, 0, 2},
			out:   12,
		},
		{
			desc:  "0 number for 2 digital, orther value for orther 2 digital (mid-night times, day times, out range hours)",
			input: []int{0, 2, 0, 3},
			out:   9,
		},
		{
			desc:  "0 number for 3 digital, one number valid",
			input: []int{0, 0, 1, 0},
			out:   4,
		},
		{
			desc:  "0 number for 3 digital, one number out of range hour time",
			input: []int{0, 0, 3, 0},
			out:   3,
		},
		{
			desc:  "0 number for 1 digital, min values for 3 digital, MAX case",
			input: []int{0, 1, 2, 3},
			out:   18,
		},
		{
			desc:  "0 number for 3 digital, one number out of range minute time",
			input: []int{0, 0, 7, 0},
			out:   2,
		},
		{
			desc:  "1 value for all digital (valid times)",
			input: []int{2, 2, 2, 2},
			out:   1,
		},
		{
			desc:  "1 value for all digital (invalid hour times)",
			input: []int{3, 3, 3, 3},
			out:   0,
		},
		{
			desc:  "2 values for 4 digital, just max value valid minute times, 11:33, 13:13, 13:31",
			input: []int{1, 1, 3, 3},
			out:   3,
		},
		{
			desc:  "2 values for 4 digital, valid times",
			input: []int{1, 1, 2, 2},
			out:   6,
		},
		{
			desc:  "Valid hour, but out of range minute times",
			input: []int{2, 4, 6, 6},
			out:   0,
		},
		{
			desc:  "Valid minute, but out of range hour times",
			input: []int{3, 3, 5, 9},
			out:   0,
		},
		{
			desc:  "1234 values",
			input: []int{1, 2, 3, 4},
			out:   10,
		},
		{
			desc:  "Test case 01 of challenge",
			input: []int{1, 8, 3, 2},
			out:   6,
		},
		{
			desc:  "Test case 02 of challenge",
			input: []int{2, 3, 3, 2},
			out:   3,
		},
		{
			desc:  "Test case 03 of challenge",
			input: []int{6, 2, 4, 7},
			out:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := clocktwo.Solution(tC.input); rs != tC.out {
				t.Errorf("Got result=%d, wants result=%d", rs, tC.out)
			}
		})
	}
}
