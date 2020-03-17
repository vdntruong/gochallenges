package interestingtime_test

import (
	"testing"

	"github.com/vdntruong/chl/interestingtime"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc   string
		inputS string
		inputT string
		output int
	}{
		{
			desc:   "Test case 01 of challenge",
			inputS: "15:15:00",
			inputT: "15:15:12",
			output: 1,
		},
		{
			desc:   "Test case 03",
			inputS: "15:15:00",
			inputT: "16:16:00",
			output: 16,
		},
		{
			desc:   "Test case 03",
			inputS: "15:15:00",
			inputT: "17:17:00",
			output: 32,
		},
		{
			desc:   "Test case 04",
			inputS: "23:22:21",
			inputT: "23:22:30",
			output: 2,
		},
		// {
		// 	desc:   "Test case 02 of challenge",
		// 	inputS: "22:22:21",
		// 	inputT: "22:22:23",
		// 	output: 3,
		// },
		// {
		// 	desc:   "Test case 03 of challenge",
		// 	inputS: "00:00:00",
		// 	inputT: "22:22:23",
		// 	output: 3,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := interestingtime.Solution(tC.inputS, tC.inputT); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
