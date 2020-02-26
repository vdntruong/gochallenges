package thering_test

import (
	"testing"

	"github.com/vdntruong/chl/thering"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc             string
		inputInner       int
		inputOuter       int
		inputX           []int
		inputY           []int
		outputTotalPoint int
	}{
		{
			desc:             "Test case 01 of challenge, in/on inner cỉrcle",
			inputInner:       1,
			inputOuter:       3,
			inputX:           []int{0, 0, 0, 1, -1},
			inputY:           []int{0, 1, -1, 0, 0},
			outputTotalPoint: 0,
		},
		{
			desc:       "Test case 01 of challenge, out/on outer cỉrcle",
			inputInner: 1,
			inputOuter: 3,
			inputX: []int{
				0, 1, 2, 3,
				3, 3, 3,
				3, 3, 3,
				2, 1, 0,
				-1, -2, -3,
				-3, -3, -3,
				-3, -3, -3,
				-2, -1, 0},
			inputY: []int{
				3, 3, 3, 3,
				2, 1, 0,
				-1, -2, -3,
				-3, -3, -3,
				-3, -3, -3,
				-2, -1, 0,
				1, 2, 3,
				3, 3, 3},
			outputTotalPoint: 0,
		},
		{
			desc:       "Test case 01 of challenge - all valid points",
			inputInner: 1,
			inputOuter: 3,
			inputX: []int{
				1, -2,
				2, 2, 2, 1,
				0, -1, -2, -1,
				-2, -1,
				-1, 0, 1, 2, 1, 2,
			},
			inputY: []int{
				1, 1,
				0, 1, 2, 2,
				2, 2, -2, 1,
				-1, -1,
				-2, -2, -2, -2, -1, -1,
			},
			outputTotalPoint: 18,
		},
		{
			desc:       "Test case 01 of challenge - mixed",
			inputInner: 1,
			inputOuter: 3,
			inputX: []int{0, 1, 2, -2, 3,
				0, 0, 1, -1,
				0, 1, 2, 3,
				3, 3, 3,
				3, 3, 3,
				2, 1, 0,
				-1, -2, -3,
				-3, -3, -3,
				-3, -3, -3,
				-2, -1, 0,
				0, 1, 0, -1,

				2, 2, 2, 1,
				0, -1, -2, -1,
				0, -2, -1,
				-1, 1, 2, 1, 2,
			},
			inputY: []int{0, 1, 4, 1, 0,
				1, -1, 0, 0,
				3, 3, 3, 3,
				2, 1, 0,
				-1, -2, -3,
				-3, -3, -3,
				-3, -3, -3,
				-2, -1, 0,
				1, 2, 3,
				3, 3, 3,
				1, 0, -1, 0,

				0, 1, 2, 2,
				2, 2, -2, 1,
				-2, -1, -1,
				-2, -2, -2, -1, -1,
			},
			outputTotalPoint: 18,
		},
		{
			desc:       "Test case 02 of challenge - in/on inner circle",
			inputInner: 2,
			inputOuter: 4,
			inputX: []int{
				0, 0, 0, 0, 0,
				-1, -1, -1, -2,
				1, 1, 1, 2,
			},
			inputY: []int{
				2, 1, 0, -1, -2,
				1, 0, -1, 0,
				1, 0, -1, 0,
			},
			outputTotalPoint: 0,
		},
		{
			desc:       "Test case 02 of challenge - all valid points",
			inputInner: 2,
			inputOuter: 4,
			inputX: []int{
				-3, -3, -3, -3, -3,
				-2, -2, -2, -2, -2, -2,
				-1, -1, -1, -1,
				0, 0,
				1, 1, 1, 1,
				2, 2, 2, 2, 2, 2,
				3, 3, 3, 3, 3,
			},
			inputY: []int{2, 1, 0, -1, -2,
				3, 2, 1, -1, -2, -3,
				3, 2, -2, -3,
				3, -3,
				3, 2, -2, -3,
				3, 2, 1, -1, -2, -3,
				2, 1, 0, -1, -2,
			},
			outputTotalPoint: 32,
		},
		{
			desc:       "Test case 02 of challenge - out/on outer circle",
			inputInner: 2,
			inputOuter: 4,
			inputX: []int{
				-4, -4, -4, -4, -4,
				-3, -3,
				-2, -2,
				-1, -1,
				0, 0,
				1, 1,
				2, 2,
				3, 3, 3, 3,
				4, 4, 4, 4, 4, 4, 4,
			},
			inputY: []int{
				2, 1, 0, -1, -2,
				3, -3,
				4, -4,
				4, -4,
				4, -4,
				4, -4,
				4, -4,
				4, 3, -3, -4,
				3, 2, 1, 0, -1, -2, -3,
			},
			outputTotalPoint: 0,
		},
		{
			desc:       "Test case 02 of challenge - mixed",
			inputInner: 2,
			inputOuter: 4,
			inputX: []int{
				0, 0, 0, 0, 0,
				-1, -1, -1, -2,
				1, 1, 1, 2,
				-3, -3, -3, -3, -3,
				-2, -2, -2, -2, -2, -2,
				-1, -1, -1, -1,
				0, 0,
				1, 1, 1, 1,
				2, 2, 2, 2, 2, 2,
				3, 3, 3, 3, 3,
				-4, -4, -4, -4, -4,
				-3, -3,
				-2, -2,
				-1, -1,
				0, 0,
				1, 1,
				2, 2,
				3, 3, 3, 3,
				4, 4, 4, 4, 4, 4, 4,
			},
			inputY: []int{
				2, 1, 0, -1, -2,
				1, 0, -1, 0,
				1, 0, -1, 0,
				2, 1, 0, -1, -2,
				3, 2, 1, -1, -2, -3,
				3, 2, -2, -3,
				3, -3,
				3, 2, -2, -3,
				3, 2, 1, -1, -2, -3,
				2, 1, 0, -1, -2,
				2, 1, 0, -1, -2,
				3, -3,
				4, -4,
				4, -4,
				4, -4,
				4, -4,
				4, -4,
				4, 3, -3, -4,
				3, 2, 1, 0, -1, -2, -3,
			},
			outputTotalPoint: 32,
		},
		{
			desc:             "Test case ",
			inputInner:       1,
			inputOuter:       2,
			inputX:           []int{0, 0, 0, 1, -1, 1, 1},
			inputY:           []int{0, 1, -1, 0, 0, 1, -1},
			outputTotalPoint: 0,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := thering.Solution(tC.inputInner, tC.inputOuter, tC.inputX, tC.inputY); rs != tC.outputTotalPoint {
				t.Errorf("got total=%d, wants total point=%d", rs, tC.outputTotalPoint)
			}
		})
	}
}
