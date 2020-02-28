package reformatnumber_test

import (
	"testing"

	"github.com/vdntruong/chl/reformatnumber"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc     string
		inputstr string
		output   string
	}{
		{
			desc:     "Test case 01 of challenge",
			inputstr: "00-44 48 5555 8361",
			output:   "004-448-555-583-61",
		},
		{
			desc:     "Test case 02 of challenge",
			inputstr: "0 - 22 1985--324",
			output:   "022-198-53-24",
		},
		{
			desc:     "Test case 03 of challenge",
			inputstr: "555372654",
			output:   "555-372-654",
		},
		{
			desc:     "Test case 04",
			inputstr: "12",
			output:   "12",
		},
		{
			desc:     "Test case 05",
			inputstr: "123",
			output:   "123",
		},
		{
			desc:     "Test case 06",
			inputstr: "1234",
			output:   "12-34",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := reformatnumber.Solution(tC.inputstr); rs != tC.output {
				t.Errorf("got rs=%s, wants rs=%s", rs, tC.output)
			}
		})
	}
}
