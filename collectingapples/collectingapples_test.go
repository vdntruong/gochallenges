package collectingapples_test

import (
	"testing"

	"github.com/vdntruong/chl/collectingapples"
)

type (
	inputStr struct {
		A    []int
		K, L int
	}
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name   string
		input  inputStr
		output int
	}{
		{
			name: "Case 01",
			input: inputStr{
				A: []int{1, 2, 3, 4, 5, 6},
				K: 4,
				L: 2,
			},
			output: 21,
		},
		{
			name: "Case 02",
			input: inputStr{
				A: []int{1, 2, 3, 4, 5},
				K: 4,
				L: 2,
			},
			output: -1,
		},
		{
			name: "Case 03",
			input: inputStr{
				A: []int{6, 1, 4, 6, 3, 2, 7, 4},
				K: 3,
				L: 2,
			},
			output: 24,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := collectingapples.Solution(test.input.A, test.input.K, test.input.L)
			if got != test.output {
				t.Errorf("got %v, wants %v", got, test.output)
			}
		})
	}
}
