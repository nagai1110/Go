// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestDuplicate(t *testing.T) {
	var tests = []struct {
		s    []byte
		want []byte
	}{
		{[]byte("1  2  3"), []byte("1 2 3")},
		{[]byte("AA       BB       CC"), []byte("AA BB CC")},
	}

	for _, test := range tests {
		out := removeDuplicateSpace(test.s)
		for i := 0; i < len(out); i++ {
			if out[i] != test.want[i] {
				t.Errorf("removeDuplicateSpace(%v) = %v, want %v", string(test.s), string(out), string(test.want))
				break
			}
		}
	}
}
