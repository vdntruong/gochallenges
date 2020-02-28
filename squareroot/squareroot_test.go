package squareroot_test

import (
	"testing"

	"github.com/vdntruong/chl/squareroot"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		inputA int
		inputB int
		output int
	}{
		{
			desc:   "Test case 01 of challenge",
			inputA: 10,
			inputB: 20,
			output: 2,
		},
		{
			desc:   "Test case 02 of challenge",
			inputA: 6000,
			inputB: 7000,
			output: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := squareroot.Solution(tC.inputA, tC.inputB); rs != tC.output {
				t.Errorf("Got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
