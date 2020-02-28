package sortingarray_test

import (
	"testing"

	"github.com/vdntruong/chl/sortingarray"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		inputAr []int
		output  bool
	}{
		{
			desc:    "Test case 01 of challenge",
			inputAr: []int{1, 5, 3, 3, 7},
			output:  true,
		},
		{
			desc:    "Test case 02 of challenge",
			inputAr: []int{1, 3, 5, 3, 4},
			output:  false,
		},
		{
			desc:    "Test case 03",
			inputAr: []int{0, 1, -1},
			output:  false,
		},
		{
			desc:    "Test case 04",
			inputAr: []int{1, 2, 3},
			output:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := sortingarray.Solution(tC.inputAr); rs != tC.output {
				t.Errorf("Got rs=%v, wants rs=%v", rs, tC.output)
			}
		})
	}
}
