package diversestring_test

import (
	"testing"

	"github.com/vdntruong/chl/diversestring"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc                   string
		inputA, inputB, inputC int
		output                 string
	}{
		{
			desc:   "Test case 01 of challenge",
			inputA: 0,
			inputB: 1,
			inputC: 2,
			output: "cbc",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := diversestring.Solution(tC.inputA, tC.inputB, tC.inputC); rs != tC.output {
				t.Errorf("got rs=%v, wants rs=%v", rs, tC.output)
			}
		})
	}
}
