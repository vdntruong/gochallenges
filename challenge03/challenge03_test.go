package challenge03_test

import (
	"testing"

	"github.com/vdntruong/chl/challenge03"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output int
	}{
		{
			desc:   "01",
			input:  []int{3, 4, 5, 3, 7},
			output: 3,
		},
		{
			desc:   "02",
			input:  []int{1, 2, 3, 4},
			output: -1,
		},
		{
			desc:   "03",
			input:  []int{1, 3, 1, 2},
			output: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := challenge03.Solution(tC.input); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
