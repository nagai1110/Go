// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bytes"
	"os"
	"testing"
)

func TestDup(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"command", "test1.txt", "test2.txt", "test3.txt"}, "a\t[test1.txt test2.txt test3.txt]\nbb\t[test2.txt test3.txt]\n"},
	}

	for _, test := range tests {
		os.Args = test.args
		out = new(bytes.Buffer)
		main()

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("main() = %q, want %q", got, test.want)
		}
	}
}
