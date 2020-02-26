package stringobtain_test

import (
	"testing"

	"github.com/vdntruong/chl/stringobtain"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc           string
		inputS, inputT string
		output         string
	}{
		{
			desc:   "NOTHING",
			inputS: "nice",
			inputT: "nice",
			output: "NOTHING",
		},
		{
			desc:   "Add at the end of string",
			inputS: "nice",
			inputT: "nicer",
			output: "ADD r",
		},
		{
			desc:   "Add at middle of string",
			inputS: "nice",
			inputT: "niece",
			output: "ADD e",
		},
		{
			desc:   "Add at the first of string",
			inputS: "cho",
			inputT: "echo",
			output: "ADD e",
		},
		{
			desc:   "Change words, at the middle",
			inputS: "test",
			inputT: "tent",
			output: "CHANGE s n",
		},
		{
			desc:   "Change words, at the end",
			inputS: "tent",
			inputT: "tens",
			output: "CHANGE t s",
		},
		{
			desc:   "Change words, at the first",
			inputS: "hent",
			inputT: "tent",
			output: "CHANGE h t",
		},
		{
			desc:   "Move words, at the first",
			inputS: "beans",
			inputT: "ebans",
			output: "MOVE b",
		},
		{
			desc:   "Move words, at the middle",
			inputS: "beans",
			inputT: "banes",
			output: "MOVE e",
		},
		{
			desc:   "Move words, at the end",
			inputS: "beans",
			inputT: "beasn",
			output: "MOVE s",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := stringobtain.Solution(tC.inputS, tC.inputT); rs != tC.output {
				t.Errorf("got rs=%s, wants rs=%s", rs, tC.output)
			}
		})
	}
}
