package equilibriummatrix_test

import (
	"testing"

	"github.com/vdntruong/chl/equilibriummatrix"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc	string
		inputMatrix [][]int
		output int 

	}{
		{
			desc: "Test case 01 of challenge",
			inputMatrix: [][]int{
				[]int{2,7,5},
				[]int{3,1,1},
				[]int{2,1,-7},
				[]int{0,2,1},
				[]int{1,6,8},
			},
			output: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := equilibriummatrix.Solution(tC.inputMatrix); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
