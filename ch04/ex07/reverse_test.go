// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		s    []byte
		want []byte
	}{
		{[]byte("ABCDE"), []byte("EDCBA")},
		{[]byte("あいうえお"), []byte("おえういあ")},
	}

	for _, test := range tests {
		out := reverse(test.s)
		for i := 0; i < len(out); i++ {
			if out[i] != test.want[i] {
				t.Errorf("reverse(%v) = %v, want %v", string(test.s), string(out), string(test.want))
				break
			}
		}
	}
}
