package lovelynumber_test

import (
	"math/rand"
	"testing"

	"github.com/vdntruong/chl/lovelynumber"
)

var (
	testCases = []struct {
		name  string
		input struct {
			a, b int
		}
		output int
	}{
		{
			name: "number < 111",
			input: struct {
				a int
				b int
			}{0, 110},
			output: 111,
		},
		{
			name: "a = b with lovely",
			input: struct {
				a int
				b int
			}{0, 110},
			output: 111,
		},
		{
			name: "a = b but not lovely",
			input: struct {
				a int
				b int
			}{100_000, 100_000},
			output: 0,
		},
		{
			name: "a > b case",
			input: struct {
				a int
				b int
			}{100_000, 90_000},
			output: 0,
		},
		{
			name: "case all range number is lovely",
			input: struct {
				a int
				b int
			}{90_123, 90_132},
			output: 10,
		},
	}
)

func TestSolution(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			if rs := lovelynumber.Solution(c.input.a, c.input.b); rs != c.output {
				t.Errorf("got rs=%v, but want rs=%v", rs, c.output)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i := rand.Intn(len(testCases) - 0)
		lovelynumber.Solution(testCases[i].input.a, testCases[i].input.b)
	}
}
