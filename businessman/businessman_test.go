package businessman_test

import (
	"testing"

	"github.com/vdntruong/chl/businessman"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		n string
		i string
		o int
	}{
		{
			n: "One",
			i: `Sun 10:00-20:00
			Fri 05:00-10:00
			Fri 16:30-23:50
			Sat 10:00-24:00
			Sun 01:00-04:00
			Sat 02:00-06:00
			Tue 03:30-18:15
			Tue 19:00-20:00
			Wed 04:25-15:14
			Wed 15:14-22:40
			Thu 00:00-23:59
			Mon 05:00-13:00
			Mon 15:00-21:00`,
			o: 505,
		}, {
			n: "Two",
			i: `Mon 01:00-23:00
			Tue 01:00-23:00
			Wed 01:00-23:00
			Thu 01:00-23:00
			Fri 01:00-23:00
			Sat 01:00-23:00
			Sun 01:00-21:00`,
			o: 180,
		}, {
			n: "Three",
			i: `Mon 00:00-24:00
			Tue 00:00-24:00
			Wed 00:00-05:00
			Wed 07:00-24:00
			Thu 00:00-24:00
			Fri 00:00-24:00
			Sat 00:00-24:00
			Sun 00:00-24:00`,
			o: 120,
		},
	}

	for _, test := range testCases {
		t.Run(test.n, func(t *testing.T) {
			got := businessman.Solution(test.i)
			if got != test.o {
				t.Errorf("got %v, wants %v", got, test.o)
			}
		})
	}
}
