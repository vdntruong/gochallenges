package johnvacation_test

import (
	"testing"

	"github.com/vdntruong/chl/johnvacation"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		inputY int
		inputA string
		inputB string
		inputW string
		output int
	}{
		{
			desc: "Test case 01 of challenge",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := johnvacation.Solution(tC.inputY, tC.inputA, tC.inputB, tC.inputW); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
