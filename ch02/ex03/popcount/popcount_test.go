// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{0, 0},
		{1, 1},
		{255, 8},
		{256, 1},
	}

	for _, test := range tests {
		pc := PopCount(test.x)
		if pc != test.want {
			t.Errorf("PopCount(%d) = %d, want %d", test.x, pc, test.want)
		}

		pc = PopCountByLoop(test.x)
		if pc != test.want {
			t.Errorf("PopCountByLoop(%d) = %d, want %d", test.x, pc, test.want)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByLoop(0x1234567890ABCDEF)
	}
}
