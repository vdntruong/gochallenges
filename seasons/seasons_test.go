package seasons_test

import (
	"testing"

	"github.com/vdntruong/chl/seasons"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		inputAr []int
		output  string
	}{
		{
			desc:    "Test case 01 of challenge",
			inputAr: []int{-3, -14, -5, 7, 8, 42, 8, 3},
			output:  "SUMMER",
		}, {
			desc:    "Test case 02 of challenge",
			inputAr: []int{-2, 3, 3, -5, 3, 18, 1, 10, 8, 2, 5, 13},
			output:  "SPRING",
		}, {
			desc:    "Test case 03",
			inputAr: []int{5, 8, -2, -2, 3, 3, 1, 10, 8, 2, 5, 13},
			output:  "AUTUMN",
		}, {
			desc:    "Test case 04",
			inputAr: []int{0, -1000, 1000, 0, -5, 9, 0, -1000, 1000, 2, -4, -9},
			output:  "WINTER",
		}, {
			desc:    "Test case 05",
			inputAr: []int{-2, -2, -2, -3, -5, -7, 1, 2, 3, 3, 2, 1},
			output:  "SPRING",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := seasons.Solution(tC.inputAr); rs != tC.output {
				t.Errorf("Got rs=%s, wants rs=%s", rs, tC.output)
			}
		})
	}
}
