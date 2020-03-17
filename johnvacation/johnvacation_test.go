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
		// {
		// 	desc:   "Test case 01 of challenge",
		// 	inputY: 2019,
		// 	inputA: "March",
		// 	inputB: "April",
		// 	inputW: "Tuesday",
		// 	output: 8,
		// },
		// {
		// 	desc:   "Test case 02 of challenge",
		// 	inputY: 2019,
		// 	inputA: "April",
		// 	inputB: "May",
		// 	inputW: "Tuesday",
		// 	output: 8,
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
		// {
		// 	desc:   "Test case 01 of challenge",
		// 	inputY: 2020,
		// 	inputA: "March",
		// 	inputB: "June",
		// 	inputW: "Wednesday",
		// 	output: 17,
		// },
		// {
		// 	desc:   "Test case 04",
		// 	inputY: 2020,
		// 	inputA: "March",
		// 	inputB: "March",
		// 	inputW: "Wednesday",
		// 	output: 4,
		// },
		{
			desc:   "Test case 04",
			inputY: 2019,
			inputA: "March",
			inputB: "March",
			inputW: "Tuesday",
			output: 4,
		},
		// {
		// 	desc:   "Test case example",
		// 	inputY: 2014,
		// 	inputA: "April",
		// 	inputB: "May",
		// 	inputW: "Wednesday",
		// 	output: 7,
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

func TestFirstMonday(t *testing.T) {
	testCases := []struct {
		desc   string
		y      int
		m      string
		fwd    string
		output int
	}{
		// {
		// 	desc:   "test 01",
		// 	y:      2014,
		// 	m:      "April",
		// 	fwd:    "Wednesday",
		// 	output: 7,
		// },
		{
			desc:   "test 02",
			y:      2020,
			m:      "April",
			fwd:    "Wednesday",
			output: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := johnvacation.LastSunday(tC.y, tC.m, tC.fwd); rs != tC.output {
				t.Errorf("got rs=%d, wants rs=%d", rs, tC.output)
			}
		})
	}
}
