package longestcentence_test

import (
	"testing"

	"github.com/vdntruong/chl/longestcentence"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		output int
	}{
		{
			desc:   "Test case 01 of challenge",
			input:  "We test coders. Give us a try?",
			output: 4,
		},
		{
			desc:   "Test case 02 of challenge",
			input:  "Forget    CVs.. Save time . x x",
			output: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := longestcentence.Solution(tC.input); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
