// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"testing"
)

func TestWordCounter(t *testing.T) {
	var tests = []struct {
		words []byte
		want  int
	}{
		{[]byte(""), 0},
		{[]byte("1"), 1},
		{[]byte("1 2"), 2},
		{[]byte("1 2 3"), 3},
	}

	for _, test := range tests {
		var c WordCounter
		c.Write(test.words)
		if int(c) != test.want {
			t.Errorf("Write(%v) = %d, want %d", test.words, int(c), test.want)
		}
	}
}

func TestLineCounter(t *testing.T) {
	var tests = []struct {
		words []byte
		want  int
	}{
		{[]byte(""), 0},
		{[]byte("1"), 1},
		{[]byte("1\n2 2"), 2},
		{[]byte("1\n2 2\n3 3 3"), 3},
	}

	for _, test := range tests {
		var c LineCounter
		c.Write(test.words)
		if int(c) != test.want {
			t.Errorf("Write(%v) = %d, want %d", test.words, int(c), test.want)
		}
	}
}
