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
		{
			desc:   "Test case 02",
			inputA: 8,
			inputB: 8,
			inputC: 1,
			output: "cbc",
		},
		{
			desc:   "Test case Quan",
			inputA: 8,
			inputB: 8,
			inputC: 2,
			output: "cbc",
		},
		// {
		// 	desc:   "Test case 02",
		// 	inputA: 6,
		// 	inputB: 4,
		// 	inputC: 2,
		// 	output: "aababacbaacb",
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := diversestring.Solution(tC.inputA, tC.inputB, tC.inputC); rs != tC.output {
				t.Errorf("got rs=%v, wants rs=%v", rs, tC.output)
			}
		})
	}
}
