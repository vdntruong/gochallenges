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
			desc:   "Test case 01 of challenge",
			inputY: 2019,
			inputA: "March",
			inputB: "April",
			inputW: "Tuesday",
			output: 8,
		},
		// {
		// 	desc:   "Test case 01 of challenge",
		// 	inputY: 2019,
		// 	inputA: "April",
		// 	inputB: "May",
		// 	inputW: "Tuesday",
		// 	output: 7,
		// },
		// {
		// 	desc:   "Test case 02",
		// 	inputY: 2020,
		// 	inputA: "January",
		// 	inputB: "February",
		// 	inputW: "Wednesday",
		// 	output: 7,
		// },
		// {
		// 	desc:   "Test case 03",
		// 	inputY: 2020,
		// 	inputA: "July",
		// 	inputB: "August",
		// 	inputW: "Wednesday",
		// 	output: 8,
		// },
		// {
		// 	desc:   "Test case 03.5",
		// 	inputY: 2020,
		// 	inputA: "February",
		// 	inputB: "March",
		// 	inputW: "Wednesday",
		// 	output: 8,
		// },
		// {
		// 	desc:   "Test case 04",
		// 	inputY: 2020,
		// 	inputA: "February",
		// 	inputB: "April",
		// 	inputW: "Wednesday",
		// 	output: 12,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := johnvacation.Solution(tC.inputY, tC.inputA, tC.inputB, tC.inputW); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
