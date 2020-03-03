package rotatedice_test

import (
	"testing"

	"github.com/vdntruong/chl/rotatedice"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		inputAr []int
		output  int
	}{
		{
			desc:    "Test case 01 of challenge",
			inputAr: []int{1, 2, 3},
			output:  2,
		},
		{
			desc:    "Test case 02 of challenge",
			inputAr: []int{1, 1, 6},
			output:  2,
		},
		{
			desc:    "Test case 03 of challenge",
			inputAr: []int{1, 6, 2, 3},
			output:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := rotatedice.Solution(tC.inputAr); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
