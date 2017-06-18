// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"os"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var tests = []struct {
		words []byte
		want  int64
	}{
		{[]byte("aaa"), 3},
		{[]byte("bbb"), 6},
		{[]byte("ccc"), 9},
	}

	w, c := CountingWriter(os.Stdout)
	for _, test := range tests {
		w.Write(test.words)
		if *c != test.want {
			t.Errorf("Write(%v) = %d, want %d", test.words, *c, test.want)
		}
	}
}
