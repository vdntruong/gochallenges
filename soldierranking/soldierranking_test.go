package soldierranking_test

import (
	"testing"
	"github.com/vdntruong/chl/soldierranking"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc	string
		inputAr []int
		output int
	}{
		{
			desc: "Test case 01 of challenge",
			inputAr: []int{3,4,3,0,2,2,3,0,0},
			output: 5,
		},
		{
			desc: "Test case 02 of challenge",
			inputAr: []int{4,0,2},
			output: 0,
		},
		{
			desc: "Test case 03 of challenge",
			inputAr: []int{4,4,3,3,1,0},
			output: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := soldierranking.Solution(tC.inputAr); rs != tC.output{
				t.Errorf("Got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}