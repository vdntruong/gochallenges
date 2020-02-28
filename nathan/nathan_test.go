package nathan_test

import (
	"testing"

	"github.com/vdntruong/chl/nathan"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		inputArT []string
		inputArR []string
		output   int
	}{
		{
			desc:     "",
			inputArT: []string{"testla", "test2", "testlb"},
			inputArR: []string{"Wrong answer", "OK", "Runtime error"},
			output:   33,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := nathan.Solution(tC.inputArT, tC.inputArR); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
