package cuttingtree_test

import (
	"testing"

	"github.com/vdntruong/chl/cuttingtree"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		inputAr []int
		output  int
	}{
		{
			desc:    "Test case 01 of challenge",
			inputAr: []int{3, 4, 5, 4},
			output:  2,
		},
		{
			desc:    "Test case 02 of challenge",
			inputAr: []int{4, 5, 2, 3, 4},
			output:  0,
		},
		{
			desc:    "Test case 03 of challenge",
			inputAr: []int{1, 2, 3, 3, 5, 6, 7},
			output:  7,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := cuttingtree.Solution(tC.inputAr); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
