package stableperiod_test

import (
	"github.com/vdntruong/chl/stableperiod"
	"testing"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc	string
		inputAr []int
		output int
	}{
		{
			desc: "Test case 01 of challenge",
			inputAr: []int{-1, 1, 3, 3, 3, 2, 3, 2, 1, 0},
			output: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := stableperiod.Solution(tC.inputAr); rs != tC.output{
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}