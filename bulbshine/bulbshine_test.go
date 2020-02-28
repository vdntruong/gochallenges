package bulbshine_test

import (
	"testing"

	"github.com/vdntruong/chl/bulbshine"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc    string
		inputAr []int
		output  int
	}{
		{
			desc:    "",
			inputAr: []int{2, 3, 4, 1, 5},
			output:  2,
		},
		{
			desc:    "02",
			inputAr: []int{1, 3, 4, 2, 5},
			output:  3,
		},
		{
			desc:    "03",
			inputAr: []int{5, 4, 3, 2, 1},
			output:  1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := bulbshine.Solution(tC.inputAr); rs != tC.output {
				t.Errorf("Got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
